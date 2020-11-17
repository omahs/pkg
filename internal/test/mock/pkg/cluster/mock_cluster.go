// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wetware/ww/pkg/cluster (interfaces: Announcer,Clock,PeerSet)

// Package mock_cluster is a generated GoMock package.
package mock_cluster

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	peer "github.com/libp2p/go-libp2p-core/peer"
	reflect "reflect"
	time "time"
)

// MockAnnouncer is a mock of Announcer interface
type MockAnnouncer struct {
	ctrl     *gomock.Controller
	recorder *MockAnnouncerMockRecorder
}

// MockAnnouncerMockRecorder is the mock recorder for MockAnnouncer
type MockAnnouncerMockRecorder struct {
	mock *MockAnnouncer
}

// NewMockAnnouncer creates a new mock instance
func NewMockAnnouncer(ctrl *gomock.Controller) *MockAnnouncer {
	mock := &MockAnnouncer{ctrl: ctrl}
	mock.recorder = &MockAnnouncerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAnnouncer) EXPECT() *MockAnnouncerMockRecorder {
	return m.recorder
}

// Announce mocks base method
func (m *MockAnnouncer) Announce(arg0 context.Context, arg1 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Announce", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Announce indicates an expected call of Announce
func (mr *MockAnnouncerMockRecorder) Announce(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Announce", reflect.TypeOf((*MockAnnouncer)(nil).Announce), arg0, arg1)
}

// Namespace mocks base method
func (m *MockAnnouncer) Namespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// Namespace indicates an expected call of Namespace
func (mr *MockAnnouncerMockRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockAnnouncer)(nil).Namespace))
}

// MockClock is a mock of Clock interface
type MockClock struct {
	ctrl     *gomock.Controller
	recorder *MockClockMockRecorder
}

// MockClockMockRecorder is the mock recorder for MockClock
type MockClockMockRecorder struct {
	mock *MockClock
}

// NewMockClock creates a new mock instance
func NewMockClock(ctrl *gomock.Controller) *MockClock {
	mock := &MockClock{ctrl: ctrl}
	mock.recorder = &MockClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClock) EXPECT() *MockClockMockRecorder {
	return m.recorder
}

// Advance mocks base method
func (m *MockClock) Advance(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Advance", arg0)
}

// Advance indicates an expected call of Advance
func (mr *MockClockMockRecorder) Advance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Advance", reflect.TypeOf((*MockClock)(nil).Advance), arg0)
}

// MockPeerSet is a mock of PeerSet interface
type MockPeerSet struct {
	ctrl     *gomock.Controller
	recorder *MockPeerSetMockRecorder
}

// MockPeerSetMockRecorder is the mock recorder for MockPeerSet
type MockPeerSetMockRecorder struct {
	mock *MockPeerSet
}

// NewMockPeerSet creates a new mock instance
func NewMockPeerSet(ctrl *gomock.Controller) *MockPeerSet {
	mock := &MockPeerSet{ctrl: ctrl}
	mock.recorder = &MockPeerSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPeerSet) EXPECT() *MockPeerSetMockRecorder {
	return m.recorder
}

// Contains mocks base method
func (m *MockPeerSet) Contains(arg0 peer.ID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contains", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Contains indicates an expected call of Contains
func (mr *MockPeerSetMockRecorder) Contains(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockPeerSet)(nil).Contains), arg0)
}

// Peers mocks base method
func (m *MockPeerSet) Peers() peer.IDSlice {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peers")
	ret0, _ := ret[0].(peer.IDSlice)
	return ret0
}

// Peers indicates an expected call of Peers
func (mr *MockPeerSetMockRecorder) Peers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peers", reflect.TypeOf((*MockPeerSet)(nil).Peers))
}

// Upsert mocks base method
func (m *MockPeerSet) Upsert(arg0 peer.ID, arg1 uint64, arg2 time.Duration) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockPeerSetMockRecorder) Upsert(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockPeerSet)(nil).Upsert), arg0, arg1, arg2)
}
