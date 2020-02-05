package shutdown

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	shutdown "phoenixnap.com/pnap-cli/commands/shutdown/server"
	"phoenixnap.com/pnap-cli/common/client"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
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
		Return(WithResponse(200, WithBody(client.ResponseBody{Result: "OK"})), nil)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestShutdownServerNotFound(test_framework *testing.T) {
	shutdownSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
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
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestShutdownServerClientFailure(test_framework *testing.T) {
	shutdownSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(nil, testutil.TestError)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "shutdown server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
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
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
