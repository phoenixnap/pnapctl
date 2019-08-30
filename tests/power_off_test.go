package tests

import (
	"bytes"
	"errors"
	"testing"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/poweroff"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

func TestPowerOffSetup(t *testing.T) {
	Body = bytes.NewBuffer([]byte{})
	URL = "servers/" + SERVERID + "/actions/power-off"
}

// Each test needs to have a name like `TestXXX`
// They also need a parameter of `*testing.T`
func TestPowerOffServerSuccess(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, nil), nil)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{SERVERID})
	if err != nil {
		test_framework.Errorf("Error detected: %s", err)
	}
}

func TestPowerOffServerNotFound(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, nil), nil)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{SERVERID})
	if err.Error() != "404" {
		test_framework.Errorf("Expected '404 NOT FOUND' error. Instead got %s", err)
	}
}

func TestPowerOffServerError(test_framework *testing.T) {
	bmcErr := ctlerrors.BMCError{
		Message:          "Something went wrong!",
		ValidationErrors: []string{},
	}

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(500, WithBody(bmcErr)), nil)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{SERVERID})
	if err.Error() != "500" {
		test_framework.Errorf("Expected '500 INTERNAL SERVER ERROR' error. Instead got %s", err)
	}
}

func TestPowerOffServerTooManyArgs(test_framework *testing.T) {
	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{SERVERID, "extra"})
	if err.Error() != "args" {
		test_framework.Errorf("Expected 'too many args' error. Instead got %s", err)
	}
}

func TestPowerOffServerClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, nil), errors.New("misc"))

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{SERVERID})
	if err.Error() != "client-fail" {
		test_framework.Errorf("Expected 'client failure' error. Instead got %s", err)
	}
}
