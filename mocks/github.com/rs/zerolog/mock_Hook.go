// Code generated by mockery v2.51.1. DO NOT EDIT.

package zerolog

import (
	zerolog "github.com/rs/zerolog"
	mock "github.com/stretchr/testify/mock"
)

// MockHook is an autogenerated mock type for the Hook type
type MockHook struct {
	mock.Mock
}

type MockHook_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHook) EXPECT() *MockHook_Expecter {
	return &MockHook_Expecter{mock: &_m.Mock}
}

// Run provides a mock function with given fields: e, level, message
func (_m *MockHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	_m.Called(e, level, message)
}

// MockHook_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockHook_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - e *zerolog.Event
//   - level zerolog.Level
//   - message string
func (_e *MockHook_Expecter) Run(e interface{}, level interface{}, message interface{}) *MockHook_Run_Call {
	return &MockHook_Run_Call{Call: _e.mock.On("Run", e, level, message)}
}

func (_c *MockHook_Run_Call) Run(run func(e *zerolog.Event, level zerolog.Level, message string)) *MockHook_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*zerolog.Event), args[1].(zerolog.Level), args[2].(string))
	})
	return _c
}

func (_c *MockHook_Run_Call) Return() *MockHook_Run_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockHook_Run_Call) RunAndReturn(run func(*zerolog.Event, zerolog.Level, string)) *MockHook_Run_Call {
	_c.Run(run)
	return _c
}

// NewMockHook creates a new instance of MockHook. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHook(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHook {
	mock := &MockHook{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
