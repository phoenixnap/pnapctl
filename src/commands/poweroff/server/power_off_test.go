package server

import (
	"errors"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

// Each test needs to have a name like `TestXXX`
// They also need a parameter of `*testing.T`
func TestPowerOffServerSuccess(test_framework *testing.T) {
	actionResult := generators.GenerateActionResult()
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOff(RESOURCEID).
		Return(actionResult, WithResponse(200, WithBody(actionResult)), nil)

	// Run command
	err := PowerOffServerCmd.RunE(PowerOffServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPowerOffServerNotFound(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOff(RESOURCEID).
		Return(bmcapisdk.ActionResult{}, WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := PowerOffServerCmd.RunE(PowerOffServerCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestPowerOffServerError(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOff(RESOURCEID).
		Return(bmcapisdk.ActionResult{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := PowerOffServerCmd.RunE(PowerOffServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPowerOffServerClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOff(RESOURCEID).
		Return(bmcapisdk.ActionResult{}, nil, testutil.TestError)

	// Run command
	err := PowerOffServerCmd.RunE(PowerOffServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "power-off server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPowerOffServerKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOff(RESOURCEID).
		Return(bmcapisdk.ActionResult{}, nil, testutil.TestKeycloakError)

	// Run command
	err := PowerOffServerCmd.RunE(PowerOffServerCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
