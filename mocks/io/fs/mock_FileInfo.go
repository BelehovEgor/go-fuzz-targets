// Code generated by mockery v2.51.1. DO NOT EDIT.

package fs

import (
	fs "io/fs"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockFileInfo is an autogenerated mock type for the FileInfo type
type MockFileInfo struct {
	mock.Mock
}

type MockFileInfo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFileInfo) EXPECT() *MockFileInfo_Expecter {
	return &MockFileInfo_Expecter{mock: &_m.Mock}
}

// IsDir provides a mock function with no fields
func (_m *MockFileInfo) IsDir() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsDir")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockFileInfo_IsDir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsDir'
type MockFileInfo_IsDir_Call struct {
	*mock.Call
}

// IsDir is a helper method to define mock.On call
func (_e *MockFileInfo_Expecter) IsDir() *MockFileInfo_IsDir_Call {
	return &MockFileInfo_IsDir_Call{Call: _e.mock.On("IsDir")}
}

func (_c *MockFileInfo_IsDir_Call) Run(run func()) *MockFileInfo_IsDir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileInfo_IsDir_Call) Return(_a0 bool) *MockFileInfo_IsDir_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileInfo_IsDir_Call) RunAndReturn(run func() bool) *MockFileInfo_IsDir_Call {
	_c.Call.Return(run)
	return _c
}

// ModTime provides a mock function with no fields
func (_m *MockFileInfo) ModTime() time.Time {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ModTime")
	}

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// MockFileInfo_ModTime_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ModTime'
type MockFileInfo_ModTime_Call struct {
	*mock.Call
}

// ModTime is a helper method to define mock.On call
func (_e *MockFileInfo_Expecter) ModTime() *MockFileInfo_ModTime_Call {
	return &MockFileInfo_ModTime_Call{Call: _e.mock.On("ModTime")}
}

func (_c *MockFileInfo_ModTime_Call) Run(run func()) *MockFileInfo_ModTime_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileInfo_ModTime_Call) Return(_a0 time.Time) *MockFileInfo_ModTime_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileInfo_ModTime_Call) RunAndReturn(run func() time.Time) *MockFileInfo_ModTime_Call {
	_c.Call.Return(run)
	return _c
}

// Mode provides a mock function with no fields
func (_m *MockFileInfo) Mode() fs.FileMode {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Mode")
	}

	var r0 fs.FileMode
	if rf, ok := ret.Get(0).(func() fs.FileMode); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(fs.FileMode)
	}

	return r0
}

// MockFileInfo_Mode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Mode'
type MockFileInfo_Mode_Call struct {
	*mock.Call
}

// Mode is a helper method to define mock.On call
func (_e *MockFileInfo_Expecter) Mode() *MockFileInfo_Mode_Call {
	return &MockFileInfo_Mode_Call{Call: _e.mock.On("Mode")}
}

func (_c *MockFileInfo_Mode_Call) Run(run func()) *MockFileInfo_Mode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileInfo_Mode_Call) Return(_a0 fs.FileMode) *MockFileInfo_Mode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileInfo_Mode_Call) RunAndReturn(run func() fs.FileMode) *MockFileInfo_Mode_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with no fields
func (_m *MockFileInfo) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockFileInfo_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockFileInfo_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockFileInfo_Expecter) Name() *MockFileInfo_Name_Call {
	return &MockFileInfo_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockFileInfo_Name_Call) Run(run func()) *MockFileInfo_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileInfo_Name_Call) Return(_a0 string) *MockFileInfo_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileInfo_Name_Call) RunAndReturn(run func() string) *MockFileInfo_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Size provides a mock function with no fields
func (_m *MockFileInfo) Size() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Size")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// MockFileInfo_Size_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Size'
type MockFileInfo_Size_Call struct {
	*mock.Call
}

// Size is a helper method to define mock.On call
func (_e *MockFileInfo_Expecter) Size() *MockFileInfo_Size_Call {
	return &MockFileInfo_Size_Call{Call: _e.mock.On("Size")}
}

func (_c *MockFileInfo_Size_Call) Run(run func()) *MockFileInfo_Size_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileInfo_Size_Call) Return(_a0 int64) *MockFileInfo_Size_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileInfo_Size_Call) RunAndReturn(run func() int64) *MockFileInfo_Size_Call {
	_c.Call.Return(run)
	return _c
}

// Sys provides a mock function with no fields
func (_m *MockFileInfo) Sys() interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Sys")
	}

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// MockFileInfo_Sys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sys'
type MockFileInfo_Sys_Call struct {
	*mock.Call
}

// Sys is a helper method to define mock.On call
func (_e *MockFileInfo_Expecter) Sys() *MockFileInfo_Sys_Call {
	return &MockFileInfo_Sys_Call{Call: _e.mock.On("Sys")}
}

func (_c *MockFileInfo_Sys_Call) Run(run func()) *MockFileInfo_Sys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileInfo_Sys_Call) Return(_a0 interface{}) *MockFileInfo_Sys_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileInfo_Sys_Call) RunAndReturn(run func() interface{}) *MockFileInfo_Sys_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFileInfo creates a new instance of MockFileInfo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFileInfo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFileInfo {
	mock := &MockFileInfo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
