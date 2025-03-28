// Code generated by mockery v2.51.1. DO NOT EDIT.

package trace

import (
	mock "github.com/stretchr/testify/mock"
	trace "go.opentelemetry.io/otel/trace"
)

// MockTracerOption is an autogenerated mock type for the TracerOption type
type MockTracerOption struct {
	mock.Mock
}

type MockTracerOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTracerOption) EXPECT() *MockTracerOption_Expecter {
	return &MockTracerOption_Expecter{mock: &_m.Mock}
}

// apply provides a mock function with given fields: _a0
func (_m *MockTracerOption) apply(_a0 trace.TracerConfig) trace.TracerConfig {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for apply")
	}

	var r0 trace.TracerConfig
	if rf, ok := ret.Get(0).(func(trace.TracerConfig) trace.TracerConfig); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(trace.TracerConfig)
	}

	return r0
}

// MockTracerOption_apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'apply'
type MockTracerOption_apply_Call struct {
	*mock.Call
}

// apply is a helper method to define mock.On call
//   - _a0 trace.TracerConfig
func (_e *MockTracerOption_Expecter) apply(_a0 interface{}) *MockTracerOption_apply_Call {
	return &MockTracerOption_apply_Call{Call: _e.mock.On("apply", _a0)}
}

func (_c *MockTracerOption_apply_Call) Run(run func(_a0 trace.TracerConfig)) *MockTracerOption_apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(trace.TracerConfig))
	})
	return _c
}

func (_c *MockTracerOption_apply_Call) Return(_a0 trace.TracerConfig) *MockTracerOption_apply_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTracerOption_apply_Call) RunAndReturn(run func(trace.TracerConfig) trace.TracerConfig) *MockTracerOption_apply_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTracerOption creates a new instance of MockTracerOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTracerOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTracerOption {
	mock := &MockTracerOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
