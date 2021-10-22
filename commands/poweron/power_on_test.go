package poweron

import (
	"errors"
	"testing"

	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"

	"github.com/stretchr/testify/assert"
	poweron "phoenixnap.com/pnap-cli/commands/poweron/server"
)

func TestPowerOnServerSuccess(test_framework *testing.T) {
	actionResult := generators.GenerateActionResult()

	PrepareBmcApiMockClient(test_framework).
		ServerPowerOn(SERVERID).
		Return(actionResult, WithResponse(200, WithBody(actionResult)), nil)

	err := poweron.PowerOnServerCmd.RunE(poweron.PowerOnServerCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPowerOnServerNotFound(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOn(SERVERID).
		Return(bmcapi.ActionResult{}, WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	err := poweron.PowerOnServerCmd.RunE(poweron.PowerOnServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestPowerOnServerError(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOn(SERVERID).
		Return(bmcapi.ActionResult{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	err := poweron.PowerOnServerCmd.RunE(poweron.PowerOnServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPowerOnServerKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOn(SERVERID).
		Return(bmcapi.ActionResult{}, nil, testutil.TestKeycloakError)

	// Run command
	err := poweron.PowerOnServerCmd.RunE(poweron.PowerOnServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
