// Code generated by MockGen. DO NOT EDIT.
// Source: common/client/ip/client.go

// Package sdkmocks is a generated GoMock package.
package sdkmocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ipapi "github.com/phoenixnap/go-sdk-bmc/ipapi"
)

// MockIpSdkClient is a mock of IpSdkClient interface.
type MockIpSdkClient struct {
	ctrl     *gomock.Controller
	recorder *MockIpSdkClientMockRecorder
}

// MockIpSdkClientMockRecorder is the mock recorder for MockIpSdkClient.
type MockIpSdkClientMockRecorder struct {
	mock *MockIpSdkClient
}

// NewMockIpSdkClient creates a new mock instance.
func NewMockIpSdkClient(ctrl *gomock.Controller) *MockIpSdkClient {
	mock := &MockIpSdkClient{ctrl: ctrl}
	mock.recorder = &MockIpSdkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIpSdkClient) EXPECT() *MockIpSdkClientMockRecorder {
	return m.recorder
}

// IpBlockPost mocks base method.
func (m *MockIpSdkClient) IpBlockPost(ipBlockCreate ipapi.IpBlockCreate) (*ipapi.IpBlock, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IpBlockPost", ipBlockCreate)
	ret0, _ := ret[0].(*ipapi.IpBlock)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IpBlockPost indicates an expected call of IpBlockPost.
func (mr *MockIpSdkClientMockRecorder) IpBlockPost(ipBlockCreate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpBlockPost", reflect.TypeOf((*MockIpSdkClient)(nil).IpBlockPost), ipBlockCreate)
}

// IpBlocksGet mocks base method.
func (m *MockIpSdkClient) IpBlocksGet(arg0 []string) ([]ipapi.IpBlock, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IpBlocksGet", arg0)
	ret0, _ := ret[0].([]ipapi.IpBlock)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IpBlocksGet indicates an expected call of IpBlocksGet.
func (mr *MockIpSdkClientMockRecorder) IpBlocksGet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpBlocksGet", reflect.TypeOf((*MockIpSdkClient)(nil).IpBlocksGet), arg0)
}

// IpBlocksGetById mocks base method.
func (m *MockIpSdkClient) IpBlocksGetById(ipBlockId string) (*ipapi.IpBlock, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IpBlocksGetById", ipBlockId)
	ret0, _ := ret[0].(*ipapi.IpBlock)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IpBlocksGetById indicates an expected call of IpBlocksGetById.
func (mr *MockIpSdkClientMockRecorder) IpBlocksGetById(ipBlockId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpBlocksGetById", reflect.TypeOf((*MockIpSdkClient)(nil).IpBlocksGetById), ipBlockId)
}

// IpBlocksIpBlockIdDelete mocks base method.
func (m *MockIpSdkClient) IpBlocksIpBlockIdDelete(ipBlockId string) (*ipapi.DeleteIpBlockResult, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IpBlocksIpBlockIdDelete", ipBlockId)
	ret0, _ := ret[0].(*ipapi.DeleteIpBlockResult)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IpBlocksIpBlockIdDelete indicates an expected call of IpBlocksIpBlockIdDelete.
func (mr *MockIpSdkClientMockRecorder) IpBlocksIpBlockIdDelete(ipBlockId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpBlocksIpBlockIdDelete", reflect.TypeOf((*MockIpSdkClient)(nil).IpBlocksIpBlockIdDelete), ipBlockId)
}

// IpBlocksIpBlockIdPatch mocks base method.
func (m *MockIpSdkClient) IpBlocksIpBlockIdPatch(ipBlockId string, ipBlockPatch ipapi.IpBlockPatch) (*ipapi.IpBlock, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IpBlocksIpBlockIdPatch", ipBlockId, ipBlockPatch)
	ret0, _ := ret[0].(*ipapi.IpBlock)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IpBlocksIpBlockIdPatch indicates an expected call of IpBlocksIpBlockIdPatch.
func (mr *MockIpSdkClientMockRecorder) IpBlocksIpBlockIdPatch(ipBlockId, ipBlockPatch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpBlocksIpBlockIdPatch", reflect.TypeOf((*MockIpSdkClient)(nil).IpBlocksIpBlockIdPatch), ipBlockId, ipBlockPatch)
}

// IpBlocksIpBlockIdTagsPut mocks base method.
func (m *MockIpSdkClient) IpBlocksIpBlockIdTagsPut(ipBlockId string, tag []ipapi.TagAssignmentRequest) (*ipapi.IpBlock, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IpBlocksIpBlockIdTagsPut", ipBlockId, tag)
	ret0, _ := ret[0].(*ipapi.IpBlock)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IpBlocksIpBlockIdTagsPut indicates an expected call of IpBlocksIpBlockIdTagsPut.
func (mr *MockIpSdkClientMockRecorder) IpBlocksIpBlockIdTagsPut(ipBlockId, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpBlocksIpBlockIdTagsPut", reflect.TypeOf((*MockIpSdkClient)(nil).IpBlocksIpBlockIdTagsPut), ipBlockId, tag)
}
