// Code generated by MockGen. DO NOT EDIT.
// Source: ./tss/coordinator.go

// Package mock_tss is a generated GoMock package.
package mock_tss

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

// MockTssProcess is a mock of TssProcess interface.
type MockTssProcess struct {
	ctrl     *gomock.Controller
	recorder *MockTssProcessMockRecorder
}

// MockTssProcessMockRecorder is the mock recorder for MockTssProcess.
type MockTssProcessMockRecorder struct {
	mock *MockTssProcess
}

// NewMockTssProcess creates a new mock instance.
func NewMockTssProcess(ctrl *gomock.Controller) *MockTssProcess {
	mock := &MockTssProcess{ctrl: ctrl}
	mock.recorder = &MockTssProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTssProcess) EXPECT() *MockTssProcessMockRecorder {
	return m.recorder
}

// Ready mocks base method.
func (m *MockTssProcess) Ready(readyMap map[peer.ID]bool, excludedPeers []peer.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready", readyMap, excludedPeers)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ready indicates an expected call of Ready.
func (mr *MockTssProcessMockRecorder) Ready(readyMap, excludedPeers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockTssProcess)(nil).Ready), readyMap, excludedPeers)
}

// Retryable mocks base method.
func (m *MockTssProcess) Retryable() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Retryable")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Retryable indicates an expected call of Retryable.
func (mr *MockTssProcessMockRecorder) Retryable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retryable", reflect.TypeOf((*MockTssProcess)(nil).Retryable))
}

// Run mocks base method.
func (m *MockTssProcess) Run(ctx context.Context, coordinator bool, resultChn chan interface{}, params []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx, coordinator, resultChn, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockTssProcessMockRecorder) Run(ctx, coordinator, resultChn, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockTssProcess)(nil).Run), ctx, coordinator, resultChn, params)
}

// SessionID mocks base method.
func (m *MockTssProcess) SessionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SessionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SessionID indicates an expected call of SessionID.
func (mr *MockTssProcessMockRecorder) SessionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SessionID", reflect.TypeOf((*MockTssProcess)(nil).SessionID))
}

// StartParams mocks base method.
func (m *MockTssProcess) StartParams(readyMap map[peer.ID]bool) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartParams", readyMap)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// StartParams indicates an expected call of StartParams.
func (mr *MockTssProcessMockRecorder) StartParams(readyMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartParams", reflect.TypeOf((*MockTssProcess)(nil).StartParams), readyMap)
}

// Stop mocks base method.
func (m *MockTssProcess) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockTssProcessMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTssProcess)(nil).Stop))
}

// ValidCoordinators mocks base method.
func (m *MockTssProcess) ValidCoordinators() []peer.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidCoordinators")
	ret0, _ := ret[0].([]peer.ID)
	return ret0
}

// ValidCoordinators indicates an expected call of ValidCoordinators.
func (mr *MockTssProcessMockRecorder) ValidCoordinators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidCoordinators", reflect.TypeOf((*MockTssProcess)(nil).ValidCoordinators))
}