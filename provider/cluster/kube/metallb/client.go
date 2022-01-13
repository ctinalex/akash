package metallb

import (
	"context"
	"errors"
	"fmt"
	"github.com/ovrclk/akash/manifest"
	"github.com/ovrclk/akash/provider/cluster/kube/builder"
	"github.com/ovrclk/akash/provider/cluster/kube/client_common"
	ctypes "github.com/ovrclk/akash/provider/cluster/types"
	clusterutil "github.com/ovrclk/akash/provider/cluster/util"
	mtypes "github.com/ovrclk/akash/x/market/types/v1beta2"
	"github.com/tendermint/tendermint/libs/log"
	corev1 "k8s.io/api/core/v1"
	kubeErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/pager"
	"k8s.io/client-go/util/flowcontrol"
	"math"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/common/expfmt"

	kubeclienterrors "github.com/ovrclk/akash/provider/cluster/kube/errors"
)

const (
	akashServiceTarget = "akash.network/service-target"
	akashMetalLB = "metal-lb"
	metalLbAllowSharedIp = "metallb.universe.tf/allow-shared-ip"
)


type Client interface {
	GetIPAddressUsage(ctx context.Context) (uint, uint, error)
	GetIPAddressStatusForLease(ctx context.Context, leaseID mtypes.LeaseID) ([]ctypes.IPLeaseState, error)

	CreateIPPassthrough(ctx context.Context, lID mtypes.LeaseID, directive ctypes.ClusterIPPassthroughDirective) error
	PurgeIPPassthrough(ctx context.Context, lID mtypes.LeaseID, directive ctypes.ClusterIPPassthroughDirective) error
	GetIPPassthroughs(ctx context.Context) ([]ctypes.IPPassthrough, error)

	Stop()
}

type client struct {
	kube kubernetes.Interface
	httpClient *http.Client

	log               log.Logger

	sda clusterutil.ServiceDiscoveryAgent
}


func (c *client) String() string {
	return fmt.Sprintf("metal LB client %p", c)
}

const (
	metricsPath = "/metrics"
	metricsTimeout = 10 * time.Second

	poolName = "default"

	serviceHostName = "controller.metallb-system.svc.cluster.local"

	metricNameAddrInUse = "metallb_allocator_addresses_in_use_total"
	metricNameAddrTotal = "metallb_allocator_addresses_total"
)

var (
	errMetalLB = errors.New("metal lb error")
)


func NewClient(configPath string, logger log.Logger) (Client, error){
	config, err := client_common.OpenKubeConfig(configPath, logger)
	if err != nil {
		return nil, fmt.Errorf("%w: creating kubernetes client", err)
	}
	config.RateLimiter = flowcontrol.NewFakeAlwaysRateLimiter()

	kc, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, fmt.Errorf("%w: creating kubernetes client", err)
	}

	dialer := net.Dialer{
		Timeout:       metricsTimeout,
		Deadline:      time.Time{},
		LocalAddr:     nil,
		FallbackDelay: 0,
		KeepAlive:     0,
		Resolver:      nil,
		Control:       nil,
	}

	transport := &http.Transport{
		Proxy:                  nil,
		DialContext:            dialer.DialContext,
		DialTLSContext:         nil,
		TLSClientConfig:        nil,
		TLSHandshakeTimeout:    0,
		DisableKeepAlives:      false,
		DisableCompression:     true,
		MaxIdleConns:           1,
		MaxIdleConnsPerHost:    1,
		MaxConnsPerHost:        1,
		IdleConnTimeout:        0,
		ResponseHeaderTimeout:  metricsTimeout,
		ExpectContinueTimeout:  metricsTimeout,
		TLSNextProto:           nil,
		ProxyConnectHeader:     nil,
		GetProxyConnectHeader:  nil,
		MaxResponseHeaderBytes: 0,
		WriteBufferSize:        0,
		ReadBufferSize:         0,
		ForceAttemptHTTP2:      false,
	}


	sda := clusterutil.NewServiceDiscoveryAgent(logger, "monitoring", "controller", "metallb-system", "TCP")


	return &client{
		sda: sda,
		kube: kc,
		httpClient: &http.Client{
			Transport:     transport,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       metricsTimeout,
		},

		log: logger.With("client","metallb"),
	}, nil

}

func (c *client) Stop() {
	c.sda.Stop()
}

/*
can get stuff like this to access metal lb metrics
   75  nslookup -type=SRV _monitoring._tcp.

  102  curl -I controller.metallb-system.svc.cluster.local:7472/metrics

 */


func (c *client) GetIPAddressUsage(ctx context.Context) (uint, uint,  error) {
	metalAddr, err := c.sda.GetAddress(ctx)
	if err != nil {
		return math.MaxUint32,math.MaxUint32, err
	}
	metricsURL := fmt.Sprintf("http://%s:%d%s", metalAddr.Target, metalAddr.Port, metricsPath)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, metricsURL, nil)
	if err != nil {
		return math.MaxUint32,math.MaxUint32, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return math.MaxUint32,math.MaxUint32, err
	}

	if response.StatusCode != http.StatusOK {
		return math.MaxUint32,math.MaxUint32, fmt.Errorf("%w: response status %d", errMetalLB, response.StatusCode)
	}


	var parser expfmt.TextParser
	mf, err := parser.TextToMetricFamilies(response.Body)
	if err != nil {
		return math.MaxUint32,math.MaxUint32, err
	}

	/**
	  Loooking for the following metrics
	    metallb_allocator_addresses_in_use_total{pool="default"} 0
	    metallb_allocator_addresses_total{pool="default"} 100
	 */

	available := uint(0)
	setAvailable := false
	inUse := uint(0)
	setInUse := false
	for _, entry := range mf {
		if setInUse && setAvailable {
			break
		}
		var target *uint
		var setTarget *bool
		if entry.GetName() == metricNameAddrInUse   {
			target = &inUse
			setTarget = &setInUse
		} else if entry.GetName() == metricNameAddrTotal {
			target = &available
			setTarget = &setAvailable
		} else {
			continue
		}

		metric := entry.GetMetric()
		searchLoop:
		for _, metricEntry := range metric{
			gauge := metricEntry.GetGauge()
			if gauge == nil {
				continue
			}
			for _, labelEntry := range metricEntry.Label {
				if labelEntry.GetName()  != "pool" {
					continue
				}

				if labelEntry.GetValue() != poolName {
					continue
				}

				*target = uint(*gauge.Value)
				*setTarget = true
				break searchLoop
			}
		}
	}

	if !setInUse || !setAvailable {
		return math.MaxUint32, math.MaxUint32, fmt.Errorf("%w: data not found in metrics response", errMetalLB)
	}

	return inUse, available, nil
}

type ipLeaseState struct {
	leaseID mtypes.LeaseID
	ip string
	serviceName string
	externalPort uint32
	port uint32
	sharingKey string
	protocol manifest.ServiceProtocol

}

func (ipls ipLeaseState) GetLeaseID() mtypes.LeaseID {
	return ipls.leaseID
}
func (ipls ipLeaseState) GetIP() string {
	return ipls.ip
}
func (ipls ipLeaseState) GetServiceName() string {
	return ipls.serviceName
}
func (ipls ipLeaseState)GetExternalPort() uint32 {
	return ipls.externalPort
}
func (ipls ipLeaseState)GetPort() uint32 {
	return ipls.port
}
func (ipls ipLeaseState)GetSharingKey() string {
	return ipls.sharingKey
}
func (ipls ipLeaseState)GetProtocol() manifest.ServiceProtocol {
	return ipls.protocol
}

func (c *client) GetIPAddressStatusForLease(ctx context.Context, leaseID mtypes.LeaseID) ([]ctypes.IPLeaseState, error){
	ns := builder.LidNS(leaseID)
	servicePager := pager.New(func(ctx context.Context, opts metav1.ListOptions) (runtime.Object, error){
		return c.kube.CoreV1().Services(ns).List(ctx, opts)
	})

	labelSelector := &strings.Builder{}

	_, err := fmt.Fprintf(labelSelector, "%s=true", builder.AkashManagedLabelName)
	if err != nil {
		return nil, err
	}
	_, err = fmt.Fprintf(labelSelector, ",%s=%s", akashServiceTarget, akashMetalLB)
	if err != nil {
		return nil, err
	}

	result := make([]ctypes.IPLeaseState, 0)
	err = servicePager.EachListItem(ctx, metav1.ListOptions{
		LabelSelector: labelSelector.String(),
	},
		func(obj runtime.Object) error {
			service := obj.(*corev1.Service)

			loadBalancerIngress := service.Status.LoadBalancer.Ingress
			// Logs something like this : │ load balancer status                         cmp=provider client=kube service=web-ip-80-tcp lb-ingress="[{IP:24.0.0.1 Hostname: Ports:[]}]"
			c.log.Debug("load balancer status", "service", service.ObjectMeta.Name, "lb-ingress", loadBalancerIngress)

			// There is no mechanism that would assign more than one IP to a single service entry
			if len(loadBalancerIngress) != 1 {
				// TODO return an error indicating something is invalid
				panic("invalid load balancer ingress")
			}

			ingress := loadBalancerIngress[0]

			if len(service.Spec.Ports) != 1 {
				panic("invalid port specs")
			}
			port := service.Spec.Ports[0]

			// TODO - make this some sort of utility method
			var proto manifest.ServiceProtocol
			switch(port.Protocol) {

			case corev1.ProtocolTCP:
				proto = manifest.TCP
			case corev1.ProtocolUDP:
				proto = manifest.UDP
			default:
				panic("unknown proto from kube: " + string(port.Protocol))
			}

			selectedServiceName := service.Spec.Selector[builder.AkashManifestServiceLabelName]

			// Note: don't care about node port here, even if it is assigned
			// Note: service.Name is a procedurally generated thing that doesn't mean anything to the end user
			result = append(result, ipLeaseState{
				leaseID:      leaseID,
				ip:           ingress.IP,
				serviceName:  selectedServiceName,
				externalPort: uint32(port.Port),
				port:         uint32(port.TargetPort.IntValue()),
				sharingKey:   service.ObjectMeta.Annotations[metalLbAllowSharedIp],
				protocol:     proto,
			})

			return nil
		})
	if err != nil {
		return nil, err
	}

	return result, nil
}



func (c *client) PurgeIPPassthrough(ctx context.Context, leaseID mtypes.LeaseID, directive ctypes.ClusterIPPassthroughDirective) error {
	ns := builder.LidNS(leaseID)
	portName := createIPPassthroughResourceName(directive)

	err := c.kube.CoreV1().Services(ns).Delete(ctx, portName, metav1.DeleteOptions{})

	if err != nil && kubeErrors.IsNotFound(err) {
		return nil
	}

	return err
}

func createIPPassthroughResourceName(directive ctypes.ClusterIPPassthroughDirective) string{
	return strings.ToLower(fmt.Sprintf("%s-ip-%d-%v", directive.ServiceName, directive.ExternalPort, directive.Protocol))
}

func (c *client) CreateIPPassthrough(ctx context.Context, leaseID mtypes.LeaseID, directive ctypes.ClusterIPPassthroughDirective) error {
	var proto corev1.Protocol

	switch(directive.Protocol) {
	case manifest.TCP:
		proto = corev1.ProtocolTCP
	case manifest.UDP:
		proto = corev1.ProtocolUDP
	default:
		return fmt.Errorf("%w unknown protocol %v", kubeclienterrors.ErrInternalError, directive.Protocol)
	}

	ns := builder.LidNS(leaseID)
	portName := createIPPassthroughResourceName(directive)

	foundEntry, err := c.kube.CoreV1().Services(ns).Get(ctx, portName, metav1.GetOptions{})

	exists := true
	if err != nil {
		if kubeErrors.IsNotFound(err) {
			exists = false
		} else {
			return err
		}
	}

	labels := make(map[string]string)
	builder.AppendLeaseLabels(leaseID, labels)
	labels[builder.AkashManagedLabelName] = "true"
	labels[akashServiceTarget] = akashMetalLB

	selector := map[string]string {
		builder.AkashManagedLabelName: "true",
		builder.AkashManifestServiceLabelName: directive.ServiceName,
	}
	// TODO - specify metallb.universe.tf/address-pool annotation if configured to do so only that pool is used at any time
	annotations := map[string]string {
		metalLbAllowSharedIp: directive.SharingKey,
	}

	port := corev1.ServicePort{
		Name:        portName,
		Protocol:    proto,
		Port:        int32(directive.ExternalPort),
		TargetPort:  intstr.FromInt(int(directive.Port)),
	}

	svc := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:                       portName,
			Labels: labels,
			Annotations: annotations,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				port,
			},
			Selector:                      selector,
			Type:                          corev1.ServiceTypeLoadBalancer,
		},
		Status: corev1.ServiceStatus{},
	}

	c.log.Debug("creating metal-lb service",
		"service", directive.ServiceName,
		"port", directive.Port,
		"external-port", directive.ExternalPort,
		"sharing-key", directive.SharingKey,
		"exists", exists)
	if exists {
		svc.ResourceVersion = foundEntry.ResourceVersion
		_, err = c.kube.CoreV1().Services(ns).Update(ctx, &svc, metav1.UpdateOptions{})
	} else {
		_, err = c.kube.CoreV1().Services(ns).Create(ctx, &svc, metav1.CreateOptions{})
	}

	if err != nil {
		return err
	}

	return nil
}


func (c *client) GetIPPassthroughs(ctx context.Context) ([]ctypes.IPPassthrough, error) {
	servicePager := pager.New(func(ctx context.Context, opts metav1.ListOptions) (runtime.Object, error){
		return c.kube.CoreV1().Services(metav1.NamespaceAll).List(ctx, opts)
	})

	labelSelector := &strings.Builder{}

	_, err := fmt.Fprintf(labelSelector, "%s=true", builder.AkashManagedLabelName)
	if err != nil {
		return nil, err
	}
	_, err = fmt.Fprintf(labelSelector, ",%s=%s", akashServiceTarget, akashMetalLB)
	if err != nil {
		return nil, err
	}

	result := make([]ctypes.IPPassthrough, 0)
	err = servicePager.EachListItem(ctx,
		metav1.ListOptions{
			LabelSelector: labelSelector.String(),
		},
		func(obj runtime.Object) error {
			service := obj.(*corev1.Service)

			_, hasOwner := service.ObjectMeta.Labels[builder.AkashLeaseOwnerLabelName]
			if !hasOwner {
				// Not a service related to a running deployment, so probably internal services
				return nil
			}

			if service.Spec.Type != corev1.ServiceTypeLoadBalancer {
				return fmt.Errorf("resource %q wrong type in service definition %v", service.ObjectMeta.Name, service.Spec.Type)
			}


			ports := service.Spec.Ports
			const expectedNumberOfPorts = 1
			if len(ports) != expectedNumberOfPorts {
				return fmt.Errorf("resource %q  wrong number of ports in load balancer service definition. expected %d, got %d", service.ObjectMeta.Name, expectedNumberOfPorts, len(ports))
			}

			portDefn := ports[0]
			proto := portDefn.Protocol
			port := portDefn.Port

			// TODO - use a utlity method here rather than a cast
			mproto, err := manifest.ParseServiceProtocol(string(proto))
			if err != nil {
				return err // TODO include resource name
			}

			leaseID, err := client_common.RecoverLeaseIdFromLabels(service.Labels)
			if err != nil {
				return err // TODO include resource name
			}


			serviceSelector := service.Spec.Selector
			serviceName := serviceSelector[builder.AkashManifestServiceLabelName]
			if len(serviceName) == 0 {
				return fmt.Errorf("service name cannot be empty")
			}

			sharingKey := service.ObjectMeta.Annotations[metalLbAllowSharedIp]

			v := ipPassthrough{
				lID:          leaseID,
				serviceName:  serviceName,
				externalPort: uint32(port),
				sharingKey:   sharingKey,
				protocol:     mproto,
			}

			result = append(result, v)
			return nil
		})


	return result, err
}

type ipPassthrough struct {
	lID mtypes.LeaseID
	serviceName string
	port uint32
	externalPort uint32
	sharingKey string
	protocol manifest.ServiceProtocol
}

func (ev ipPassthrough) GetLeaseID() mtypes.LeaseID {
	return ev.lID
}

func (ev ipPassthrough) GetServiceName() string {
	return ev.serviceName
}

func (ev ipPassthrough) GetPort() uint32 {
	return ev.port
}

func (ev ipPassthrough) GetExternalPort() uint32 {
	return ev.externalPort
}

func (ev ipPassthrough) GetSharingKey() string {
	return ev.sharingKey
}

func (ev ipPassthrough) GetProtocol() manifest.ServiceProtocol {
	return ev.protocol
}