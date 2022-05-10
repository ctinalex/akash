package provider

import (
	"context"
	"time"

	"github.com/boz/go-lifecycle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	btypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tendermint/tendermint/libs/log"
	tmrpc "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/ovrclk/akash/client"
	"github.com/ovrclk/akash/provider/event"
	"github.com/ovrclk/akash/provider/session"
	"github.com/ovrclk/akash/pubsub"
	"github.com/ovrclk/akash/util/runner"
	dtypes "github.com/ovrclk/akash/x/deployment/types/v1beta2"
	mtypes "github.com/ovrclk/akash/x/market/types/v1beta2"

	aclient "github.com/ovrclk/akash/client"
)

const blockPeriod = 6500 * time.Millisecond

var (
	balanceGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "provider_balance",
	})
)

type balanceChecker struct {
	session session.Session
	log     log.Logger
	lc      lifecycle.Lifecycle
	bus     pubsub.Bus
	ownAddr sdk.AccAddress
	bqc     btypes.QueryClient
	aqc     aclient.QueryClient
	leases  map[mtypes.LeaseID]*time.Timer
	cfg     BalanceCheckerConfig
}

type BalanceCheckerConfig struct {
	PollingPeriod           time.Duration
	MinimumBalanceThreshold uint64
	WithdrawalPeriod        time.Duration
	LeaseFundCheckInterval  time.Duration
}

type leaseCheckResponse struct {
	lid        mtypes.LeaseID
	checkAfter time.Duration // potential out of funds after this amount of time spent since now. 0 means lease already dummed
	err        error
}

func newBalanceChecker(ctx context.Context,
	bqc btypes.QueryClient,
	aqc client.QueryClient,
	accAddr sdk.AccAddress,
	clientSession session.Session,
	bus pubsub.Bus,
	cfg BalanceCheckerConfig) *balanceChecker {

	bc := &balanceChecker{
		session: clientSession,
		log:     clientSession.Log().With("cmp", "balance-checker"),
		lc:      lifecycle.New(),
		bus:     bus,
		ownAddr: accAddr,
		bqc:     bqc,
		aqc:     aqc,
		leases:  make(map[mtypes.LeaseID]*time.Timer),
		cfg:     cfg,
	}

	go bc.lc.WatchContext(ctx)
	go bc.run()

	return bc
}

func (bc *balanceChecker) doCheck(ctx context.Context) (bool, error) {
	// Get the current wallet balance
	query := btypes.NewQueryBalanceRequest(bc.ownAddr, "uakt")
	result, err := bc.bqc.Balance(ctx, query)
	if err != nil {
		return false, err
	}

	balance := result.Balance.Amount
	balanceGauge.Set(float64(balance.Uint64()))

	if bc.cfg.MinimumBalanceThreshold == 0 {
		return false, nil
	}

	tooLow := sdk.NewIntFromUint64(bc.cfg.MinimumBalanceThreshold).GT(balance)

	return tooLow, nil
}

func (bc *balanceChecker) runEscrowCheck(ctx context.Context, lid mtypes.LeaseID, res chan<- leaseCheckResponse) {
	go func() {
		select {
		case <-bc.lc.Done():
		case res <- bc.doEscrowCheck(ctx, lid):
		}
	}()
}

func (bc *balanceChecker) doEscrowCheck(ctx context.Context, lid mtypes.LeaseID) leaseCheckResponse {
	resp := leaseCheckResponse{
		lid: lid,
	}

	var syncInfo *tmrpc.SyncInfo
	syncInfo, resp.err = bc.session.Client().NodeSyncInfo(ctx)
	if resp.err != nil {
		return resp
	}

	if !syncInfo.CatchingUp {
		resp.err = aclient.ErrNodeNotSynced
		return resp
	}

	var dResp *dtypes.QueryDeploymentResponse
	var lResp *mtypes.QueryLeasesResponse

	// Fetch the balance of the escrow account
	dResp, resp.err = bc.aqc.Deployment(ctx, &dtypes.QueryDeploymentRequest{
		ID: lid.DeploymentID(),
	})

	if resp.err != nil {
		return resp
	}

	lResp, resp.err = bc.aqc.Leases(ctx, &mtypes.QueryLeasesRequest{
		Filters: mtypes.LeaseFilters{
			Owner: lid.Owner,
			DSeq:  lid.DSeq,
			State: "active",
		},
	})

	if resp.err != nil {
		return resp
	}

	totalLeaseAmount := sdk.NewDec(0)
	for _, lease := range lResp.Leases {
		totalLeaseAmount = totalLeaseAmount.Add(lease.Lease.Price.Amount)
	}

	if resp.err == nil {
		balance := dResp.EscrowAccount.Balance.Amount
		settledAt := dResp.EscrowAccount.SettledAt
		balanceRemaining := balance.MustFloat64() - (float64(syncInfo.LatestBlockHeight-settledAt))*totalLeaseAmount.MustFloat64()
		blocksRemaining := int64(balanceRemaining / totalLeaseAmount.MustFloat64())

		// lease is out of funds
		if blocksRemaining <= 0 {
			return resp
		}

		blocksPerCheckInterval := int64(bc.cfg.LeaseFundCheckInterval / blockPeriod)
		if blocksRemaining > blocksPerCheckInterval {
			blocksRemaining = blocksPerCheckInterval
		}

		resp.checkAfter = time.Duration(blocksRemaining) * (6500 * time.Millisecond)
	}

	return resp
}

func (bc *balanceChecker) startWithdrawAll() error {
	return bc.bus.Publish(event.LeaseWithdrawNow{})
}

func (bc *balanceChecker) run() {
	defer bc.lc.ShutdownCompleted()
	ctx, cancel := context.WithCancel(context.Background())

	tick := time.NewTicker(bc.cfg.PollingPeriod)
	withdrawalTicker := time.NewTicker(bc.cfg.WithdrawalPeriod)

	var balanceCheckResult <-chan runner.Result
	var withdrawAllResult <-chan runner.Result

	leaseCheckCh := make(chan leaseCheckResponse, 1)

	subscriber, _ := bc.bus.Subscribe()

loop:
	for {
		withdrawAllNow := false

		select {
		case err := <-bc.lc.ShutdownRequest():
			bc.log.Debug("shutting down")
			bc.lc.ShutdownInitiated(err)
			break loop
		case <-tick.C:
			tick.Stop() // Stop the timer
			// Start the balance check
			balanceCheckResult = runner.Do(func() runner.Result {
				return runner.NewResult(bc.doCheck(ctx))
			})
		case evt := <-subscriber.Events():
			switch ev := evt.(type) {
			case event.LeaseAddFundsMonitor:
				bc.leases[ev.LeaseID] = nil
				if bc.cfg.LeaseFundCheckInterval > 0 {
					bc.runEscrowCheck(ctx, ev.LeaseID, leaseCheckCh)
				}
			case event.LeaseRemoveFundsMonitor:
				tm, exists := bc.leases[ev.LeaseID]
				if !exists {
					break
				}

				delete(bc.leases, ev.LeaseID)

				if tm != nil && !tm.Stop() {
					<-tm.C
				}
			}
		case res := <-leaseCheckCh:
			// we may have timer fired just a heart beat ahead of lease remove event.
			if _, exists := bc.leases[res.lid]; !exists {
				break
			}

			if res.err != nil {
				bc.log.Info("couldn't check lease balance. retrying in 1m", "leaseId", res.lid, "error", res.err.Error())
				bc.leases[res.lid] = time.AfterFunc(time.Minute, func() {
					select {
					case <-bc.lc.Done():
					case leaseCheckCh <- bc.doEscrowCheck(ctx, res.lid):
					}
				})
				break
			}

			if res.checkAfter == 0 {
				bc.log.Debug("lease is out of funds. closing", "leaseId", res.lid)

				_ = runner.Do(func() runner.Result {
					return runner.NewResult(nil, bc.session.Client().Tx().Broadcast(ctx, mtypes.NewMsgCloseBid(res.lid.BidID())))
				})

				break
			}

			bc.leases[res.lid] = time.AfterFunc(res.checkAfter, func() {
				select {
				case <-bc.lc.Done():
				case leaseCheckCh <- bc.doEscrowCheck(ctx, res.lid):
				}
			})
		case balanceCheck := <-balanceCheckResult:
			balanceCheckResult = nil
			tick.Reset(bc.cfg.PollingPeriod) // Re-enable the timer
			err := balanceCheck.Error()
			if err != nil {
				bc.log.Error("failed to check balance", "err", err)
				break
			}

			tooLow := balanceCheck.Value().(bool)
			if tooLow {
				// trigger the withdrawal
				bc.log.Info("balance below target amount")
				withdrawAllNow = true
			}
		case withdrawAll := <-withdrawAllResult:
			withdrawAllResult = nil
			withdrawalTicker.Reset(bc.cfg.PollingPeriod) // Re-enable the timer
			if err := withdrawAll.Error(); err != nil {
				bc.log.Error("failed to started withdrawals", "err", err)
			}
		case <-withdrawalTicker.C:
			withdrawAllNow = true
			withdrawalTicker.Stop()
		}

		if withdrawAllNow {
			bc.log.Info("balance below target amount, withdrawing now")
			withdrawAllResult = runner.Do(func() runner.Result {
				return runner.NewResult(nil, bc.startWithdrawAll())
			})
		}
	}
	cancel()

	bc.log.Debug("shutdown complete")
}
