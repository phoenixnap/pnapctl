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

func shutdownSetup() {
	Body = bytes.NewBuffer([]byte{})
	URL = "servers/" + SERVERID + "/actions/shutdown"
}

func TestShutdownServerSuccess(test_framework *testing.T) {
	shutdownSetup()

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
	shutdownSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, nil), nil)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New("Server with ID " + SERVERID + " not found.")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestShutdownServerError(test_framework *testing.T) {
	shutdownSetup()

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

func TestShutdownServerClientFailure(test_framework *testing.T) {
	shutdownSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, nil), testutil.TestError)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(nil, "shutdown")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestShutdownServerKeycloakFailure(test_framework *testing.T) {
	shutdownSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(nil, testutil.TestKeycloakError)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Assertions
	testutil.AssertEqual(test_framework, testutil.TestKeycloakError, err)
}
