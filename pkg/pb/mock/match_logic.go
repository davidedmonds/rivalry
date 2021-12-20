// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/amg84/om-stream/pkg/pb (interfaces: ValidationServiceClient,ValidationServiceServer,DataServiceClient,DataServiceServer,MatchMakerServiceClient,MatchMakerService_MakeMatchesClient,MatchMakerServiceServer,MatchMakerService_MakeMatchesServer,AssignmentServiceClient,AssignmentServiceServer)

// Package mock_pb is a generated GoMock package.
package mock_pb

import (
	context "context"
	reflect "reflect"

	pb "github.com/amg84/om-stream/pkg/pb"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockValidationServiceClient is a mock of ValidationServiceClient interface.
type MockValidationServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockValidationServiceClientMockRecorder
}

// MockValidationServiceClientMockRecorder is the mock recorder for MockValidationServiceClient.
type MockValidationServiceClientMockRecorder struct {
	mock *MockValidationServiceClient
}

// NewMockValidationServiceClient creates a new mock instance.
func NewMockValidationServiceClient(ctrl *gomock.Controller) *MockValidationServiceClient {
	mock := &MockValidationServiceClient{ctrl: ctrl}
	mock.recorder = &MockValidationServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidationServiceClient) EXPECT() *MockValidationServiceClientMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockValidationServiceClient) Validate(arg0 context.Context, arg1 *pb.Ticket, arg2 ...grpc.CallOption) (*pb.ValidateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Validate", varargs...)
	ret0, _ := ret[0].(*pb.ValidateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockValidationServiceClientMockRecorder) Validate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidationServiceClient)(nil).Validate), varargs...)
}

// MockValidationServiceServer is a mock of ValidationServiceServer interface.
type MockValidationServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockValidationServiceServerMockRecorder
}

// MockValidationServiceServerMockRecorder is the mock recorder for MockValidationServiceServer.
type MockValidationServiceServerMockRecorder struct {
	mock *MockValidationServiceServer
}

// NewMockValidationServiceServer creates a new mock instance.
func NewMockValidationServiceServer(ctrl *gomock.Controller) *MockValidationServiceServer {
	mock := &MockValidationServiceServer{ctrl: ctrl}
	mock.recorder = &MockValidationServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidationServiceServer) EXPECT() *MockValidationServiceServerMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockValidationServiceServer) Validate(arg0 context.Context, arg1 *pb.Ticket) (*pb.ValidateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(*pb.ValidateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockValidationServiceServerMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidationServiceServer)(nil).Validate), arg0, arg1)
}

// MockDataServiceClient is a mock of DataServiceClient interface.
type MockDataServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDataServiceClientMockRecorder
}

// MockDataServiceClientMockRecorder is the mock recorder for MockDataServiceClient.
type MockDataServiceClientMockRecorder struct {
	mock *MockDataServiceClient
}

// NewMockDataServiceClient creates a new mock instance.
func NewMockDataServiceClient(ctrl *gomock.Controller) *MockDataServiceClient {
	mock := &MockDataServiceClient{ctrl: ctrl}
	mock.recorder = &MockDataServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataServiceClient) EXPECT() *MockDataServiceClientMockRecorder {
	return m.recorder
}

// GatherData mocks base method.
func (m *MockDataServiceClient) GatherData(arg0 context.Context, arg1 *pb.Ticket, arg2 ...grpc.CallOption) (*pb.GatherDataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GatherData", varargs...)
	ret0, _ := ret[0].(*pb.GatherDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GatherData indicates an expected call of GatherData.
func (mr *MockDataServiceClientMockRecorder) GatherData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatherData", reflect.TypeOf((*MockDataServiceClient)(nil).GatherData), varargs...)
}

// MockDataServiceServer is a mock of DataServiceServer interface.
type MockDataServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockDataServiceServerMockRecorder
}

// MockDataServiceServerMockRecorder is the mock recorder for MockDataServiceServer.
type MockDataServiceServerMockRecorder struct {
	mock *MockDataServiceServer
}

// NewMockDataServiceServer creates a new mock instance.
func NewMockDataServiceServer(ctrl *gomock.Controller) *MockDataServiceServer {
	mock := &MockDataServiceServer{ctrl: ctrl}
	mock.recorder = &MockDataServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataServiceServer) EXPECT() *MockDataServiceServerMockRecorder {
	return m.recorder
}

// GatherData mocks base method.
func (m *MockDataServiceServer) GatherData(arg0 context.Context, arg1 *pb.Ticket) (*pb.GatherDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatherData", arg0, arg1)
	ret0, _ := ret[0].(*pb.GatherDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GatherData indicates an expected call of GatherData.
func (mr *MockDataServiceServerMockRecorder) GatherData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatherData", reflect.TypeOf((*MockDataServiceServer)(nil).GatherData), arg0, arg1)
}

// MockMatchMakerServiceClient is a mock of MatchMakerServiceClient interface.
type MockMatchMakerServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockMatchMakerServiceClientMockRecorder
}

// MockMatchMakerServiceClientMockRecorder is the mock recorder for MockMatchMakerServiceClient.
type MockMatchMakerServiceClientMockRecorder struct {
	mock *MockMatchMakerServiceClient
}

// NewMockMatchMakerServiceClient creates a new mock instance.
func NewMockMatchMakerServiceClient(ctrl *gomock.Controller) *MockMatchMakerServiceClient {
	mock := &MockMatchMakerServiceClient{ctrl: ctrl}
	mock.recorder = &MockMatchMakerServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatchMakerServiceClient) EXPECT() *MockMatchMakerServiceClientMockRecorder {
	return m.recorder
}

// MakeMatches mocks base method.
func (m *MockMatchMakerServiceClient) MakeMatches(arg0 context.Context, arg1 *pb.MakeMatchesRequest, arg2 ...grpc.CallOption) (pb.MatchMakerService_MakeMatchesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MakeMatches", varargs...)
	ret0, _ := ret[0].(pb.MatchMakerService_MakeMatchesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeMatches indicates an expected call of MakeMatches.
func (mr *MockMatchMakerServiceClientMockRecorder) MakeMatches(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeMatches", reflect.TypeOf((*MockMatchMakerServiceClient)(nil).MakeMatches), varargs...)
}

// MockMatchMakerService_MakeMatchesClient is a mock of MatchMakerService_MakeMatchesClient interface.
type MockMatchMakerService_MakeMatchesClient struct {
	ctrl     *gomock.Controller
	recorder *MockMatchMakerService_MakeMatchesClientMockRecorder
}

// MockMatchMakerService_MakeMatchesClientMockRecorder is the mock recorder for MockMatchMakerService_MakeMatchesClient.
type MockMatchMakerService_MakeMatchesClientMockRecorder struct {
	mock *MockMatchMakerService_MakeMatchesClient
}

// NewMockMatchMakerService_MakeMatchesClient creates a new mock instance.
func NewMockMatchMakerService_MakeMatchesClient(ctrl *gomock.Controller) *MockMatchMakerService_MakeMatchesClient {
	mock := &MockMatchMakerService_MakeMatchesClient{ctrl: ctrl}
	mock.recorder = &MockMatchMakerService_MakeMatchesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatchMakerService_MakeMatchesClient) EXPECT() *MockMatchMakerService_MakeMatchesClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockMatchMakerService_MakeMatchesClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockMatchMakerService_MakeMatchesClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockMatchMakerService_MakeMatchesClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockMatchMakerService_MakeMatchesClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockMatchMakerService_MakeMatchesClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockMatchMakerService_MakeMatchesClient)(nil).Context))
}

// Header mocks base method.
func (m *MockMatchMakerService_MakeMatchesClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockMatchMakerService_MakeMatchesClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockMatchMakerService_MakeMatchesClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockMatchMakerService_MakeMatchesClient) Recv() (*pb.MakeMatchesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*pb.MakeMatchesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockMatchMakerService_MakeMatchesClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockMatchMakerService_MakeMatchesClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockMatchMakerService_MakeMatchesClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockMatchMakerService_MakeMatchesClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockMatchMakerService_MakeMatchesClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockMatchMakerService_MakeMatchesClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockMatchMakerService_MakeMatchesClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockMatchMakerService_MakeMatchesClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockMatchMakerService_MakeMatchesClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockMatchMakerService_MakeMatchesClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockMatchMakerService_MakeMatchesClient)(nil).Trailer))
}

// MockMatchMakerServiceServer is a mock of MatchMakerServiceServer interface.
type MockMatchMakerServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockMatchMakerServiceServerMockRecorder
}

// MockMatchMakerServiceServerMockRecorder is the mock recorder for MockMatchMakerServiceServer.
type MockMatchMakerServiceServerMockRecorder struct {
	mock *MockMatchMakerServiceServer
}

// NewMockMatchMakerServiceServer creates a new mock instance.
func NewMockMatchMakerServiceServer(ctrl *gomock.Controller) *MockMatchMakerServiceServer {
	mock := &MockMatchMakerServiceServer{ctrl: ctrl}
	mock.recorder = &MockMatchMakerServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatchMakerServiceServer) EXPECT() *MockMatchMakerServiceServerMockRecorder {
	return m.recorder
}

// MakeMatches mocks base method.
func (m *MockMatchMakerServiceServer) MakeMatches(arg0 *pb.MakeMatchesRequest, arg1 pb.MatchMakerService_MakeMatchesServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeMatches", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeMatches indicates an expected call of MakeMatches.
func (mr *MockMatchMakerServiceServerMockRecorder) MakeMatches(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeMatches", reflect.TypeOf((*MockMatchMakerServiceServer)(nil).MakeMatches), arg0, arg1)
}

// MockMatchMakerService_MakeMatchesServer is a mock of MatchMakerService_MakeMatchesServer interface.
type MockMatchMakerService_MakeMatchesServer struct {
	ctrl     *gomock.Controller
	recorder *MockMatchMakerService_MakeMatchesServerMockRecorder
}

// MockMatchMakerService_MakeMatchesServerMockRecorder is the mock recorder for MockMatchMakerService_MakeMatchesServer.
type MockMatchMakerService_MakeMatchesServerMockRecorder struct {
	mock *MockMatchMakerService_MakeMatchesServer
}

// NewMockMatchMakerService_MakeMatchesServer creates a new mock instance.
func NewMockMatchMakerService_MakeMatchesServer(ctrl *gomock.Controller) *MockMatchMakerService_MakeMatchesServer {
	mock := &MockMatchMakerService_MakeMatchesServer{ctrl: ctrl}
	mock.recorder = &MockMatchMakerService_MakeMatchesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatchMakerService_MakeMatchesServer) EXPECT() *MockMatchMakerService_MakeMatchesServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockMatchMakerService_MakeMatchesServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockMatchMakerService_MakeMatchesServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockMatchMakerService_MakeMatchesServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockMatchMakerService_MakeMatchesServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockMatchMakerService_MakeMatchesServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockMatchMakerService_MakeMatchesServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockMatchMakerService_MakeMatchesServer) Send(arg0 *pb.MakeMatchesResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockMatchMakerService_MakeMatchesServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMatchMakerService_MakeMatchesServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockMatchMakerService_MakeMatchesServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockMatchMakerService_MakeMatchesServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockMatchMakerService_MakeMatchesServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockMatchMakerService_MakeMatchesServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockMatchMakerService_MakeMatchesServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockMatchMakerService_MakeMatchesServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockMatchMakerService_MakeMatchesServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockMatchMakerService_MakeMatchesServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockMatchMakerService_MakeMatchesServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockMatchMakerService_MakeMatchesServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockMatchMakerService_MakeMatchesServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockMatchMakerService_MakeMatchesServer)(nil).SetTrailer), arg0)
}

// MockAssignmentServiceClient is a mock of AssignmentServiceClient interface.
type MockAssignmentServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAssignmentServiceClientMockRecorder
}

// MockAssignmentServiceClientMockRecorder is the mock recorder for MockAssignmentServiceClient.
type MockAssignmentServiceClientMockRecorder struct {
	mock *MockAssignmentServiceClient
}

// NewMockAssignmentServiceClient creates a new mock instance.
func NewMockAssignmentServiceClient(ctrl *gomock.Controller) *MockAssignmentServiceClient {
	mock := &MockAssignmentServiceClient{ctrl: ctrl}
	mock.recorder = &MockAssignmentServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssignmentServiceClient) EXPECT() *MockAssignmentServiceClientMockRecorder {
	return m.recorder
}

// MakeAssignment mocks base method.
func (m *MockAssignmentServiceClient) MakeAssignment(arg0 context.Context, arg1 *pb.Match, arg2 ...grpc.CallOption) (*pb.Assignment, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MakeAssignment", varargs...)
	ret0, _ := ret[0].(*pb.Assignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeAssignment indicates an expected call of MakeAssignment.
func (mr *MockAssignmentServiceClientMockRecorder) MakeAssignment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeAssignment", reflect.TypeOf((*MockAssignmentServiceClient)(nil).MakeAssignment), varargs...)
}

// MockAssignmentServiceServer is a mock of AssignmentServiceServer interface.
type MockAssignmentServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAssignmentServiceServerMockRecorder
}

// MockAssignmentServiceServerMockRecorder is the mock recorder for MockAssignmentServiceServer.
type MockAssignmentServiceServerMockRecorder struct {
	mock *MockAssignmentServiceServer
}

// NewMockAssignmentServiceServer creates a new mock instance.
func NewMockAssignmentServiceServer(ctrl *gomock.Controller) *MockAssignmentServiceServer {
	mock := &MockAssignmentServiceServer{ctrl: ctrl}
	mock.recorder = &MockAssignmentServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAssignmentServiceServer) EXPECT() *MockAssignmentServiceServerMockRecorder {
	return m.recorder
}

// MakeAssignment mocks base method.
func (m *MockAssignmentServiceServer) MakeAssignment(arg0 context.Context, arg1 *pb.Match) (*pb.Assignment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeAssignment", arg0, arg1)
	ret0, _ := ret[0].(*pb.Assignment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeAssignment indicates an expected call of MakeAssignment.
func (mr *MockAssignmentServiceServerMockRecorder) MakeAssignment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeAssignment", reflect.TypeOf((*MockAssignmentServiceServer)(nil).MakeAssignment), arg0, arg1)
}
