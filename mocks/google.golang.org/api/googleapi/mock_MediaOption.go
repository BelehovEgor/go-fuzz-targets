// Code generated by mockery v2.51.1. DO NOT EDIT.

package googleapi

import (
	mock "github.com/stretchr/testify/mock"
	googleapi "google.golang.org/api/googleapi"
)

// MockMediaOption is an autogenerated mock type for the MediaOption type
type MockMediaOption struct {
	mock.Mock
}

type MockMediaOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMediaOption) EXPECT() *MockMediaOption_Expecter {
	return &MockMediaOption_Expecter{mock: &_m.Mock}
}

// setOptions provides a mock function with given fields: o
func (_m *MockMediaOption) setOptions(o *googleapi.MediaOptions) {
	_m.Called(o)
}

// MockMediaOption_setOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'setOptions'
type MockMediaOption_setOptions_Call struct {
	*mock.Call
}

// setOptions is a helper method to define mock.On call
//   - o *googleapi.MediaOptions
func (_e *MockMediaOption_Expecter) setOptions(o interface{}) *MockMediaOption_setOptions_Call {
	return &MockMediaOption_setOptions_Call{Call: _e.mock.On("setOptions", o)}
}

func (_c *MockMediaOption_setOptions_Call) Run(run func(o *googleapi.MediaOptions)) *MockMediaOption_setOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*googleapi.MediaOptions))
	})
	return _c
}

func (_c *MockMediaOption_setOptions_Call) Return() *MockMediaOption_setOptions_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMediaOption_setOptions_Call) RunAndReturn(run func(*googleapi.MediaOptions)) *MockMediaOption_setOptions_Call {
	_c.Run(run)
	return _c
}

// NewMockMediaOption creates a new instance of MockMediaOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMediaOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMediaOption {
	mock := &MockMediaOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
