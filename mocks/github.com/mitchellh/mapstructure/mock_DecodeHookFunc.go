// Code generated by mockery v2.51.1. DO NOT EDIT.

package mapstructure

import mock "github.com/stretchr/testify/mock"

// MockDecodeHookFunc is an autogenerated mock type for the DecodeHookFunc type
type MockDecodeHookFunc struct {
	mock.Mock
}

type MockDecodeHookFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDecodeHookFunc) EXPECT() *MockDecodeHookFunc_Expecter {
	return &MockDecodeHookFunc_Expecter{mock: &_m.Mock}
}

// NewMockDecodeHookFunc creates a new instance of MockDecodeHookFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDecodeHookFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDecodeHookFunc {
	mock := &MockDecodeHookFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
