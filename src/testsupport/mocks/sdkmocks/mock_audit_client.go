// Code generated by MockGen. DO NOT EDIT.
// Source: common/client/audit/client.go

// Package sdkmocks is a generated GoMock package.
package sdkmocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	auditapi "github.com/phoenixnap/go-sdk-bmc/auditapi/v2"
	audit "phoenixnap.com/pnapctl/common/models/queryparams/audit"
	auditapi "github.com/phoenixnap/go-sdk-bmc/auditapi/v2"
	auditmodels "phoenixnap.com/pnapctl/common/models/auditmodels"
)

// MockAuditSdkClient is a mock of AuditSdkClient interface.
type MockAuditSdkClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuditSdkClientMockRecorder
}

// MockAuditSdkClientMockRecorder is the mock recorder for MockAuditSdkClient.
type MockAuditSdkClientMockRecorder struct {
	mock *MockAuditSdkClient
}

// NewMockAuditSdkClient creates a new mock instance.
func NewMockAuditSdkClient(ctrl *gomock.Controller) *MockAuditSdkClient {
	mock := &MockAuditSdkClient{ctrl: ctrl}
	mock.recorder = &MockAuditSdkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditSdkClient) EXPECT() *MockAuditSdkClientMockRecorder {
	return m.recorder
}

// EventsGet mocks base method.
func (m *MockAuditSdkClient) EventsGet(queryParams audit.EventsGetQueryParams) ([]auditapi.Event, *http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsGet", queryParams)
	ret0, _ := ret[0].([]auditapi.Event)
	ret1, _ := ret[1].(*http.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EventsGet indicates an expected call of EventsGet.
func (mr *MockAuditSdkClientMockRecorder) EventsGet(queryParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsGet", reflect.TypeOf((*MockAuditSdkClient)(nil).EventsGet), queryParams)
}
