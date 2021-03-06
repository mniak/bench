// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/bench/domain (interfaces: Tester)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/mniak/bench/domain"
)

// MockTester is a mock of Tester interface.
type MockTester struct {
	ctrl     *gomock.Controller
	recorder *MockTesterMockRecorder
}

// MockTesterMockRecorder is the mock recorder for MockTester.
type MockTesterMockRecorder struct {
	mock *MockTester
}

// NewMockTester creates a new mock instance.
func NewMockTester(ctrl *gomock.Controller) *MockTester {
	mock := &MockTester{ctrl: ctrl}
	mock.recorder = &MockTesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTester) EXPECT() *MockTesterMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockTester) Start(arg0 domain.Test) (domain.StartedTest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(domain.StartedTest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start.
func (mr *MockTesterMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTester)(nil).Start), arg0)
}

// Wait mocks base method.
func (m *MockTester) Wait(arg0 domain.StartedTest) (domain.TestResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(domain.TestResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Wait indicates an expected call of Wait.
func (mr *MockTesterMockRecorder) Wait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockTester)(nil).Wait), arg0)
}
