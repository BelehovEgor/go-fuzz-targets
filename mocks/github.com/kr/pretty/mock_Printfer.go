// Code generated by mockery v2.51.1. DO NOT EDIT.

package pretty

import mock "github.com/stretchr/testify/mock"

// MockPrintfer is an autogenerated mock type for the Printfer type
type MockPrintfer struct {
	mock.Mock
}

type MockPrintfer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPrintfer) EXPECT() *MockPrintfer_Expecter {
	return &MockPrintfer_Expecter{mock: &_m.Mock}
}

// Printf provides a mock function with given fields: format, a
func (_m *MockPrintfer) Printf(format string, a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// MockPrintfer_Printf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Printf'
type MockPrintfer_Printf_Call struct {
	*mock.Call
}

// Printf is a helper method to define mock.On call
//   - format string
//   - a ...interface{}
func (_e *MockPrintfer_Expecter) Printf(format interface{}, a ...interface{}) *MockPrintfer_Printf_Call {
	return &MockPrintfer_Printf_Call{Call: _e.mock.On("Printf",
		append([]interface{}{format}, a...)...)}
}

func (_c *MockPrintfer_Printf_Call) Run(run func(format string, a ...interface{})) *MockPrintfer_Printf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockPrintfer_Printf_Call) Return() *MockPrintfer_Printf_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPrintfer_Printf_Call) RunAndReturn(run func(string, ...interface{})) *MockPrintfer_Printf_Call {
	_c.Run(run)
	return _c
}

// NewMockPrintfer creates a new instance of MockPrintfer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPrintfer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPrintfer {
	mock := &MockPrintfer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
