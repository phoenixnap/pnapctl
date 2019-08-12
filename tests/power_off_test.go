package tests

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/poweroff"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	mocks "phoenixnap.com/pnap-cli/pnapctl/mocks"

	"github.com/golang/mock/gomock"
)

// Each test needs to have a name like `TestXXX`
// They also need a parameter of `*testing.T`
func TestPowerOffServerSuccess(test_framework *testing.T) {
	// Creates a new mock controller.
	controller := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mockClient := mocks.NewMockWebClient(controller)

	// Manufacture response
	response := http.Response{
		StatusCode: 200,
	}

	// Sets main client to be the mock client.
	client.MainClient = mockClient

	// A call to the URL given should return the response and no errors.
	mockClient.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-off", bytes.NewBuffer([]byte{})).
		Return(&response, nil)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{serverID})
	if err != nil {
		test_framework.Errorf("Error detected: %s", err)
	}
}

func TestPowerOffServerConflict(test_framework *testing.T) {
	controller := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mockClient := mocks.NewMockWebClient(controller)

	// Manufacture response
	response := http.Response{
		StatusCode: 409,
	}

	// Change main client to be mock
	client.MainClient = mockClient

	// Mock return
	mockClient.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-off", bytes.NewBuffer([]byte{})).
		Return(&response, nil)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{serverID})
	if err.Error() != "409" {
		test_framework.Errorf("Expected '409 CONFLICT' error. Instead got %s", err)
	}
}

func TestPowerOffServerNotFound(test_framework *testing.T) {
	controller := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mockClient := mocks.NewMockWebClient(controller)

	// Formulate response
	response := http.Response{
		StatusCode: 404,
	}

	// Change main client to be mock
	client.MainClient = mockClient

	// Mock return
	mockClient.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-off", bytes.NewBuffer([]byte{})).
		Return(&response, nil)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{serverID})
	if err.Error() != "404" {
		test_framework.Errorf("Expected '404 NOT FOUND' error. Instead got %s", err)
	}
}

func TestPowerOffServerInternalServerError(test_framework *testing.T) {
	controller := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mockClient := mocks.NewMockWebClient(controller)

	// Formulate response
	response := http.Response{
		StatusCode: 500,
	}

	// Change main client to be mock
	client.MainClient = mockClient

	// Mock return
	mockClient.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-off", bytes.NewBuffer([]byte{})).
		Return(&response, nil)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{serverID})
	if err.Error() != "500" {
		test_framework.Errorf("Expected '500 INTERNAL SERVER ERROR' error. Instead got %s", err)
	}
}

func TestPowerOffServerTooManyArgs(test_framework *testing.T) {
	controller := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mockClient := mocks.NewMockWebClient(controller)

	// Formulate response
	response := http.Response{
		StatusCode: 200,
	}

	// Change main client to be mock
	client.MainClient = mockClient

	// Mock return
	mockClient.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-off", bytes.NewBuffer([]byte{})).
		Return(&response, nil)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{serverID, "extra"})
	if err.Error() != "args" {
		test_framework.Errorf("Expected 'too many args' error. Instead got %s", err)
	}
}

func TestPowerOffServerClientFailure(test_framework *testing.T) {
	controller := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mockClient := mocks.NewMockWebClient(controller)

	// Formulate response
	response := http.Response{
		StatusCode: 200,
	}

	// Change main client to mock
	client.MainClient = mockClient

	// Mock return
	mockClient.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-off", bytes.NewBuffer([]byte{})).
		Return(&response, errors.New("fake error"))

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{serverID})
	if err.Error() != "client-fail" {
		test_framework.Errorf("Expected 'client failure' error. Instead got %s", err)
	}
}
