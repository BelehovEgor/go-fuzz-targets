// Code generated by MockGen. DO NOT EDIT.
// Source: examples/interface.go
//
// Generated by this command:
//
//	mockgen -source=examples/interface.go -destination=mocks/mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCar is a mock of Car interface.
type MockCar struct {
	ctrl     *gomock.Controller
	recorder *MockCarMockRecorder
}

// MockCarMockRecorder is the mock recorder for MockCar.
type MockCarMockRecorder struct {
	mock *MockCar
}

// NewMockCar creates a new mock instance.
func NewMockCar(ctrl *gomock.Controller) *MockCar {
	mock := &MockCar{ctrl: ctrl}
	mock.recorder = &MockCarMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCar) EXPECT() *MockCarMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCar) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockCarMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCar)(nil).Close))
}

// GoForward mocks base method.
func (m *MockCar) GoForward(km int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GoForward", km)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GoForward indicates an expected call of GoForward.
func (mr *MockCarMockRecorder) GoForward(km any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoForward", reflect.TypeOf((*MockCar)(nil).GoForward), km)
}

// Open mocks base method.
func (m *MockCar) Open() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Open")
}

// Open indicates an expected call of Open.
func (mr *MockCarMockRecorder) Open() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockCar)(nil).Open))
}

// Price mocks base method.
func (m *MockCar) Price() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Price")
	ret0, _ := ret[0].(int)
	return ret0
}

// Price indicates an expected call of Price.
func (mr *MockCarMockRecorder) Price() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Price", reflect.TypeOf((*MockCar)(nil).Price))
}
