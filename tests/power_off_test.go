package tests

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl"

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

	// Prepare arguments to be used by Cobra
	os.Args = []string{"pnapctl", "bmc", "power-off", serverID}
	pnapctl.Execute()
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

	// Call
	os.Args = []string{"pnapctl", "bmc", "power-off", serverID}
	pnapctl.Execute()

	// Expect the error code to be set right.
	// ErrorCode is our own variable, made to log the status of the request.
	if poweroff.ErrorCode != "409" {
		// Errorf displays the error and fails the test.
		test_framework.Errorf("Expected '409 CONFLICT' error. Instead got %s", poweroff.ErrorCode)
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

	// Call
	os.Args = []string{"pnapctl", "bmc", "power-off", serverID}
	pnapctl.Execute()

	// Assert
	if poweroff.ErrorCode != "404" {
		test_framework.Errorf("Expected '404 NOT FOUND' error. Instead got %s", poweroff.ErrorCode)
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

	// Call
	os.Args = []string{"pnapctl", "bmc", "power-off", serverID}
	pnapctl.Execute()

	// Assert
	if poweroff.ErrorCode != "500" {
		test_framework.Errorf("Expected '500 INTERNAL SERVER ERROR' error. Instead got %s", poweroff.ErrorCode)
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

	// Call
	os.Args = []string{"pnapctl", "bmc", "power-off", serverID, "extra"}
	pnapctl.Execute()

	// Assert
	if poweroff.ErrorCode != "ARGS" {
		test_framework.Errorf("Expected 'too many args' error. Instead got %s", poweroff.ErrorCode)
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

	// Call
	os.Args = []string{"pnapctl", "bmc", "power-off", serverID}
	pnapctl.Execute()

	// Assert
	if poweroff.ErrorCode != "CLIENT" {
		test_framework.Errorf("Expected 'Client failed' error. Instead got %s", poweroff.ErrorCode)
	}
}
