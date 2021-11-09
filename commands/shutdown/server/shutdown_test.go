package server

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestShutdownServerSuccess(test_framework *testing.T) {
	actionResult := generators.GenerateActionResult()
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(RESOURCEID).
		Return(actionResult, WithResponse(200, WithBody(actionResult)), nil)

	// Run command
	err := ShutdownCmd.RunE(ShutdownCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestShutdownServerNotFound(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(RESOURCEID).
		Return(bmcapisdk.ActionResult{}, WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := ShutdownCmd.RunE(ShutdownCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestShutdownServerError(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(RESOURCEID).
		Return(bmcapisdk.ActionResult{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := ShutdownCmd.RunE(ShutdownCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestShutdownServerClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(RESOURCEID).
		Return(bmcapisdk.ActionResult{}, nil, testutil.TestError)

	// Run command
	err := ShutdownCmd.RunE(ShutdownCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "shutdown server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestShutdownServerKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(RESOURCEID).
		Return(bmcapisdk.ActionResult{}, nil, testutil.TestKeycloakError)

	// Run command
	err := ShutdownCmd.RunE(ShutdownCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}