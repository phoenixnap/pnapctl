package tests

import (
	"bytes"
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/shutdown"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
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

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestShutdownServerNotFound(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, nil), nil)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New("Error: Server with ID " + SERVERID + " not found.")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestShutdownServerError(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestShutdownServerTooManyArgs(test_framework *testing.T) {
	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID, "BONUS"})

	// Expected error
	expectedErr := ctlerrors.InvalidNumberOfArgs(1, 2, "shutdown")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestShutdownServerClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, nil), testutil.TestError)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.ShutdownServerGenericError(testutil.TestError)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}
