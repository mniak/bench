// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/bench (interfaces: ProgramFinder,ToolchainProducer,Builder)

// Package mock_bench is a generated GoMock package.
package mock_bench

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	toolchain "github.com/mniak/bench/toolchain"
)

// MockProgramFinder is a mock of ProgramFinder interface.
type MockProgramFinder struct {
	ctrl     *gomock.Controller
	recorder *MockProgramFinderMockRecorder
}

// MockProgramFinderMockRecorder is the mock recorder for MockProgramFinder.
type MockProgramFinderMockRecorder struct {
	mock *MockProgramFinder
}

// NewMockProgramFinder creates a new mock instance.
func NewMockProgramFinder(ctrl *gomock.Controller) *MockProgramFinder {
	mock := &MockProgramFinder{ctrl: ctrl}
	mock.recorder = &MockProgramFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProgramFinder) EXPECT() *MockProgramFinderMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockProgramFinder) Find(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockProgramFinderMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockProgramFinder)(nil).Find), arg0)
}

// MockToolchainProducer is a mock of ToolchainProducer interface.
type MockToolchainProducer struct {
	ctrl     *gomock.Controller
	recorder *MockToolchainProducerMockRecorder
}

// MockToolchainProducerMockRecorder is the mock recorder for MockToolchainProducer.
type MockToolchainProducerMockRecorder struct {
	mock *MockToolchainProducer
}

// NewMockToolchainProducer creates a new mock instance.
func NewMockToolchainProducer(ctrl *gomock.Controller) *MockToolchainProducer {
	mock := &MockToolchainProducer{ctrl: ctrl}
	mock.recorder = &MockToolchainProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToolchainProducer) EXPECT() *MockToolchainProducerMockRecorder {
	return m.recorder
}

// Produce mocks base method.
func (m *MockToolchainProducer) Produce(arg0 string) (toolchain.Toolchain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", arg0)
	ret0, _ := ret[0].(toolchain.Toolchain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Produce indicates an expected call of Produce.
func (mr *MockToolchainProducerMockRecorder) Produce(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockToolchainProducer)(nil).Produce), arg0)
}

// MockBuilder is a mock of Builder interface.
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder.
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance.
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockBuilder) Build(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockBuilderMockRecorder) Build(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBuilder)(nil).Build), arg0)
}
