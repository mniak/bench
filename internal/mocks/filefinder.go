// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/bench/domain (interfaces: FileFinder)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileFinder is a mock of FileFinder interface.
type MockFileFinder struct {
	ctrl     *gomock.Controller
	recorder *MockFileFinderMockRecorder
}

// MockFileFinderMockRecorder is the mock recorder for MockFileFinder.
type MockFileFinderMockRecorder struct {
	mock *MockFileFinder
}

// NewMockFileFinder creates a new mock instance.
func NewMockFileFinder(ctrl *gomock.Controller) *MockFileFinder {
	mock := &MockFileFinder{ctrl: ctrl}
	mock.recorder = &MockFileFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileFinder) EXPECT() *MockFileFinderMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockFileFinder) Find(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockFileFinderMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFileFinder)(nil).Find), arg0)
}
