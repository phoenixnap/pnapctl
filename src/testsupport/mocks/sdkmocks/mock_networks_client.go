// Code generated by MockGen. DO NOT EDIT.
// Source: common/client/networks/client.go

// Package sdkmocks is a generated GoMock package.
package sdkmocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	networkapi "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
)

// MockNetworkSdkClient is a mock of NetworkSdkClient interface.
type MockNetworkSdkClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkSdkClientMockRecorder
}

// MockNetworkSdkClientMockRecorder is the mock recorder for MockNetworkSdkClient.
type MockNetworkSdkClientMockRecorder struct {
	mock *MockNetworkSdkClient
}

// NewMockNetworkSdkClient creates a new mock instance.
func NewMockNetworkSdkClient(ctrl *gomock.Controller) *MockNetworkSdkClient {
	mock := &MockNetworkSdkClient{ctrl: ctrl}
	mock.recorder = &MockNetworkSdkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkSdkClient) EXPECT() *MockNetworkSdkClientMockRecorder {
	return m.recorder
}

// PrivateNetworkDelete mocks base method.
func (m *MockNetworkSdkClient) PrivateNetworkDelete(networkId string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateNetworkDelete", networkId)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateNetworkDelete indicates an expected call of PrivateNetworkDelete.
func (mr *MockNetworkSdkClientMockRecorder) PrivateNetworkDelete(networkId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateNetworkDelete", reflect.TypeOf((*MockNetworkSdkClient)(nil).PrivateNetworkDelete), networkId)
}

// PrivateNetworkGetById mocks base method.
func (m *MockNetworkSdkClient) PrivateNetworkGetById(networkId string) (*networkapi.PrivateNetwork, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateNetworkGetById", networkId)
	ret0, _ := ret[0].(*networkapi.PrivateNetwork)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PrivateNetworkGetById indicates an expected call of PrivateNetworkGetById.
func (mr *MockNetworkSdkClientMockRecorder) PrivateNetworkGetById(networkId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateNetworkGetById", reflect.TypeOf((*MockNetworkSdkClient)(nil).PrivateNetworkGetById), networkId)
}

// PrivateNetworkPut mocks base method.
func (m *MockNetworkSdkClient) PrivateNetworkPut(networkId string, privateNetworkUpdate networkapi.PrivateNetworkModify) (*networkapi.PrivateNetwork, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateNetworkPut", networkId, privateNetworkUpdate)
	ret0, _ := ret[0].(*networkapi.PrivateNetwork)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PrivateNetworkPut indicates an expected call of PrivateNetworkPut.
func (mr *MockNetworkSdkClientMockRecorder) PrivateNetworkPut(networkId, privateNetworkUpdate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateNetworkPut", reflect.TypeOf((*MockNetworkSdkClient)(nil).PrivateNetworkPut), networkId, privateNetworkUpdate)
}

// PrivateNetworksGet mocks base method.
func (m *MockNetworkSdkClient) PrivateNetworksGet(location string) ([]networkapi.PrivateNetwork, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateNetworksGet", location)
	ret0, _ := ret[0].([]networkapi.PrivateNetwork)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PrivateNetworksGet indicates an expected call of PrivateNetworksGet.
func (mr *MockNetworkSdkClientMockRecorder) PrivateNetworksGet(location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateNetworksGet", reflect.TypeOf((*MockNetworkSdkClient)(nil).PrivateNetworksGet), location)
}

// PrivateNetworksPost mocks base method.
func (m *MockNetworkSdkClient) PrivateNetworksPost(privateNetworkCreate networkapi.PrivateNetworkCreate) (*networkapi.PrivateNetwork, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateNetworksPost", privateNetworkCreate)
	ret0, _ := ret[0].(*networkapi.PrivateNetwork)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PrivateNetworksPost indicates an expected call of PrivateNetworksPost.
func (mr *MockNetworkSdkClientMockRecorder) PrivateNetworksPost(privateNetworkCreate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateNetworksPost", reflect.TypeOf((*MockNetworkSdkClient)(nil).PrivateNetworksPost), privateNetworkCreate)
}

// PublicNetworkDelete mocks base method.
func (m *MockNetworkSdkClient) PublicNetworkDelete(networkId string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicNetworkDelete", networkId)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublicNetworkDelete indicates an expected call of PublicNetworkDelete.
func (mr *MockNetworkSdkClientMockRecorder) PublicNetworkDelete(networkId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicNetworkDelete", reflect.TypeOf((*MockNetworkSdkClient)(nil).PublicNetworkDelete), networkId)
}

// PublicNetworkGetById mocks base method.
func (m *MockNetworkSdkClient) PublicNetworkGetById(networkId string) (*networkapi.PublicNetwork, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicNetworkGetById", networkId)
	ret0, _ := ret[0].(*networkapi.PublicNetwork)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PublicNetworkGetById indicates an expected call of PublicNetworkGetById.
func (mr *MockNetworkSdkClientMockRecorder) PublicNetworkGetById(networkId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicNetworkGetById", reflect.TypeOf((*MockNetworkSdkClient)(nil).PublicNetworkGetById), networkId)
}

// PublicNetworkIpBlockDelete mocks base method.
func (m *MockNetworkSdkClient) PublicNetworkIpBlockDelete(networkId, ipBlockId string) (string, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicNetworkIpBlockDelete", networkId, ipBlockId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PublicNetworkIpBlockDelete indicates an expected call of PublicNetworkIpBlockDelete.
func (mr *MockNetworkSdkClientMockRecorder) PublicNetworkIpBlockDelete(networkId, ipBlockId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicNetworkIpBlockDelete", reflect.TypeOf((*MockNetworkSdkClient)(nil).PublicNetworkIpBlockDelete), networkId, ipBlockId)
}

// PublicNetworkIpBlockPost mocks base method.
func (m *MockNetworkSdkClient) PublicNetworkIpBlockPost(networkId string, idBlockCreate networkapi.PublicNetworkIpBlock) (*networkapi.PublicNetworkIpBlock, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicNetworkIpBlockPost", networkId, idBlockCreate)
	ret0, _ := ret[0].(*networkapi.PublicNetworkIpBlock)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PublicNetworkIpBlockPost indicates an expected call of PublicNetworkIpBlockPost.
func (mr *MockNetworkSdkClientMockRecorder) PublicNetworkIpBlockPost(networkId, idBlockCreate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicNetworkIpBlockPost", reflect.TypeOf((*MockNetworkSdkClient)(nil).PublicNetworkIpBlockPost), networkId, idBlockCreate)
}

// PublicNetworkPatch mocks base method.
func (m *MockNetworkSdkClient) PublicNetworkPatch(networkId string, publicNetworkPatch networkapi.PublicNetworkModify) (*networkapi.PublicNetwork, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicNetworkPatch", networkId, publicNetworkPatch)
	ret0, _ := ret[0].(*networkapi.PublicNetwork)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PublicNetworkPatch indicates an expected call of PublicNetworkPatch.
func (mr *MockNetworkSdkClientMockRecorder) PublicNetworkPatch(networkId, publicNetworkPatch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicNetworkPatch", reflect.TypeOf((*MockNetworkSdkClient)(nil).PublicNetworkPatch), networkId, publicNetworkPatch)
}

// PublicNetworksGet mocks base method.
func (m *MockNetworkSdkClient) PublicNetworksGet(arg0 string) ([]networkapi.PublicNetwork, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicNetworksGet", arg0)
	ret0, _ := ret[0].([]networkapi.PublicNetwork)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PublicNetworksGet indicates an expected call of PublicNetworksGet.
func (mr *MockNetworkSdkClientMockRecorder) PublicNetworksGet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicNetworksGet", reflect.TypeOf((*MockNetworkSdkClient)(nil).PublicNetworksGet), arg0)
}

// PublicNetworksPost mocks base method.
func (m *MockNetworkSdkClient) PublicNetworksPost(publicNetworkCreate networkapi.PublicNetworkCreate) (*networkapi.PublicNetwork, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicNetworksPost", publicNetworkCreate)
	ret0, _ := ret[0].(*networkapi.PublicNetwork)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PublicNetworksPost indicates an expected call of PublicNetworksPost.
func (mr *MockNetworkSdkClientMockRecorder) PublicNetworksPost(publicNetworkCreate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicNetworksPost", reflect.TypeOf((*MockNetworkSdkClient)(nil).PublicNetworksPost), publicNetworkCreate)
}
