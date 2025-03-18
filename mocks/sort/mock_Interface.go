// Code generated by mockery v2.51.1. DO NOT EDIT.

package sort

import mock "github.com/stretchr/testify/mock"

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

type MockInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface) EXPECT() *MockInterface_Expecter {
	return &MockInterface_Expecter{mock: &_m.Mock}
}

// Len provides a mock function with no fields
func (_m *MockInterface) Len() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Len")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockInterface_Len_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Len'
type MockInterface_Len_Call struct {
	*mock.Call
}

// Len is a helper method to define mock.On call
func (_e *MockInterface_Expecter) Len() *MockInterface_Len_Call {
	return &MockInterface_Len_Call{Call: _e.mock.On("Len")}
}

func (_c *MockInterface_Len_Call) Run(run func()) *MockInterface_Len_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_Len_Call) Return(_a0 int) *MockInterface_Len_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Len_Call) RunAndReturn(run func() int) *MockInterface_Len_Call {
	_c.Call.Return(run)
	return _c
}

// Less provides a mock function with given fields: i, j
func (_m *MockInterface) Less(i int, j int) bool {
	ret := _m.Called(i, j)

	if len(ret) == 0 {
		panic("no return value specified for Less")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, int) bool); ok {
		r0 = rf(i, j)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockInterface_Less_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Less'
type MockInterface_Less_Call struct {
	*mock.Call
}

// Less is a helper method to define mock.On call
//   - i int
//   - j int
func (_e *MockInterface_Expecter) Less(i interface{}, j interface{}) *MockInterface_Less_Call {
	return &MockInterface_Less_Call{Call: _e.mock.On("Less", i, j)}
}

func (_c *MockInterface_Less_Call) Run(run func(i int, j int)) *MockInterface_Less_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *MockInterface_Less_Call) Return(_a0 bool) *MockInterface_Less_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_Less_Call) RunAndReturn(run func(int, int) bool) *MockInterface_Less_Call {
	_c.Call.Return(run)
	return _c
}

// Swap provides a mock function with given fields: i, j
func (_m *MockInterface) Swap(i int, j int) {
	_m.Called(i, j)
}

// MockInterface_Swap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Swap'
type MockInterface_Swap_Call struct {
	*mock.Call
}

// Swap is a helper method to define mock.On call
//   - i int
//   - j int
func (_e *MockInterface_Expecter) Swap(i interface{}, j interface{}) *MockInterface_Swap_Call {
	return &MockInterface_Swap_Call{Call: _e.mock.On("Swap", i, j)}
}

func (_c *MockInterface_Swap_Call) Run(run func(i int, j int)) *MockInterface_Swap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *MockInterface_Swap_Call) Return() *MockInterface_Swap_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockInterface_Swap_Call) RunAndReturn(run func(int, int)) *MockInterface_Swap_Call {
	_c.Run(run)
	return _c
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
