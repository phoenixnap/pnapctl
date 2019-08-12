package tests

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnap-cli/pnapctl/bmc/reboot"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	mocks "phoenixnap.com/pnap-cli/pnapctl/mocks"
)

func TestRebootServerSuccess(test_framework *testing.T) {
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

	// Mocks call to PerformPost
	mockClient.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/reboot", bytes.NewBuffer([]byte{})).
		Return(&response, nil)

	// Run command
	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{serverID})
	if err != nil {
		test_framework.Errorf("Error detected: %s", err)
	}
}

func TestRebootServerArgFail(test_framework *testing.T) {
	serverID := "fake_id"

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{serverID, "extra"})
	if err.Error() != "args" {
		test_framework.Errorf("Expected invalid args error - found %s", err)
	}
}

func TestRebootServerClientFail(test_framework *testing.T) {
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

	// Mocks call to PerformPost
	mockClient.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/reboot", bytes.NewBuffer([]byte{})).
		Return(&response, errors.New("fake"))

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{serverID})
	if err.Error() != "client-fail" {
		test_framework.Errorf("Error: Expected client failure error, found %s", err)
	}
}

func TestRebootServerNotFoundFail(test_framework *testing.T) {
	// Creates a new mock controller.
	controller := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mockClient := mocks.NewMockWebClient(controller)

	// Manufacture response
	response := http.Response{
		StatusCode: 404,
	}

	// Sets main client to be the mock client.
	client.MainClient = mockClient

	// Mocks call to PerformPost
	mockClient.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/reboot", bytes.NewBuffer([]byte{})).
		Return(&response, nil)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{serverID})
	if err.Error() != "404" {
		test_framework.Errorf("Error: not found error, found %s", err)
	}
}

func TestRebootServerConflictFail(test_framework *testing.T) {
	// Creates a new mock controller.
	controller := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mockClient := mocks.NewMockWebClient(controller)

	// Manufacture response
	response := http.Response{
		StatusCode: 409,
	}

	// Sets main client to be the mock client.
	client.MainClient = mockClient

	// Mocks call to PerformPost
	mockClient.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/reboot", bytes.NewBuffer([]byte{})).
		Return(&response, nil)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{serverID})
	if err.Error() != "409" {
		test_framework.Errorf("Error: Expected conflict error, found %s", err)
	}
}

func TestRebootServerInternalServerErrorFail(test_framework *testing.T) {
	// Creates a new mock controller.
	controller := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mockClient := mocks.NewMockWebClient(controller)

	// Manufacture response
	response := http.Response{
		StatusCode: 500,
	}

	// Sets main client to be the mock client.
	client.MainClient = mockClient

	// Mocks call to PerformPost
	mockClient.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/reboot", bytes.NewBuffer([]byte{})).
		Return(&response, nil)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{serverID})
	if err.Error() != "500" {
		test_framework.Errorf("Error: Expected internal server error, found %s", err)
	}
}
