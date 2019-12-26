// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/networkservicemesh/networkservicemesh/forwarder/api/forwarder (interfaces: ForwarderClient)

// Package tests is a generated GoMock package.
package tests

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	crossconnect "github.com/networkservicemesh/networkservicemesh/controlplane/api/crossconnect"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockForwarderClient is a mock of ForwarderClient interface
type MockForwarderClient struct {
	ctrl     *gomock.Controller
	recorder *MockForwarderClientMockRecorder
}

// MockForwarderClientMockRecorder is the mock recorder for MockForwarderClient
type MockForwarderClientMockRecorder struct {
	mock *MockForwarderClient
}

// NewMockForwarderClient creates a new mock instance
func NewMockForwarderClient(ctrl *gomock.Controller) *MockForwarderClient {
	mock := &MockForwarderClient{ctrl: ctrl}
	mock.recorder = &MockForwarderClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockForwarderClient) EXPECT() *MockForwarderClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockForwarderClient) Close(arg0 context.Context, arg1 *crossconnect.CrossConnect, arg2 ...grpc.CallOption) (*empty.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Close", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Close indicates an expected call of Close
func (mr *MockForwarderClientMockRecorder) Close(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockForwarderClient)(nil).Close), varargs...)
}

// Request mocks base method
func (m *MockForwarderClient) Request(arg0 context.Context, arg1 *crossconnect.CrossConnect, arg2 ...grpc.CallOption) (*crossconnect.CrossConnect, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Request", varargs...)
	ret0, _ := ret[0].(*crossconnect.CrossConnect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request
func (mr *MockForwarderClientMockRecorder) Request(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockForwarderClient)(nil).Request), varargs...)
}