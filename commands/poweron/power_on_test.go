package poweron

import (
	"bytes"
	"errors"
	"testing"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"

	"github.com/stretchr/testify/assert"
	poweron "phoenixnap.com/pnap-cli/commands/poweron/server"
	"phoenixnap.com/pnap-cli/common/client"
)

func powerOnSetup() {
	Body = bytes.NewBuffer([]byte{})
	URL = "servers/" + SERVERID + "/actions/power-on"
}

func TestPowerOnServerSuccess(test_framework *testing.T) {
	powerOnSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, WithBody(client.ResponseBody{Result: "OK"})), nil)

	err := poweron.PowerOnServerCmd.RunE(poweron.PowerOnServerCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPowerOnServerNotFound(test_framework *testing.T) {
	powerOnSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	err := poweron.PowerOnServerCmd.RunE(poweron.PowerOnServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestPowerOnServerError(test_framework *testing.T) {
	powerOnSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	err := poweron.PowerOnServerCmd.RunE(poweron.PowerOnServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPowerOnServerKeycloakFailure(test_framework *testing.T) {
	powerOnSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(nil, testutil.TestKeycloakError)

	// Run command
	err := poweron.PowerOnServerCmd.RunE(poweron.PowerOnServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
