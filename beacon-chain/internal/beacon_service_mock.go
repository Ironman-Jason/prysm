// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1 (interfaces: BeaconServiceServer,BeaconService_LatestAttestationServer,BeaconService_ValidatorAssignmentsServer)

package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	v10 "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	metadata "google.golang.org/grpc/metadata"
)

// MockBeaconServiceServer is a mock of BeaconServiceServer interface
type MockBeaconServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconServiceServerMockRecorder
}

// MockBeaconServiceServerMockRecorder is the mock recorder for MockBeaconServiceServer
type MockBeaconServiceServerMockRecorder struct {
	mock *MockBeaconServiceServer
}

// NewMockBeaconServiceServer creates a new mock instance
func NewMockBeaconServiceServer(ctrl *gomock.Controller) *MockBeaconServiceServer {
	mock := &MockBeaconServiceServer{ctrl: ctrl}
	mock.recorder = &MockBeaconServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconServiceServer) EXPECT() *MockBeaconServiceServerMockRecorder {
	return m.recorder
}

// CanonicalHead mocks base method
func (m *MockBeaconServiceServer) CanonicalHead(arg0 context.Context, arg1 *empty.Empty) (*v1.BeaconBlock, error) {
	ret := m.ctrl.Call(m, "CanonicalHead", arg0, arg1)
	ret0, _ := ret[0].(*v1.BeaconBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanonicalHead indicates an expected call of CanonicalHead
func (mr *MockBeaconServiceServerMockRecorder) CanonicalHead(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanonicalHead", reflect.TypeOf((*MockBeaconServiceServer)(nil).CanonicalHead), arg0, arg1)
}

// CurrentAssignmentsAndGenesisTime mocks base method
func (m *MockBeaconServiceServer) CurrentAssignmentsAndGenesisTime(arg0 context.Context, arg1 *v10.ValidatorAssignmentRequest) (*v10.CurrentAssignmentsResponse, error) {
	ret := m.ctrl.Call(m, "CurrentAssignmentsAndGenesisTime", arg0, arg1)
	ret0, _ := ret[0].(*v10.CurrentAssignmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentAssignmentsAndGenesisTime indicates an expected call of CurrentAssignmentsAndGenesisTime
func (mr *MockBeaconServiceServerMockRecorder) CurrentAssignmentsAndGenesisTime(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentAssignmentsAndGenesisTime", reflect.TypeOf((*MockBeaconServiceServer)(nil).CurrentAssignmentsAndGenesisTime), arg0, arg1)
}

// LatestAttestation mocks base method
func (m *MockBeaconServiceServer) LatestAttestation(arg0 *empty.Empty, arg1 v10.BeaconService_LatestAttestationServer) error {
	ret := m.ctrl.Call(m, "LatestAttestation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LatestAttestation indicates an expected call of LatestAttestation
func (mr *MockBeaconServiceServerMockRecorder) LatestAttestation(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestAttestation", reflect.TypeOf((*MockBeaconServiceServer)(nil).LatestAttestation), arg0, arg1)
}

// ValidatorAssignments mocks base method
func (m *MockBeaconServiceServer) ValidatorAssignments(arg0 *v10.ValidatorAssignmentRequest, arg1 v10.BeaconService_ValidatorAssignmentsServer) error {
	ret := m.ctrl.Call(m, "ValidatorAssignments", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidatorAssignments indicates an expected call of ValidatorAssignments
func (mr *MockBeaconServiceServerMockRecorder) ValidatorAssignments(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorAssignments", reflect.TypeOf((*MockBeaconServiceServer)(nil).ValidatorAssignments), arg0, arg1)
}

// MockBeaconService_LatestAttestationServer is a mock of BeaconService_LatestAttestationServer interface
type MockBeaconService_LatestAttestationServer struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconService_LatestAttestationServerMockRecorder
}

// MockBeaconService_LatestAttestationServerMockRecorder is the mock recorder for MockBeaconService_LatestAttestationServer
type MockBeaconService_LatestAttestationServerMockRecorder struct {
	mock *MockBeaconService_LatestAttestationServer
}

// NewMockBeaconService_LatestAttestationServer creates a new mock instance
func NewMockBeaconService_LatestAttestationServer(ctrl *gomock.Controller) *MockBeaconService_LatestAttestationServer {
	mock := &MockBeaconService_LatestAttestationServer{ctrl: ctrl}
	mock.recorder = &MockBeaconService_LatestAttestationServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconService_LatestAttestationServer) EXPECT() *MockBeaconService_LatestAttestationServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockBeaconService_LatestAttestationServer) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBeaconService_LatestAttestationServerMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconService_LatestAttestationServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockBeaconService_LatestAttestationServer) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBeaconService_LatestAttestationServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconService_LatestAttestationServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockBeaconService_LatestAttestationServer) Send(arg0 *v1.AggregatedAttestation) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockBeaconService_LatestAttestationServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBeaconService_LatestAttestationServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockBeaconService_LatestAttestationServer) SendHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockBeaconService_LatestAttestationServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBeaconService_LatestAttestationServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockBeaconService_LatestAttestationServer) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBeaconService_LatestAttestationServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconService_LatestAttestationServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockBeaconService_LatestAttestationServer) SetHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockBeaconService_LatestAttestationServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBeaconService_LatestAttestationServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockBeaconService_LatestAttestationServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockBeaconService_LatestAttestationServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBeaconService_LatestAttestationServer)(nil).SetTrailer), arg0)
}

// MockBeaconService_ValidatorAssignmentsServer is a mock of BeaconService_ValidatorAssignmentsServer interface
type MockBeaconService_ValidatorAssignmentsServer struct {
	ctrl     *gomock.Controller
	recorder *MockBeaconService_ValidatorAssignmentsServerMockRecorder
}

// MockBeaconService_ValidatorAssignmentsServerMockRecorder is the mock recorder for MockBeaconService_ValidatorAssignmentsServer
type MockBeaconService_ValidatorAssignmentsServerMockRecorder struct {
	mock *MockBeaconService_ValidatorAssignmentsServer
}

// NewMockBeaconService_ValidatorAssignmentsServer creates a new mock instance
func NewMockBeaconService_ValidatorAssignmentsServer(ctrl *gomock.Controller) *MockBeaconService_ValidatorAssignmentsServer {
	mock := &MockBeaconService_ValidatorAssignmentsServer{ctrl: ctrl}
	mock.recorder = &MockBeaconService_ValidatorAssignmentsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeaconService_ValidatorAssignmentsServer) EXPECT() *MockBeaconService_ValidatorAssignmentsServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockBeaconService_ValidatorAssignmentsServer) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockBeaconService_ValidatorAssignmentsServerMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBeaconService_ValidatorAssignmentsServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockBeaconService_ValidatorAssignmentsServer) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockBeaconService_ValidatorAssignmentsServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockBeaconService_ValidatorAssignmentsServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockBeaconService_ValidatorAssignmentsServer) Send(arg0 *v10.ValidatorAssignmentResponse) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockBeaconService_ValidatorAssignmentsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBeaconService_ValidatorAssignmentsServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockBeaconService_ValidatorAssignmentsServer) SendHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockBeaconService_ValidatorAssignmentsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockBeaconService_ValidatorAssignmentsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockBeaconService_ValidatorAssignmentsServer) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockBeaconService_ValidatorAssignmentsServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockBeaconService_ValidatorAssignmentsServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockBeaconService_ValidatorAssignmentsServer) SetHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockBeaconService_ValidatorAssignmentsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockBeaconService_ValidatorAssignmentsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockBeaconService_ValidatorAssignmentsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockBeaconService_ValidatorAssignmentsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockBeaconService_ValidatorAssignmentsServer)(nil).SetTrailer), arg0)
}
