package tests

import (
	"bytes"
	"net/http"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/poweron"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	mocks "phoenixnap.com/pnap-cli/pnapctl/mocks"

	// "phoenixnap.com/pnap-cli/pnapctl/mocks"
	"github.com/golang/mock/gomock"
)

func TestPowerOnServerSuccess(test_framework *testing.T) {
	ctrl := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mock_client := mocks.NewMockWebClient(ctrl)

	resp := http.Response{
		StatusCode: 200,
	}

	client.MainClient = mock_client

	mock_client.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-on", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{serverID})
	if err != nil {
		test_framework.Errorf("Error detected: %s", err)
	}
}

func TestPowerOnServerConflict(test_framework *testing.T) {
	ctrl := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mock_client := mocks.NewMockWebClient(ctrl)

	resp := http.Response{
		StatusCode: 409,
	}

	client.MainClient = mock_client

	mock_client.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-on", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{serverID})

	if err.Error() != "409" {
		test_framework.Errorf("Expected '409 CONFLICT' error. Instead got %s", err.Error())
	}
}

func TestPowerOnServerNotFound(test_framework *testing.T) {
	ctrl := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mock_client := mocks.NewMockWebClient(ctrl)

	resp := http.Response{
		StatusCode: 404,
	}

	client.MainClient = mock_client

	mock_client.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-on", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{serverID})

	if err.Error() != "404" {
		test_framework.Errorf("Expected '404 NOT FOUND' error. Instead got %s", err.Error())
	}
}

func TestPowerOnServerInternalServerError(test_framework *testing.T) {
	ctrl := gomock.NewController(test_framework)
	serverID := "fake_id"

	// Init mock client
	mock_client := mocks.NewMockWebClient(ctrl)

	resp := http.Response{
		StatusCode: 500,
	}

	client.MainClient = mock_client

	mock_client.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-on", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{serverID})

	if err.Error() != "500" {
		test_framework.Errorf("Expected '500 INTERNAL SERVER ERROR' error. Instead got %s", err.Error())
	}
}
