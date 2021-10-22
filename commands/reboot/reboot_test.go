package reboot

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	reboot "phoenixnap.com/pnap-cli/commands/reboot/server"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestRebootServerSuccess(test_framework *testing.T) {
	// Mocking
	actionResult := generators.GenerateActionResult()
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(SERVERID).
		Return(actionResult, WithResponse(200, WithBody(actionResult)), nil)

	// Run command
	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestRebootServerClientFail(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(SERVERID).
		Return(bmcapi.ActionResult{}, nil, testutil.TestError)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "reboot server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestRebootServerKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(SERVERID).
		Return(bmcapi.ActionResult{}, nil, testutil.TestKeycloakError)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestRebootServerNotFoundFail(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(SERVERID).
		Return(bmcapi.ActionResult{}, WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestRebootServerErrorFail(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(SERVERID).
		Return(bmcapi.ActionResult{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}
