// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1beta1 "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
)

// MetricsV1beta1Interface is an autogenerated mock type for the MetricsV1beta1Interface type
type MetricsV1beta1Interface struct {
	mock.Mock
}

// NodeMetricses provides a mock function with given fields:
func (_m *MetricsV1beta1Interface) NodeMetricses() v1beta1.NodeMetricsInterface {
	ret := _m.Called()

	var r0 v1beta1.NodeMetricsInterface
	if rf, ok := ret.Get(0).(func() v1beta1.NodeMetricsInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.NodeMetricsInterface)
		}
	}

	return r0
}

// PodMetricses provides a mock function with given fields: namespace
func (_m *MetricsV1beta1Interface) PodMetricses(namespace string) v1beta1.PodMetricsInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.PodMetricsInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.PodMetricsInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.PodMetricsInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *MetricsV1beta1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}
