// Code generated by mockery v2.51.1. DO NOT EDIT.

package viper

import (
	viper "github.com/spf13/viper"
	mock "github.com/stretchr/testify/mock"
)

// MockFlagValueSet is an autogenerated mock type for the FlagValueSet type
type MockFlagValueSet struct {
	mock.Mock
}

type MockFlagValueSet_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFlagValueSet) EXPECT() *MockFlagValueSet_Expecter {
	return &MockFlagValueSet_Expecter{mock: &_m.Mock}
}

// VisitAll provides a mock function with given fields: fn
func (_m *MockFlagValueSet) VisitAll(fn func(viper.FlagValue)) {
	_m.Called(fn)
}

// MockFlagValueSet_VisitAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VisitAll'
type MockFlagValueSet_VisitAll_Call struct {
	*mock.Call
}

// VisitAll is a helper method to define mock.On call
//   - fn func(viper.FlagValue)
func (_e *MockFlagValueSet_Expecter) VisitAll(fn interface{}) *MockFlagValueSet_VisitAll_Call {
	return &MockFlagValueSet_VisitAll_Call{Call: _e.mock.On("VisitAll", fn)}
}

func (_c *MockFlagValueSet_VisitAll_Call) Run(run func(fn func(viper.FlagValue))) *MockFlagValueSet_VisitAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(viper.FlagValue)))
	})
	return _c
}

func (_c *MockFlagValueSet_VisitAll_Call) Return() *MockFlagValueSet_VisitAll_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockFlagValueSet_VisitAll_Call) RunAndReturn(run func(func(viper.FlagValue))) *MockFlagValueSet_VisitAll_Call {
	_c.Run(run)
	return _c
}

// NewMockFlagValueSet creates a new instance of MockFlagValueSet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFlagValueSet(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFlagValueSet {
	mock := &MockFlagValueSet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
