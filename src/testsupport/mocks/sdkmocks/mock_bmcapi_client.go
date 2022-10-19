// Code generated by MockGen. DO NOT EDIT.
// Source: common/client/bmcapi/client.go

// Package sdkmocks is a generated GoMock package.
package sdkmocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	bmcapi "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
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

// QuotaEditById mocks base method.
func (m *MockBmcApiSdkClient) QuotaEditById(quotaId string, quotaEditRequest bmcapi.QuotaEditLimitRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuotaEditById", quotaId, quotaEditRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// QuotaEditById indicates an expected call of QuotaEditById.
func (mr *MockBmcApiSdkClientMockRecorder) QuotaEditById(quotaId, quotaEditRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuotaEditById", reflect.TypeOf((*MockBmcApiSdkClient)(nil).QuotaEditById), quotaId, quotaEditRequest)
}

// QuotaGetById mocks base method.
func (m *MockBmcApiSdkClient) QuotaGetById(quotaId string) (*bmcapi.Quota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuotaGetById", quotaId)
	ret0, _ := ret[0].(*bmcapi.Quota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuotaGetById indicates an expected call of QuotaGetById.
func (mr *MockBmcApiSdkClientMockRecorder) QuotaGetById(quotaId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuotaGetById", reflect.TypeOf((*MockBmcApiSdkClient)(nil).QuotaGetById), quotaId)
}

// QuotasGet mocks base method.
func (m *MockBmcApiSdkClient) QuotasGet() ([]bmcapi.Quota, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuotasGet")
	ret0, _ := ret[0].([]bmcapi.Quota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuotasGet indicates an expected call of QuotasGet.
func (mr *MockBmcApiSdkClientMockRecorder) QuotasGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuotasGet", reflect.TypeOf((*MockBmcApiSdkClient)(nil).QuotasGet))
}

// ServerDelete mocks base method.
func (m *MockBmcApiSdkClient) ServerDelete(serverId string) (*bmcapi.DeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerDelete", serverId)
	ret0, _ := ret[0].(*bmcapi.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerDelete indicates an expected call of ServerDelete.
func (mr *MockBmcApiSdkClientMockRecorder) ServerDelete(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerDelete", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerDelete), serverId)
}

// ServerDeprovision mocks base method.
func (m *MockBmcApiSdkClient) ServerDeprovision(serverId string, relinquishIpBlock bmcapi.RelinquishIpBlock) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerDeprovision", serverId, relinquishIpBlock)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerDeprovision indicates an expected call of ServerDeprovision.
func (mr *MockBmcApiSdkClientMockRecorder) ServerDeprovision(serverId, relinquishIpBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerDeprovision", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerDeprovision), serverId, relinquishIpBlock)
}

// ServerGetById mocks base method.
func (m *MockBmcApiSdkClient) ServerGetById(serverId string) (*bmcapi.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerGetById", serverId)
	ret0, _ := ret[0].(*bmcapi.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerGetById indicates an expected call of ServerGetById.
func (mr *MockBmcApiSdkClientMockRecorder) ServerGetById(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerGetById", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerGetById), serverId)
}

// ServerIpBlockDelete mocks base method.
func (m *MockBmcApiSdkClient) ServerIpBlockDelete(serverId, ipBlockId string, relinquishIpBlock bmcapi.RelinquishIpBlock) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerIpBlockDelete", serverId, ipBlockId, relinquishIpBlock)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerIpBlockDelete indicates an expected call of ServerIpBlockDelete.
func (mr *MockBmcApiSdkClientMockRecorder) ServerIpBlockDelete(serverId, ipBlockId, relinquishIpBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerIpBlockDelete", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerIpBlockDelete), serverId, ipBlockId, relinquishIpBlock)
}

// ServerIpBlockPost mocks base method.
func (m *MockBmcApiSdkClient) ServerIpBlockPost(serverId string, serverIpBlock bmcapi.ServerIpBlock) (*bmcapi.ServerIpBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerIpBlockPost", serverId, serverIpBlock)
	ret0, _ := ret[0].(*bmcapi.ServerIpBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerIpBlockPost indicates an expected call of ServerIpBlockPost.
func (mr *MockBmcApiSdkClientMockRecorder) ServerIpBlockPost(serverId, serverIpBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerIpBlockPost", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerIpBlockPost), serverId, serverIpBlock)
}

// ServerPatch mocks base method.
func (m *MockBmcApiSdkClient) ServerPatch(serverId string, serverPatch bmcapi.ServerPatch) (*bmcapi.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPatch", serverId, serverPatch)
	ret0, _ := ret[0].(*bmcapi.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPatch indicates an expected call of ServerPatch.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPatch(serverId, serverPatch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPatch", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPatch), serverId, serverPatch)
}

// ServerPowerOff mocks base method.
func (m *MockBmcApiSdkClient) ServerPowerOff(serverId string) (*bmcapi.ActionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPowerOff", serverId)
	ret0, _ := ret[0].(*bmcapi.ActionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPowerOff indicates an expected call of ServerPowerOff.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPowerOff(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPowerOff", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPowerOff), serverId)
}

// ServerPowerOn mocks base method.
func (m *MockBmcApiSdkClient) ServerPowerOn(serverId string) (*bmcapi.ActionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPowerOn", serverId)
	ret0, _ := ret[0].(*bmcapi.ActionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPowerOn indicates an expected call of ServerPowerOn.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPowerOn(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPowerOn", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPowerOn), serverId)
}

// ServerPrivateNetworkDelete mocks base method.
func (m *MockBmcApiSdkClient) ServerPrivateNetworkDelete(serverId, networkId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPrivateNetworkDelete", serverId, networkId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPrivateNetworkDelete indicates an expected call of ServerPrivateNetworkDelete.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPrivateNetworkDelete(serverId, networkId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPrivateNetworkDelete", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPrivateNetworkDelete), serverId, networkId)
}

// ServerPrivateNetworkPost mocks base method.
func (m *MockBmcApiSdkClient) ServerPrivateNetworkPost(serverId string, serverPrivateNetwork bmcapi.ServerPrivateNetwork) (*bmcapi.ServerPrivateNetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPrivateNetworkPost", serverId, serverPrivateNetwork)
	ret0, _ := ret[0].(*bmcapi.ServerPrivateNetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPrivateNetworkPost indicates an expected call of ServerPrivateNetworkPost.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPrivateNetworkPost(serverId, serverPrivateNetwork interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPrivateNetworkPost", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPrivateNetworkPost), serverId, serverPrivateNetwork)
}

// ServerPublicNetworkDelete mocks base method.
func (m *MockBmcApiSdkClient) ServerPublicNetworkDelete(serverId, networkId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPublicNetworkDelete", serverId, networkId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPublicNetworkDelete indicates an expected call of ServerPublicNetworkDelete.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPublicNetworkDelete(serverId, networkId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPublicNetworkDelete", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPublicNetworkDelete), serverId, networkId)
}

// ServerPublicNetworkPost mocks base method.
func (m *MockBmcApiSdkClient) ServerPublicNetworkPost(serverId string, serverPublicNetwork bmcapi.ServerPublicNetwork) (*bmcapi.ServerPublicNetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPublicNetworkPost", serverId, serverPublicNetwork)
	ret0, _ := ret[0].(*bmcapi.ServerPublicNetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerPublicNetworkPost indicates an expected call of ServerPublicNetworkPost.
func (mr *MockBmcApiSdkClientMockRecorder) ServerPublicNetworkPost(serverId, serverPublicNetwork interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPublicNetworkPost", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerPublicNetworkPost), serverId, serverPublicNetwork)
}

// ServerReboot mocks base method.
func (m *MockBmcApiSdkClient) ServerReboot(serverId string) (*bmcapi.ActionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerReboot", serverId)
	ret0, _ := ret[0].(*bmcapi.ActionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerReboot indicates an expected call of ServerReboot.
func (mr *MockBmcApiSdkClientMockRecorder) ServerReboot(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerReboot", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerReboot), serverId)
}

// ServerReserve mocks base method.
func (m *MockBmcApiSdkClient) ServerReserve(serverId string, serverReserve bmcapi.ServerReserve) (*bmcapi.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerReserve", serverId, serverReserve)
	ret0, _ := ret[0].(*bmcapi.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerReserve indicates an expected call of ServerReserve.
func (mr *MockBmcApiSdkClientMockRecorder) ServerReserve(serverId, serverReserve interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerReserve", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerReserve), serverId, serverReserve)
}

// ServerReset mocks base method.
func (m *MockBmcApiSdkClient) ServerReset(serverId string, serverReset bmcapi.ServerReset) (*bmcapi.ResetResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerReset", serverId, serverReset)
	ret0, _ := ret[0].(*bmcapi.ResetResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerReset indicates an expected call of ServerReset.
func (mr *MockBmcApiSdkClientMockRecorder) ServerReset(serverId, serverReset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerReset", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerReset), serverId, serverReset)
}

// ServerShutdown mocks base method.
func (m *MockBmcApiSdkClient) ServerShutdown(serverId string) (*bmcapi.ActionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerShutdown", serverId)
	ret0, _ := ret[0].(*bmcapi.ActionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerShutdown indicates an expected call of ServerShutdown.
func (mr *MockBmcApiSdkClientMockRecorder) ServerShutdown(serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerShutdown", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerShutdown), serverId)
}

// ServerTag mocks base method.
func (m *MockBmcApiSdkClient) ServerTag(serverId string, tagAssignmentRequests []bmcapi.TagAssignmentRequest) (*bmcapi.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerTag", serverId, tagAssignmentRequests)
	ret0, _ := ret[0].(*bmcapi.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServerTag indicates an expected call of ServerTag.
func (mr *MockBmcApiSdkClientMockRecorder) ServerTag(serverId, tagAssignmentRequests interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerTag", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServerTag), serverId, tagAssignmentRequests)
}

// ServersGet mocks base method.
func (m *MockBmcApiSdkClient) ServersGet(arg0 []string) ([]bmcapi.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServersGet", arg0)
	ret0, _ := ret[0].([]bmcapi.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServersGet indicates an expected call of ServersGet.
func (mr *MockBmcApiSdkClientMockRecorder) ServersGet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServersGet", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServersGet), arg0)
}

// ServersPost mocks base method.
func (m *MockBmcApiSdkClient) ServersPost(serverCreate bmcapi.ServerCreate) (*bmcapi.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServersPost", serverCreate)
	ret0, _ := ret[0].(*bmcapi.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServersPost indicates an expected call of ServersPost.
func (mr *MockBmcApiSdkClientMockRecorder) ServersPost(serverCreate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServersPost", reflect.TypeOf((*MockBmcApiSdkClient)(nil).ServersPost), serverCreate)
}

// SshKeyDelete mocks base method.
func (m *MockBmcApiSdkClient) SshKeyDelete(sshKeyId string) (*bmcapi.DeleteSshKeyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SshKeyDelete", sshKeyId)
	ret0, _ := ret[0].(*bmcapi.DeleteSshKeyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SshKeyDelete indicates an expected call of SshKeyDelete.
func (mr *MockBmcApiSdkClientMockRecorder) SshKeyDelete(sshKeyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SshKeyDelete", reflect.TypeOf((*MockBmcApiSdkClient)(nil).SshKeyDelete), sshKeyId)
}

// SshKeyGetById mocks base method.
func (m *MockBmcApiSdkClient) SshKeyGetById(sshKeyId string) (*bmcapi.SshKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SshKeyGetById", sshKeyId)
	ret0, _ := ret[0].(*bmcapi.SshKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SshKeyGetById indicates an expected call of SshKeyGetById.
func (mr *MockBmcApiSdkClientMockRecorder) SshKeyGetById(sshKeyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SshKeyGetById", reflect.TypeOf((*MockBmcApiSdkClient)(nil).SshKeyGetById), sshKeyId)
}

// SshKeyPost mocks base method.
func (m *MockBmcApiSdkClient) SshKeyPost(sshkeyCreate bmcapi.SshKeyCreate) (*bmcapi.SshKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SshKeyPost", sshkeyCreate)
	ret0, _ := ret[0].(*bmcapi.SshKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SshKeyPost indicates an expected call of SshKeyPost.
func (mr *MockBmcApiSdkClientMockRecorder) SshKeyPost(sshkeyCreate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SshKeyPost", reflect.TypeOf((*MockBmcApiSdkClient)(nil).SshKeyPost), sshkeyCreate)
}

// SshKeyPut mocks base method.
func (m *MockBmcApiSdkClient) SshKeyPut(sshKeyId string, sshKeyUpdate bmcapi.SshKeyUpdate) (*bmcapi.SshKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SshKeyPut", sshKeyId, sshKeyUpdate)
	ret0, _ := ret[0].(*bmcapi.SshKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SshKeyPut indicates an expected call of SshKeyPut.
func (mr *MockBmcApiSdkClientMockRecorder) SshKeyPut(sshKeyId, sshKeyUpdate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SshKeyPut", reflect.TypeOf((*MockBmcApiSdkClient)(nil).SshKeyPut), sshKeyId, sshKeyUpdate)
}

// SshKeysGet mocks base method.
func (m *MockBmcApiSdkClient) SshKeysGet() ([]bmcapi.SshKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SshKeysGet")
	ret0, _ := ret[0].([]bmcapi.SshKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SshKeysGet indicates an expected call of SshKeysGet.
func (mr *MockBmcApiSdkClientMockRecorder) SshKeysGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SshKeysGet", reflect.TypeOf((*MockBmcApiSdkClient)(nil).SshKeysGet))
}
