package tests

import (
	"bytes"
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
