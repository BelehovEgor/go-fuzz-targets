// Code generated by mockery v2.51.1. DO NOT EDIT.

package metric

import (
	mock "github.com/stretchr/testify/mock"
	metric "go.opentelemetry.io/otel/metric"
)

// MockFloat64GaugeOption is an autogenerated mock type for the Float64GaugeOption type
type MockFloat64GaugeOption struct {
	mock.Mock
}

type MockFloat64GaugeOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFloat64GaugeOption) EXPECT() *MockFloat64GaugeOption_Expecter {
	return &MockFloat64GaugeOption_Expecter{mock: &_m.Mock}
}

// applyFloat64Gauge provides a mock function with given fields: _a0
func (_m *MockFloat64GaugeOption) applyFloat64Gauge(_a0 metric.Float64GaugeConfig) metric.Float64GaugeConfig {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for applyFloat64Gauge")
	}

	var r0 metric.Float64GaugeConfig
	if rf, ok := ret.Get(0).(func(metric.Float64GaugeConfig) metric.Float64GaugeConfig); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(metric.Float64GaugeConfig)
	}

	return r0
}

// MockFloat64GaugeOption_applyFloat64Gauge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'applyFloat64Gauge'
type MockFloat64GaugeOption_applyFloat64Gauge_Call struct {
	*mock.Call
}

// applyFloat64Gauge is a helper method to define mock.On call
//   - _a0 metric.Float64GaugeConfig
func (_e *MockFloat64GaugeOption_Expecter) applyFloat64Gauge(_a0 interface{}) *MockFloat64GaugeOption_applyFloat64Gauge_Call {
	return &MockFloat64GaugeOption_applyFloat64Gauge_Call{Call: _e.mock.On("applyFloat64Gauge", _a0)}
}

func (_c *MockFloat64GaugeOption_applyFloat64Gauge_Call) Run(run func(_a0 metric.Float64GaugeConfig)) *MockFloat64GaugeOption_applyFloat64Gauge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metric.Float64GaugeConfig))
	})
	return _c
}

func (_c *MockFloat64GaugeOption_applyFloat64Gauge_Call) Return(_a0 metric.Float64GaugeConfig) *MockFloat64GaugeOption_applyFloat64Gauge_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFloat64GaugeOption_applyFloat64Gauge_Call) RunAndReturn(run func(metric.Float64GaugeConfig) metric.Float64GaugeConfig) *MockFloat64GaugeOption_applyFloat64Gauge_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFloat64GaugeOption creates a new instance of MockFloat64GaugeOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFloat64GaugeOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFloat64GaugeOption {
	mock := &MockFloat64GaugeOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
