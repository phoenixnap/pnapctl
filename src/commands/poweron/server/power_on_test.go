package server

import (
	"errors"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
)

func TestPowerOnServerSuccess(test_framework *testing.T) {
	actionResult := servermodels.GenerateActionResultSdk()

	PrepareBmcApiMockClient(test_framework).
		ServerPowerOn(RESOURCEID).
		Return(actionResult, WithResponse(200, WithBody(actionResult)), nil)

	err := PowerOnServerCmd.RunE(PowerOnServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPowerOnServerNotFound(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOn(RESOURCEID).
		Return(bmcapisdk.ActionResult{}, WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	err := PowerOnServerCmd.RunE(PowerOnServerCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestPowerOnServerError(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOn(RESOURCEID).
		Return(bmcapisdk.ActionResult{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	err := PowerOnServerCmd.RunE(PowerOnServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPowerOnServerKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOn(RESOURCEID).
		Return(bmcapisdk.ActionResult{}, nil, testutil.TestKeycloakError)

	// Run command
	err := PowerOnServerCmd.RunE(PowerOnServerCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
