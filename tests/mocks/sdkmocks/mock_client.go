// Code generated by MockGen. DO NOT EDIT.
// Source: common/client/bmcapi/client.go

// Package sdkmocks is a generated GoMock package.
package sdkmocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
)

// MockBmcApiSdkClient is a mock of BmcApiSdkClient interface.
type MockBmcApiSdkClient struct {
	ctrl     *gomock.Controller
	recorder *MockBmcApiSdkClientMockRecorder
}

// MockBmcApiSdkClientMockRecorder is the mock recorder for MockBmcApiSdkClient.
type MockBmcApiSdkClientMockRecorder struct {
	mock *MockBmcApiSdkClient
}

// NewMockBmcApiSdkClient creates a new mock instance.
func NewMockBmcApiSdkClient(ctrl *gomock.Controller) *MockBmcApiSdkClient {
	mock := &MockBmcApiSdkClient{ctrl: ctrl}
	mock.recorder = &MockBmcApiSdkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBmcApiSdkClient) EXPECT() *MockBmcApiSdkClientMockRecorder {
	return m.recorder
}

// ServerDelete mocks base method.
func (m *MockBmcApiSdkClient) ServerDelete(serverId string) (bmcapisdk.DeleteResult, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerDelete", serverId)
	ret0, _ := ret[0].(bmcapisdk.DeleteResult)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerDelete indicates an expected call of ServerDelete.
func (mr *MockBmcApiSdkClientMockRecorder) ServerDelete(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerDelete", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerDelete), serverId)
}

// ServerGetById mocks base method.
func (m *MockBmcApiSdkClient) ServerGetById(serverId string) (bmcapisdk.Server, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerGetById", serverId)
	ret0, _ := ret[0].(bmcapisdk.Server)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerGetById indicates an expected call of ServerGetById.
func (mr *MockBmcApiSdkClientMockRecorder) ServerGetById(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerGetById", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerGetById), serverId)
}

// ServerPowerOff mocks base method.
func (m *MockBmcApiSdkClient) ServerPowerOff(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPowerOff", serverId)
	ret0, _ := ret[0].(bmcapisdk.ActionResult)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerPowerOff indicates an expected call of ServerPowerOff.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPowerOff(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPowerOff", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPowerOff), serverId)
}

// ServerPowerOn mocks base method.
func (m *MockBmcApiSdkClient) ServerPowerOn(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPowerOn", serverId)
	ret0, _ := ret[0].(bmcapisdk.ActionResult)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerPowerOn indicates an expected call of ServerPowerOn.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPowerOn(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPowerOn", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPowerOn), serverId)
}

// ServerReboot mocks base method.
func (m *MockBmcApiSdkClient) ServerReboot(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerReboot", serverId)
	ret0, _ := ret[0].(bmcapisdk.ActionResult)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerReboot indicates an expected call of ServerReboot.
func (mr *MockBmcApiSdkClientMockRecorder) ServerReboot(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerReboot", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerReboot), serverId)
}

// ServerReset mocks base method.
func (m *MockBmcApiSdkClient) ServerReset(serverId string, serverReset bmcapisdk.ServerReset) (bmcapisdk.ResetResult, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerReset", serverId, serverReset)
	ret0, _ := ret[0].(bmcapisdk.ResetResult)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerReset indicates an expected call of ServerReset.
func (mr *MockBmcApiSdkClientMockRecorder) ServerReset(serverId, serverReset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerReset", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerReset), serverId, serverReset)
}

// ServerShutdown mocks base method.
func (m *MockBmcApiSdkClient) ServerShutdown(serverId string) (bmcapisdk.ActionResult, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerShutdown", serverId)
	ret0, _ := ret[0].(bmcapisdk.ActionResult)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerShutdown indicates an expected call of ServerShutdown.
func (mr *MockBmcApiSdkClientMockRecorder) ServerShutdown(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerShutdown", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerShutdown), serverId)
}

// ServerReserve mocks base method.
func (m *MockBmcApiSdkClient) ServerReserve(serverId string, serverReserve bmcapisdk.ServerReserve) (bmcapisdk.Server, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerReserve", serverReserve)
	ret0, _ := ret[0].(bmcapisdk.Server)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerReserve indicates an expected call of ServerReserve.
func (mr *MockBmcApiSdkClientMockRecorder) ServerReserve(serverId, serverReserve interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerReserve", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerReserve), serverReserve)
}

// ServersGet mocks base method.
func (m *MockBmcApiSdkClient) ServersGet(tags []string) ([]bmcapisdk.Server, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServersGet")
	ret0, _ := ret[0].([]bmcapisdk.Server)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServersGet indicates an expected call of ServersGet.
func (mr *MockBmcApiSdkClientMockRecorder) ServersGet(tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServersGet", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServersGet))
}

// ServersPost mocks base method.
func (m *MockBmcApiSdkClient) ServersPost(serverCreate bmcapisdk.ServerCreate) (bmcapisdk.Server, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServersPost", serverCreate)
	ret0, _ := ret[0].(bmcapisdk.Server)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServersPost indicates an expected call of ServersPost.
func (mr *MockBmcApiSdkClientMockRecorder) ServersPost(serverCreate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServersPost", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServersPost), serverCreate)
}

// ServerPatch mocks base method.
func (m *MockBmcApiSdkClient) ServerPatch(serverId string, serverPatch bmcapisdk.ServerPatch) (bmcapisdk.Server, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPatch", serverPatch)
	ret0, _ := ret[0].(bmcapisdk.Server)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerPatch indicates an expected call of ServerPatch.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPatch(serverId, serverPatch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPatch", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPatch), serverPatch)
}

// ServerTag mocks base method.
func (m *MockBmcApiSdkClient) ServerTag(serverId string, tagAssignmentRequests []bmcapisdk.TagAssignmentRequest) (bmcapisdk.Server, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerTag", tagAssignmentRequests)
	ret0, _ := ret[0].(bmcapisdk.Server)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerTag indicates an expected call of ServerTag.
func (mr *MockBmcApiSdkClientMockRecorder) ServerTag(serverId, tagAssignmentRequests interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerTag", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerTag), tagAssignmentRequests)
}

// ServerPrivateNetworkPost mocks base method.
func (m *MockBmcApiSdkClient) ServerPrivateNetworkPost(serverId string, serverPrivateNetwork bmcapisdk.ServerPrivateNetwork) (bmcapisdk.ServerPrivateNetwork, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPrivateNetworkPost", serverPrivateNetwork)
	ret0, _ := ret[0].(bmcapisdk.ServerPrivateNetwork)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerPrivateNetworkPost indicates an expected call of ServerPrivateNetworkPost.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPrivateNetworkPost(serverId, serverPrivateNetwork interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPrivateNetworkPost", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPrivateNetworkPost), serverPrivateNetwork)
}

// ServerPrivateNetworkDelete mocks base method.
func (m *MockBmcApiSdkClient) ServerPrivateNetworkDelete(serverId string, networkId string) (string, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPrivateNetworkDelete", serverId, networkId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ServerPrivateNetworkDelete indicates an expected call of ServerPrivateNetworkDelete.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPrivateNetworkDelete(serverId, networkId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPrivateNetworkDelete", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPrivateNetworkDelete), serverId, networkId)
}

/*---- QUOTA -----*/
// QuotaGetById mocks base method.
func (m *MockBmcApiSdkClient) QuotaGetById(quotaId string) (bmcapisdk.Quota, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuotaGetById", quotaId)
	ret0, _ := ret[0].(bmcapisdk.Quota)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QuotaGetById indicates an expected call of ServerGetById.
func (mr *MockBmcApiSdkClientMockRecorder) QuotaGetById(quotaId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuotaGetById", reflect.TypeOf((*MockBmcApiSdkClient)(nil).QuotaGetById), quotaId)
}

// QuotasGet mocks base method.
func (m *MockBmcApiSdkClient) QuotasGet() ([]bmcapisdk.Quota, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuotasGet")
	ret0, _ := ret[0].([]bmcapisdk.Quota)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QuotasGet indicates an expected call of ServersGet.
func (mr *MockBmcApiSdkClientMockRecorder) QuotasGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuotasGet", reflect.TypeOf((*MockBmcApiSdkClient)(nil).QuotasGet))
}

// QuotasEdit mocks base method.
func (m *MockBmcApiSdkClient) QuotaEditById(quotaId string, quotaEditRequest bmcapisdk.QuotaEditLimitRequest) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuotaEditById")
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuotasEdit indicates an expected call of ServersGet.
func (mr *MockBmcApiSdkClientMockRecorder) QuotaEditById(quotaId interface{}, quotaEditRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuotaEditById", reflect.TypeOf((*MockBmcApiSdkClient)(nil).QuotaEditById))
}
