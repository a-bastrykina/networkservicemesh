// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/user/Develop/NSM/networkservicemesh/controlplane/pkg/api/nsm/nsm.go

// Package mock_nsm is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"

	connection "github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	crossconnect "github.com/networkservicemesh/networkservicemesh/controlplane/api/crossconnect"
	networkservice "github.com/networkservicemesh/networkservicemesh/controlplane/api/networkservice"
	registry "github.com/networkservicemesh/networkservicemesh/controlplane/api/registry"
	x "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/api/nsm"
	model "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/model"
	properties "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/properties"
	serviceregistry "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/serviceregistry"
	connectionmonitor "github.com/networkservicemesh/networkservicemesh/sdk/monitor/connectionmonitor"
	crossconnect0 "github.com/networkservicemesh/networkservicemesh/sdk/monitor/crossconnect"
)

// MockClientConnection is a mock_nsm of ClientConnection interface
type MockClientConnection struct {
	ctrl     *gomock.Controller
	recorder *MockClientConnectionMockRecorder
}

// MockClientConnectionMockRecorder is the mock_nsm recorder for MockClientConnection
type MockClientConnectionMockRecorder struct {
	mock *MockClientConnection
}

// NewMockClientConnection creates a new mock_nsm instance
func NewMockClientConnection(ctrl *gomock.Controller) *MockClientConnection {
	mock := &MockClientConnection{ctrl: ctrl}
	mock.recorder = &MockClientConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientConnection) EXPECT() *MockClientConnectionMockRecorder {
	return m.recorder
}

// GetID mocks base method
func (m *MockClientConnection) GetID() string {
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockClientConnectionMockRecorder) GetID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockClientConnection)(nil).GetID))
}

// GetConnectionSource mocks base method
func (m *MockClientConnection) GetConnectionSource() *connection.Connection {
	ret := m.ctrl.Call(m, "GetConnectionSource")
	ret0, _ := ret[0].(*connection.Connection)
	return ret0
}

// GetConnectionSource indicates an expected call of GetConnectionSource
func (mr *MockClientConnectionMockRecorder) GetConnectionSource() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectionSource", reflect.TypeOf((*MockClientConnection)(nil).GetConnectionSource))
}

// GetConnectionDestination mocks base method
func (m *MockClientConnection) GetConnectionDestination() *connection.Connection {
	ret := m.ctrl.Call(m, "GetConnectionDestination")
	ret0, _ := ret[0].(*connection.Connection)
	return ret0
}

// GetConnectionDestination indicates an expected call of GetConnectionDestination
func (mr *MockClientConnectionMockRecorder) GetConnectionDestination() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectionDestination", reflect.TypeOf((*MockClientConnection)(nil).GetConnectionDestination))
}

// GetNetworkService mocks base method
func (m *MockClientConnection) GetNetworkService() string {
	ret := m.ctrl.Call(m, "GetNetworkService")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNetworkService indicates an expected call of GetNetworkService
func (mr *MockClientConnectionMockRecorder) GetNetworkService() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkService", reflect.TypeOf((*MockClientConnection)(nil).GetNetworkService))
}

// MockNetworkServiceClient is a mock_nsm of NetworkServiceClient interface
type MockNetworkServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceClientMockRecorder
}

// MockNetworkServiceClientMockRecorder is the mock_nsm recorder for MockNetworkServiceClient
type MockNetworkServiceClientMockRecorder struct {
	mock *MockNetworkServiceClient
}

// NewMockNetworkServiceClient creates a new mock_nsm instance
func NewMockNetworkServiceClient(ctrl *gomock.Controller) *MockNetworkServiceClient {
	mock := &MockNetworkServiceClient{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkServiceClient) EXPECT() *MockNetworkServiceClientMockRecorder {
	return m.recorder
}

// Request mocks base method
func (m *MockNetworkServiceClient) Request(ctx context.Context, request *networkservice.NetworkServiceRequest) (*connection.Connection, error) {
	ret := m.ctrl.Call(m, "Request", ctx, request)
	ret0, _ := ret[0].(*connection.Connection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request
func (mr *MockNetworkServiceClientMockRecorder) Request(ctx, request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockNetworkServiceClient)(nil).Request), ctx, request)
}

// Close mocks base method
func (m *MockNetworkServiceClient) Close(ctx context.Context, connection *connection.Connection) error {
	ret := m.ctrl.Call(m, "Close", ctx, connection)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockNetworkServiceClientMockRecorder) Close(ctx, connection interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockNetworkServiceClient)(nil).Close), ctx, connection)
}

// Cleanup mocks base method
func (m *MockNetworkServiceClient) Cleanup() error {
	ret := m.ctrl.Call(m, "Cleanup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Cleanup indicates an expected call of Cleanup
func (mr *MockNetworkServiceClientMockRecorder) Cleanup() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockNetworkServiceClient)(nil).Cleanup))
}

// MockNetworkServiceRequestManager is a mock_nsm of NetworkServiceRequestManager interface
type MockNetworkServiceRequestManager struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceRequestManagerMockRecorder
}

// MockNetworkServiceRequestManagerMockRecorder is the mock_nsm recorder for MockNetworkServiceRequestManager
type MockNetworkServiceRequestManagerMockRecorder struct {
	mock *MockNetworkServiceRequestManager
}

// NewMockNetworkServiceRequestManager creates a new mock_nsm instance
func NewMockNetworkServiceRequestManager(ctrl *gomock.Controller) *MockNetworkServiceRequestManager {
	mock := &MockNetworkServiceRequestManager{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceRequestManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkServiceRequestManager) EXPECT() *MockNetworkServiceRequestManagerMockRecorder {
	return m.recorder
}

// LocalManager mocks base method
func (m *MockNetworkServiceRequestManager) LocalManager(clientConnection x.ClientConnection) networkservice.NetworkServiceServer {
	ret := m.ctrl.Call(m, "LocalManager", clientConnection)
	ret0, _ := ret[0].(networkservice.NetworkServiceServer)
	return ret0
}

// LocalManager indicates an expected call of LocalManager
func (mr *MockNetworkServiceRequestManagerMockRecorder) LocalManager(clientConnection interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalManager", reflect.TypeOf((*MockNetworkServiceRequestManager)(nil).LocalManager), clientConnection)
}

// RemoteManager mocks base method
func (m *MockNetworkServiceRequestManager) RemoteManager() networkservice.NetworkServiceServer {
	ret := m.ctrl.Call(m, "RemoteManager")
	ret0, _ := ret[0].(networkservice.NetworkServiceServer)
	return ret0
}

// RemoteManager indicates an expected call of RemoteManager
func (mr *MockNetworkServiceRequestManagerMockRecorder) RemoteManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteManager", reflect.TypeOf((*MockNetworkServiceRequestManager)(nil).RemoteManager))
}

// MockNetworkServiceHealProcessor is a mock_nsm of NetworkServiceHealProcessor interface
type MockNetworkServiceHealProcessor struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceHealProcessorMockRecorder
}

// MockNetworkServiceHealProcessorMockRecorder is the mock_nsm recorder for MockNetworkServiceHealProcessor
type MockNetworkServiceHealProcessorMockRecorder struct {
	mock *MockNetworkServiceHealProcessor
}

// NewMockNetworkServiceHealProcessor creates a new mock_nsm instance
func NewMockNetworkServiceHealProcessor(ctrl *gomock.Controller) *MockNetworkServiceHealProcessor {
	mock := &MockNetworkServiceHealProcessor{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceHealProcessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkServiceHealProcessor) EXPECT() *MockNetworkServiceHealProcessorMockRecorder {
	return m.recorder
}

// Heal mocks base method
func (m *MockNetworkServiceHealProcessor) Heal(ctx context.Context, clientConnection x.ClientConnection, healState x.HealState) {
	m.ctrl.Call(m, "Heal", ctx, clientConnection, healState)
}

// Heal indicates an expected call of Heal
func (mr *MockNetworkServiceHealProcessorMockRecorder) Heal(ctx, clientConnection, healState interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Heal", reflect.TypeOf((*MockNetworkServiceHealProcessor)(nil).Heal), ctx, clientConnection, healState)
}

// CloseConnection mocks base method
func (m *MockNetworkServiceHealProcessor) CloseConnection(ctx context.Context, clientConnection x.ClientConnection) error {
	ret := m.ctrl.Call(m, "CloseConnection", ctx, clientConnection)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseConnection indicates an expected call of CloseConnection
func (mr *MockNetworkServiceHealProcessorMockRecorder) CloseConnection(ctx, clientConnection interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseConnection", reflect.TypeOf((*MockNetworkServiceHealProcessor)(nil).CloseConnection), ctx, clientConnection)
}

// MockMonitorManager is a mock_nsm of MonitorManager interface
type MockMonitorManager struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorManagerMockRecorder
}

// MockMonitorManagerMockRecorder is the mock_nsm recorder for MockMonitorManager
type MockMonitorManagerMockRecorder struct {
	mock *MockMonitorManager
}

// NewMockMonitorManager creates a new mock_nsm instance
func NewMockMonitorManager(ctrl *gomock.Controller) *MockMonitorManager {
	mock := &MockMonitorManager{ctrl: ctrl}
	mock.recorder = &MockMonitorManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMonitorManager) EXPECT() *MockMonitorManagerMockRecorder {
	return m.recorder
}

// CrossConnectMonitor mocks base method
func (m *MockMonitorManager) CrossConnectMonitor() crossconnect0.MonitorServer {
	ret := m.ctrl.Call(m, "CrossConnectMonitor")
	ret0, _ := ret[0].(crossconnect0.MonitorServer)
	return ret0
}

// CrossConnectMonitor indicates an expected call of CrossConnectMonitor
func (mr *MockMonitorManagerMockRecorder) CrossConnectMonitor() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossConnectMonitor", reflect.TypeOf((*MockMonitorManager)(nil).CrossConnectMonitor))
}

// LocalConnectionMonitor mocks base method
func (m *MockMonitorManager) LocalConnectionMonitor(workspace string) connectionmonitor.MonitorServer {
	ret := m.ctrl.Call(m, "LocalConnectionMonitor", workspace)
	ret0, _ := ret[0].(connectionmonitor.MonitorServer)
	return ret0
}

// LocalConnectionMonitor indicates an expected call of LocalConnectionMonitor
func (mr *MockMonitorManagerMockRecorder) LocalConnectionMonitor(workspace interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalConnectionMonitor", reflect.TypeOf((*MockMonitorManager)(nil).LocalConnectionMonitor), workspace)
}

// MockNetworkServiceManager is a mock_nsm of NetworkServiceManager interface
type MockNetworkServiceManager struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceManagerMockRecorder
}

// MockNetworkServiceManagerMockRecorder is the mock_nsm recorder for MockNetworkServiceManager
type MockNetworkServiceManagerMockRecorder struct {
	mock *MockNetworkServiceManager
}

// NewMockNetworkServiceManager creates a new mock_nsm instance
func NewMockNetworkServiceManager(ctrl *gomock.Controller) *MockNetworkServiceManager {
	mock := &MockNetworkServiceManager{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkServiceManager) EXPECT() *MockNetworkServiceManagerMockRecorder {
	return m.recorder
}

// GetHealProperties mocks base method
func (m *MockNetworkServiceManager) GetHealProperties() *properties.Properties {
	ret := m.ctrl.Call(m, "GetHealProperties")
	ret0, _ := ret[0].(*properties.Properties)
	return ret0
}

// GetHealProperties indicates an expected call of GetHealProperties
func (mr *MockNetworkServiceManagerMockRecorder) GetHealProperties() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHealProperties", reflect.TypeOf((*MockNetworkServiceManager)(nil).GetHealProperties))
}

// WaitForForwarder mocks base method
func (m *MockNetworkServiceManager) WaitForForwarder(ctx context.Context, duration time.Duration) error {
	ret := m.ctrl.Call(m, "WaitForForwarder", ctx, duration)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForForwarder indicates an expected call of WaitForForwarder
func (mr *MockNetworkServiceManagerMockRecorder) WaitForForwarder(ctx, duration interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForForwarder", reflect.TypeOf((*MockNetworkServiceManager)(nil).WaitForForwarder), ctx, duration)
}

// RemoteConnectionLost mocks base method
func (m *MockNetworkServiceManager) RemoteConnectionLost(ctx context.Context, clientConnection x.ClientConnection) {
	m.ctrl.Call(m, "RemoteConnectionLost", ctx, clientConnection)
}

// RemoteConnectionLost indicates an expected call of RemoteConnectionLost
func (mr *MockNetworkServiceManagerMockRecorder) RemoteConnectionLost(ctx, clientConnection interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteConnectionLost", reflect.TypeOf((*MockNetworkServiceManager)(nil).RemoteConnectionLost), ctx, clientConnection)
}

// NotifyRenamedEndpoint mocks base method
func (m *MockNetworkServiceManager) NotifyRenamedEndpoint(nseOldName, nseNewName string) {
	m.ctrl.Call(m, "NotifyRenamedEndpoint", nseOldName, nseNewName)
}

// NotifyRenamedEndpoint indicates an expected call of NotifyRenamedEndpoint
func (mr *MockNetworkServiceManagerMockRecorder) NotifyRenamedEndpoint(nseOldName, nseNewName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyRenamedEndpoint", reflect.TypeOf((*MockNetworkServiceManager)(nil).NotifyRenamedEndpoint), nseOldName, nseNewName)
}

// NseManager mocks base method
func (m *MockNetworkServiceManager) NseManager() x.NetworkServiceEndpointManager {
	ret := m.ctrl.Call(m, "NseManager")
	ret0, _ := ret[0].(x.NetworkServiceEndpointManager)
	return ret0
}

// NseManager indicates an expected call of NseManager
func (mr *MockNetworkServiceManagerMockRecorder) NseManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NseManager", reflect.TypeOf((*MockNetworkServiceManager)(nil).NseManager))
}

// SetRemoteServer mocks base method
func (m *MockNetworkServiceManager) SetRemoteServer(server networkservice.NetworkServiceServer) {
	m.ctrl.Call(m, "SetRemoteServer", server)
}

// SetRemoteServer indicates an expected call of SetRemoteServer
func (mr *MockNetworkServiceManagerMockRecorder) SetRemoteServer(server interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRemoteServer", reflect.TypeOf((*MockNetworkServiceManager)(nil).SetRemoteServer), server)
}

// Model mocks base method
func (m *MockNetworkServiceManager) Model() model.Model {
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(model.Model)
	return ret0
}

// Model indicates an expected call of Model
func (mr *MockNetworkServiceManagerMockRecorder) Model() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockNetworkServiceManager)(nil).Model))
}

// Heal mocks base method
func (m *MockNetworkServiceManager) Heal(ctx context.Context, clientConnection x.ClientConnection, healState x.HealState) {
	m.ctrl.Call(m, "Heal", ctx, clientConnection, healState)
}

// Heal indicates an expected call of Heal
func (mr *MockNetworkServiceManagerMockRecorder) Heal(ctx, clientConnection, healState interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Heal", reflect.TypeOf((*MockNetworkServiceManager)(nil).Heal), ctx, clientConnection, healState)
}

// CloseConnection mocks base method
func (m *MockNetworkServiceManager) CloseConnection(ctx context.Context, clientConnection x.ClientConnection) error {
	ret := m.ctrl.Call(m, "CloseConnection", ctx, clientConnection)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseConnection indicates an expected call of CloseConnection
func (mr *MockNetworkServiceManagerMockRecorder) CloseConnection(ctx, clientConnection interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseConnection", reflect.TypeOf((*MockNetworkServiceManager)(nil).CloseConnection), ctx, clientConnection)
}

// ServiceRegistry mocks base method
func (m *MockNetworkServiceManager) ServiceRegistry() serviceregistry.ServiceRegistry {
	ret := m.ctrl.Call(m, "ServiceRegistry")
	ret0, _ := ret[0].(serviceregistry.ServiceRegistry)
	return ret0
}

// ServiceRegistry indicates an expected call of ServiceRegistry
func (mr *MockNetworkServiceManagerMockRecorder) ServiceRegistry() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceRegistry", reflect.TypeOf((*MockNetworkServiceManager)(nil).ServiceRegistry))
}

// RestoreConnections mocks base method
func (m *MockNetworkServiceManager) RestoreConnections(xcons []*crossconnect.CrossConnect, forwarder string, manager x.MonitorManager) {
	m.ctrl.Call(m, "RestoreConnections", xcons, forwarder, manager)
}

// RestoreConnections indicates an expected call of RestoreConnections
func (mr *MockNetworkServiceManagerMockRecorder) RestoreConnections(xcons, forwarder, manager interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreConnections", reflect.TypeOf((*MockNetworkServiceManager)(nil).RestoreConnections), xcons, forwarder, manager)
}

// MockNetworkServiceEndpointManager is a mock_nsm of NetworkServiceEndpointManager interface
type MockNetworkServiceEndpointManager struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceEndpointManagerMockRecorder
}

// MockNetworkServiceEndpointManagerMockRecorder is the mock_nsm recorder for MockNetworkServiceEndpointManager
type MockNetworkServiceEndpointManagerMockRecorder struct {
	mock *MockNetworkServiceEndpointManager
}

// NewMockNetworkServiceEndpointManager creates a new mock_nsm instance
func NewMockNetworkServiceEndpointManager(ctrl *gomock.Controller) *MockNetworkServiceEndpointManager {
	mock := &MockNetworkServiceEndpointManager{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceEndpointManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkServiceEndpointManager) EXPECT() *MockNetworkServiceEndpointManagerMockRecorder {
	return m.recorder
}

// GetEndpoint mocks base method
func (m *MockNetworkServiceEndpointManager) GetEndpoint(ctx context.Context, requestConnection *connection.Connection, ignoreEndpoints map[registry.EndpointNSMName]*registry.NSERegistration) (*registry.NSERegistration, error) {
	ret := m.ctrl.Call(m, "GetEndpoint", ctx, requestConnection, ignoreEndpoints)
	ret0, _ := ret[0].(*registry.NSERegistration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEndpoint indicates an expected call of GetEndpoint
func (mr *MockNetworkServiceEndpointManagerMockRecorder) GetEndpoint(ctx, requestConnection, ignoreEndpoints interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEndpoint", reflect.TypeOf((*MockNetworkServiceEndpointManager)(nil).GetEndpoint), ctx, requestConnection, ignoreEndpoints)
}

// CreateNSEClient mocks base method
func (m *MockNetworkServiceEndpointManager) CreateNSEClient(ctx context.Context, endpoint *registry.NSERegistration) (x.NetworkServiceClient, error) {
	ret := m.ctrl.Call(m, "CreateNSEClient", ctx, endpoint)
	ret0, _ := ret[0].(x.NetworkServiceClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNSEClient indicates an expected call of CreateNSEClient
func (mr *MockNetworkServiceEndpointManagerMockRecorder) CreateNSEClient(ctx, endpoint interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNSEClient", reflect.TypeOf((*MockNetworkServiceEndpointManager)(nil).CreateNSEClient), ctx, endpoint)
}

// IsLocalEndpoint mocks base method
func (m *MockNetworkServiceEndpointManager) IsLocalEndpoint(endpoint *registry.NSERegistration) bool {
	ret := m.ctrl.Call(m, "IsLocalEndpoint", endpoint)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLocalEndpoint indicates an expected call of IsLocalEndpoint
func (mr *MockNetworkServiceEndpointManagerMockRecorder) IsLocalEndpoint(endpoint interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLocalEndpoint", reflect.TypeOf((*MockNetworkServiceEndpointManager)(nil).IsLocalEndpoint), endpoint)
}

// CheckUpdateNSE mocks base method
func (m *MockNetworkServiceEndpointManager) CheckUpdateNSE(ctx context.Context, reg *registry.NSERegistration) bool {
	ret := m.ctrl.Call(m, "CheckUpdateNSE", ctx, reg)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckUpdateNSE indicates an expected call of CheckUpdateNSE
func (mr *MockNetworkServiceEndpointManagerMockRecorder) CheckUpdateNSE(ctx, reg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUpdateNSE", reflect.TypeOf((*MockNetworkServiceEndpointManager)(nil).CheckUpdateNSE), ctx, reg)
}
