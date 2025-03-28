// Code generated by mockery v2.51.1. DO NOT EDIT.

package io

import mock "github.com/stretchr/testify/mock"

// MockWriter is an autogenerated mock type for the Writer type
type MockWriter struct {
	mock.Mock
}

type MockWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWriter) EXPECT() *MockWriter_Expecter {
	return &MockWriter_Expecter{mock: &_m.Mock}
}

// Write provides a mock function with given fields: p
func (_m *MockWriter) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWriter_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type MockWriter_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - p []byte
func (_e *MockWriter_Expecter) Write(p interface{}) *MockWriter_Write_Call {
	return &MockWriter_Write_Call{Call: _e.mock.On("Write", p)}
}

func (_c *MockWriter_Write_Call) Run(run func(p []byte)) *MockWriter_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MockWriter_Write_Call) Return(n int, err error) *MockWriter_Write_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *MockWriter_Write_Call) RunAndReturn(run func([]byte) (int, error)) *MockWriter_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWriter creates a new instance of MockWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWriter {
	mock := &MockWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
