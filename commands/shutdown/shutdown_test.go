package shutdown

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	shutdown "phoenixnap.com/pnap-cli/commands/shutdown/server"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestShutdownServerSuccess(test_framework *testing.T) {
	actionResult := generators.GenerateActionResult()
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(SERVERID).
		Return(actionResult, WithResponse(200, WithBody(actionResult)), nil)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestShutdownServerNotFound(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(SERVERID).
		Return(bmcapi.ActionResult{}, WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestShutdownServerError(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(SERVERID).
		Return(bmcapi.ActionResult{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestShutdownServerClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(SERVERID).
		Return(bmcapi.ActionResult{}, nil, testutil.TestError)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "shutdown server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestShutdownServerKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(SERVERID).
		Return(bmcapi.ActionResult{}, nil, testutil.TestKeycloakError)

	// Run command
	err := shutdown.ShutdownCmd.RunE(shutdown.ShutdownCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
