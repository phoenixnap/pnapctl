// Code generated by MockGen. DO NOT EDIT.
// Source: common/client/locations/client.go

// Package sdkmocks is a generated GoMock package.
package sdkmocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	locationapi "github.com/phoenixnap/go-sdk-bmc/locationapi/v2"
)

// MockLocationSdkClient is a mock of LocationSdkClient interface.
type MockLocationSdkClient struct {
	ctrl     *gomock.Controller
	recorder *MockLocationSdkClientMockRecorder
}

// MockLocationSdkClientMockRecorder is the mock recorder for MockLocationSdkClient.
type MockLocationSdkClientMockRecorder struct {
	mock *MockLocationSdkClient
}

// NewMockLocationSdkClient creates a new mock instance.
func NewMockLocationSdkClient(ctrl *gomock.Controller) *MockLocationSdkClient {
	mock := &MockLocationSdkClient{ctrl: ctrl}
	mock.recorder = &MockLocationSdkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocationSdkClient) EXPECT() *MockLocationSdkClientMockRecorder {
	return m.recorder
}

// LocationsGet mocks base method.
func (m *MockLocationSdkClient) LocationsGet(location, productCategory string) ([]locationapi.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocationsGet", location, productCategory)
	ret0, _ := ret[0].([]locationapi.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocationsGet indicates an expected call of LocationsGet.
func (mr *MockLocationSdkClientMockRecorder) LocationsGet(location, productCategory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocationsGet", reflect.TypeOf((*MockLocationSdkClient)(nil).LocationsGet), location, productCategory)
}
