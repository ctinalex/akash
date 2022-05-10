package provider

import (
	"context"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	tmrpc "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/ovrclk/akash/provider/event"
	"github.com/ovrclk/akash/provider/session"
	"github.com/ovrclk/akash/pubsub"
	"github.com/ovrclk/akash/testutil"
	dtypes "github.com/ovrclk/akash/x/deployment/types/v1beta2"
	"github.com/ovrclk/akash/x/escrow/types/v1beta2"
	mtypes "github.com/ovrclk/akash/x/market/types/v1beta2"
	ptypes "github.com/ovrclk/akash/x/provider/types/v1beta2"

	akashmock "github.com/ovrclk/akash/client/mocks"
	cosmosMock "github.com/ovrclk/akash/testutil/cosmos_mock"
)

type broadcaster struct {
	txCh chan sdk.Msg
}

type scaffold struct {
	testAddr  sdk.AccAddress
	testBus   pubsub.Bus
	ctx       context.Context
	cancel    context.CancelFunc
	qc        *cosmosMock.QueryClient
	aqc       *akashmock.QueryClient
	bc        *balanceChecker
	broadcast *broadcaster
}

func (bc *broadcaster) Broadcast(_ context.Context, msg ...sdk.Msg) error {
	bc.txCh <- msg[0]
	return nil
}

func balanceCheckerForTest(t *testing.T, balance int64, pollPeriod time.Duration) (*scaffold, *balanceChecker) {
	s := &scaffold{}

	s.ctx, s.cancel = context.WithCancel(context.Background())
	myLog := testutil.Logger(t)

	s.testAddr = testutil.AccAddress(t)

	queryClient := &cosmosMock.QueryClient{}
	aqc := akashmock.NewQueryClient(t)

	query := bankTypes.NewQueryBalanceRequest(s.testAddr, testutil.CoinDenom)
	coin := sdk.NewCoin(testutil.CoinDenom, sdk.NewInt(balance))
	result := &bankTypes.QueryBalanceResponse{
		Balance: &coin,
	}
	queryClient.On("Balance", mock.Anything, query).Return(result, nil)

	myProvider := &ptypes.Provider{
		Owner:      s.testAddr.String(),
		HostURI:    "http://test.localhost:7443",
		Attributes: nil,
	}
	mySession := session.New(myLog, nil, myProvider, -1)
	s.testBus = pubsub.NewBus()

	bc := newBalanceChecker(s.ctx, queryClient, aqc, s.testAddr, mySession, s.testBus, BalanceCheckerConfig{
		PollingPeriod:           pollPeriod,
		MinimumBalanceThreshold: 100,
		WithdrawalPeriod:        time.Hour * 24,
		LeaseFundCheckInterval:  time.Second * 30,
	})

	s.qc = queryClient
	s.bc = bc

	return s, bc
}

func leaseMonitorForTest(t *testing.T, balance int64, pollPeriod time.Duration) (*scaffold, *balanceChecker) {
	s := &scaffold{
		broadcast: &broadcaster{
			txCh: make(chan sdk.Msg, 1),
		},
	}

	s.ctx, s.cancel = context.WithCancel(context.Background())
	myLog := testutil.Logger(t)

	s.testAddr = testutil.AccAddress(t)

	aqc := akashmock.NewQueryClient(t)

	startedAt := time.Now()

	nodeSyncInfo := &tmrpc.SyncInfo{
		LatestBlockHeight: 1,
		CatchingUp:        true,
	}

	client := akashmock.NewClient(t)
	client.On("Tx").Return(s.broadcast)
	client.On("NodeSyncInfo", mock.Anything).Run(func(args mock.Arguments) {
		nodeSyncInfo.LatestBlockHeight = int64(time.Now().Sub(startedAt) / blockPeriod)
	}).Return(nodeSyncInfo, nil)

	queryClient := &cosmosMock.QueryClient{}

	deploymentResp := &dtypes.QueryDeploymentResponse{
		Deployment: dtypes.Deployment{
			DeploymentID: dtypes.DeploymentID{},
			State:        0,
			Version:      nil,
			CreatedAt:    0,
		},
		Groups: nil,
		EscrowAccount: v1beta2.Account{
			ID:          v1beta2.AccountID{},
			Owner:       "",
			State:       0,
			Balance:     sdk.NewDecCoin("uakt", sdk.NewInt(2000)),
			Transferred: sdk.DecCoin{},
			SettledAt:   1,
			Depositor:   "",
			Funds:       sdk.DecCoin{},
		},
	}

	aqc.On("Deployment", mock.Anything, mock.Anything).Return(deploymentResp, nil)

	leasesResp := &mtypes.QueryLeasesResponse{
		Leases: []mtypes.QueryLeaseResponse{
			{
				Lease: mtypes.Lease{
					Price: sdk.NewDecCoin("uakt", sdk.NewInt(1500)),
				},
			},
		},
	}
	aqc.On("Leases", mock.Anything, mock.Anything).Return(leasesResp, nil)

	myProvider := &ptypes.Provider{
		Owner:      s.testAddr.String(),
		HostURI:    "http://test.localhost:7443",
		Attributes: nil,
	}
	mySession := session.New(myLog, client, myProvider, -1)
	s.testBus = pubsub.NewBus()

	bc := newBalanceChecker(s.ctx, queryClient, aqc, s.testAddr, mySession, s.testBus, BalanceCheckerConfig{
		PollingPeriod:           pollPeriod,
		MinimumBalanceThreshold: 100,
		WithdrawalPeriod:        time.Hour * 24,
		LeaseFundCheckInterval:  time.Second * 30,
	})

	s.qc = queryClient
	s.bc = bc

	return s, bc
}

func TestBalanceCheckerChecksBalance(t *testing.T) {
	testScaffold, bc := balanceCheckerForTest(t, 9999999999999, time.Millisecond*100)
	defer testScaffold.testBus.Close()
	subscriber, err := testScaffold.testBus.Subscribe()
	require.NoError(t, err)

	firstEvent := make(chan pubsub.Event, 1)
	go func() {
		defer subscriber.Close()
		select {
		case ev := <-subscriber.Events():
			firstEvent <- ev
		case <-bc.lc.Done():

		}
	}()

	time.Sleep(bc.cfg.PollingPeriod * 3)
	testScaffold.cancel()
	<-bc.lc.Done()

	testScaffold.qc.AssertExpectations(t)

	// Make sure no event is sent
	select {
	case <-firstEvent:
		t.Fatal("should not have an event to read")
	default:

	}
}

func TestBalanceCheckerStartsWithdrawal(t *testing.T) {
	testScaffold, bc := balanceCheckerForTest(t, 1, time.Millisecond*100)
	defer testScaffold.testBus.Close()
	subscriber, err := testScaffold.testBus.Subscribe()
	require.NoError(t, err)

	firstEvent := make(chan pubsub.Event, 1)
	go func() {
		defer subscriber.Close()
		select {
		case ev := <-subscriber.Events():
			firstEvent <- ev
		case <-bc.lc.Done():
		}
	}()

	// Make sure the event is sent
	select {
	case ev := <-firstEvent:
		_, ok := ev.(event.LeaseWithdrawNow)
		require.True(t, ok)
	case <-time.After(15 * time.Second):
		t.Fatal("should have an event to read")
	}

	time.Sleep(bc.cfg.PollingPeriod)
	testScaffold.cancel()

	select {
	case <-bc.lc.Done():
	case <-time.After(15 * time.Second):
		t.Fatal("timed out waiting for completion")
	}

	testScaffold.qc.AssertExpectations(t)
}

func TestBalanceCheckerMonitorsFunds(t *testing.T) {
	testScaffold, bc := leaseMonitorForTest(t, 1, time.Second*100)
	defer testScaffold.testBus.Close()

	subscriber, err := testScaffold.testBus.Subscribe()
	require.NoError(t, err)
	defer subscriber.Close()

	lid := mtypes.LeaseID{
		Owner:    bc.ownAddr.String(),
		DSeq:     1,
		GSeq:     0,
		OSeq:     0,
		Provider: bc.ownAddr.String(),
	}

	time.Sleep(time.Second)

	err = testScaffold.testBus.Publish(event.LeaseAddFundsMonitor{LeaseID: lid})
	require.NoError(t, err)

	select {
	case msg := <-testScaffold.broadcast.txCh:
		switch msg.(type) {
		case *mtypes.MsgCloseBid:
		default:
			t.Errorf("received unexpected message")
		}
	case <-time.NewTimer(time.Second * 25).C:
		t.Errorf("has not received bid close message")
	case <-testScaffold.ctx.Done():
		t.Fail()
		return
	}

	testScaffold.qc.AssertExpectations(t)
}
