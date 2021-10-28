// Code generated by MockGen. DO NOT EDIT.
// Source: ./common/printer/printer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPrinter is a mock of Printer interface.
type MockPrinter struct {
	ctrl     *gomock.Controller
	recorder *MockPrinterMockRecorder
}

// MockPrinterMockRecorder is the mock recorder for MockPrinter.
type MockPrinterMockRecorder struct {
	mock *MockPrinter
}

// NewMockPrinter creates a new mock instance.
func NewMockPrinter(ctrl *gomock.Controller) *MockPrinter {
	mock := &MockPrinter{ctrl: ctrl}
	mock.recorder = &MockPrinterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrinter) EXPECT() *MockPrinterMockRecorder {
	return m.recorder
}

// PrintOutput mocks base method.
func (m *MockPrinter) PrintOutput(construct interface{}, commandName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrintOutput", construct, commandName)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrintOutput indicates an expected call of PrintOutput.
func (mr *MockPrinterMockRecorder) PrintOutput(construct, commandName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintOutput", reflect.TypeOf((*MockPrinter)(nil).PrintOutput), construct, commandName)
}
