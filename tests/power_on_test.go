package tests

import (
	"bytes"
	"testing"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/poweron"
	// "phoenixnap.com/pnap-cli/pnapctl/mocks"
)

var body = bytes.NewBuffer([]byte{})

const serverID = "mock_id"
const url = "servers/" + serverID + "/actions/power-on"

func TestPowerOnServerSuccess(test_framework *testing.T) {
	// Init mock client
	PrepareMockClient(test_framework).
		PerformPost(url, body).
		Return(WithResponse(200, nil), nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{serverID})

	if err != nil {
		test_framework.Errorf("Expected no error. Instead got %s", err.Error())
	}
}

func TestPowerOnServerConflict(test_framework *testing.T) {
	// init mock client
	PrepareMockClient(test_framework).
		PerformPost(url, body).
		Return(WithResponse(409, nil), nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{serverID})

	if err.Error() != "409" {
		test_framework.Errorf("Expected '409 CONFLICT' error. Instead got %s", err.Error())
	}
}

func TestPowerOnServerNotFound(test_framework *testing.T) {
	// init
	PrepareMockClient(test_framework).
		PerformPost(url, body).
		Return(WithResponse(404, nil), nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{serverID})

	if err.Error() != "404" {
		test_framework.Errorf("Expected '404 NOT FOUND' error. Instead got %s", err.Error())
	}
}

func TestPowerOnServerInternalServerError(test_framework *testing.T) {
	PrepareMockClient(test_framework).
		PerformPost(url, body).
		Return(WithResponse(500, nil), nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{serverID})

	if err.Error() != "500" {
		test_framework.Errorf("Expected '500 INTERNAL SERVER ERROR' error. Instead got %s", err.Error())
	}
}
