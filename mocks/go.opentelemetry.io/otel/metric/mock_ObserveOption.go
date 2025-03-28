// Code generated by mockery v2.51.1. DO NOT EDIT.

package metric

import (
	mock "github.com/stretchr/testify/mock"
	metric "go.opentelemetry.io/otel/metric"
)

// MockObserveOption is an autogenerated mock type for the ObserveOption type
type MockObserveOption struct {
	mock.Mock
}

type MockObserveOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockObserveOption) EXPECT() *MockObserveOption_Expecter {
	return &MockObserveOption_Expecter{mock: &_m.Mock}
}

// applyObserve provides a mock function with given fields: _a0
func (_m *MockObserveOption) applyObserve(_a0 metric.ObserveConfig) metric.ObserveConfig {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for applyObserve")
	}

	var r0 metric.ObserveConfig
	if rf, ok := ret.Get(0).(func(metric.ObserveConfig) metric.ObserveConfig); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(metric.ObserveConfig)
	}

	return r0
}

// MockObserveOption_applyObserve_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'applyObserve'
type MockObserveOption_applyObserve_Call struct {
	*mock.Call
}

// applyObserve is a helper method to define mock.On call
//   - _a0 metric.ObserveConfig
func (_e *MockObserveOption_Expecter) applyObserve(_a0 interface{}) *MockObserveOption_applyObserve_Call {
	return &MockObserveOption_applyObserve_Call{Call: _e.mock.On("applyObserve", _a0)}
}

func (_c *MockObserveOption_applyObserve_Call) Run(run func(_a0 metric.ObserveConfig)) *MockObserveOption_applyObserve_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metric.ObserveConfig))
	})
	return _c
}

func (_c *MockObserveOption_applyObserve_Call) Return(_a0 metric.ObserveConfig) *MockObserveOption_applyObserve_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockObserveOption_applyObserve_Call) RunAndReturn(run func(metric.ObserveConfig) metric.ObserveConfig) *MockObserveOption_applyObserve_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockObserveOption creates a new instance of MockObserveOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockObserveOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockObserveOption {
	mock := &MockObserveOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
