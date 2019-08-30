package tests

import (
	"bytes"
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/shutdown"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
)

func TestShutdownSetup(t *testing.T) {
	Body = bytes.NewBuffer([]byte{})
	URL = "servers/" + SERVERID + "/actions/shutdown"
}

func TestShutdownServerSuccess(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, nil), nil)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})
	if err != nil {
		test_framework.Errorf("Error detected: %s", err)
	}
}

func TestShutdownServerNotFound(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, nil), nil)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})
	if err.Error() != "404" {
		test_framework.Errorf("Expected '404 NOT FOUND' error. Instead got %s", err)
	}
}

func TestShutdownServerError(test_framework *testing.T) {
	bmcErr := ctlerrors.BMCError{
		Message:          "Something went wrong!",
		ValidationErrors: []string{},
	}
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(500, WithBody(bmcErr)), nil)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})
	if err.Error() != "500" {
		test_framework.Errorf("Expected '500 INTERNAL SERVER ERROR' error. Instead got %s", err)
	}
}

func TestShutdownServerTooManyArgs(test_framework *testing.T) {
	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID, "BONUS"})
	if err.Error() != "args" {
		test_framework.Errorf("Expected 'args' error. Instead got %s", err)
	}
}

func TestShutdownServerClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, nil), errors.New("oops"))

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})
	if err.Error() != "client-fail" {
		test_framework.Errorf("Expected 'client failure' error. Instead got %s", err)
	}
}
