// Code generated by mockery 2.12.1. DO NOT EDIT.

package kubernetes_mocks

import (
	context "context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mock "github.com/stretchr/testify/mock"

	rest "k8s.io/client-go/rest"

	testing "testing"

	v1 "k8s.io/api/core/v1"

	v1beta1 "k8s.io/api/policy/v1beta1"
)

// PodExpansion is an autogenerated mock type for the PodExpansion type
type PodExpansion struct {
	mock.Mock
}

// Bind provides a mock function with given fields: ctx, binding, opts
func (_m *PodExpansion) Bind(ctx context.Context, binding *v1.Binding, opts metav1.CreateOptions) error {
	ret := _m.Called(ctx, binding, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.Binding, metav1.CreateOptions) error); ok {
		r0 = rf(ctx, binding, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Evict provides a mock function with given fields: ctx, eviction
func (_m *PodExpansion) Evict(ctx context.Context, eviction *v1beta1.Eviction) error {
	ret := _m.Called(ctx, eviction)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.Eviction) error); ok {
		r0 = rf(ctx, eviction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetLogs provides a mock function with given fields: name, opts
func (_m *PodExpansion) GetLogs(name string, opts *v1.PodLogOptions) *rest.Request {
	ret := _m.Called(name, opts)

	var r0 *rest.Request
	if rf, ok := ret.Get(0).(func(string, *v1.PodLogOptions) *rest.Request); ok {
		r0 = rf(name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Request)
		}
	}

	return r0
}

// ProxyGet provides a mock function with given fields: scheme, name, port, path, params
func (_m *PodExpansion) ProxyGet(scheme string, name string, port string, path string, params map[string]string) rest.ResponseWrapper {
	ret := _m.Called(scheme, name, port, path, params)

	var r0 rest.ResponseWrapper
	if rf, ok := ret.Get(0).(func(string, string, string, string, map[string]string) rest.ResponseWrapper); ok {
		r0 = rf(scheme, name, port, path, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.ResponseWrapper)
		}
	}

	return r0
}

// NewPodExpansion creates a new instance of PodExpansion. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPodExpansion(t testing.TB) *PodExpansion {
	mock := &PodExpansion{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
