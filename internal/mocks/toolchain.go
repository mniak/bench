// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/bench/domain (interfaces: Toolchain)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/mniak/bench/domain"
)

// MockToolchain is a mock of Toolchain interface.
type MockToolchain struct {
	ctrl     *gomock.Controller
	recorder *MockToolchainMockRecorder
}

// MockToolchainMockRecorder is the mock recorder for MockToolchain.
type MockToolchainMockRecorder struct {
	mock *MockToolchain
}

// NewMockToolchain creates a new mock instance.
func NewMockToolchain(ctrl *gomock.Controller) *MockToolchain {
	mock := &MockToolchain{ctrl: ctrl}
	mock.recorder = &MockToolchainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToolchain) EXPECT() *MockToolchainMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockToolchain) Build(arg0 domain.BuildRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockToolchainMockRecorder) Build(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockToolchain)(nil).Build), arg0)
}
