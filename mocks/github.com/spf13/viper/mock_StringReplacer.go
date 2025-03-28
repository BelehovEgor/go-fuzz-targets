// Code generated by mockery v2.51.1. DO NOT EDIT.

package viper

import mock "github.com/stretchr/testify/mock"

// MockStringReplacer is an autogenerated mock type for the StringReplacer type
type MockStringReplacer struct {
	mock.Mock
}

type MockStringReplacer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStringReplacer) EXPECT() *MockStringReplacer_Expecter {
	return &MockStringReplacer_Expecter{mock: &_m.Mock}
}

// Replace provides a mock function with given fields: s
func (_m *MockStringReplacer) Replace(s string) string {
	ret := _m.Called(s)

	if len(ret) == 0 {
		panic("no return value specified for Replace")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockStringReplacer_Replace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Replace'
type MockStringReplacer_Replace_Call struct {
	*mock.Call
}

// Replace is a helper method to define mock.On call
//   - s string
func (_e *MockStringReplacer_Expecter) Replace(s interface{}) *MockStringReplacer_Replace_Call {
	return &MockStringReplacer_Replace_Call{Call: _e.mock.On("Replace", s)}
}

func (_c *MockStringReplacer_Replace_Call) Run(run func(s string)) *MockStringReplacer_Replace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStringReplacer_Replace_Call) Return(_a0 string) *MockStringReplacer_Replace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStringReplacer_Replace_Call) RunAndReturn(run func(string) string) *MockStringReplacer_Replace_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStringReplacer creates a new instance of MockStringReplacer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStringReplacer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStringReplacer {
	mock := &MockStringReplacer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
