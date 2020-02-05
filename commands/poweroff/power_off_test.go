package poweroff

import (
	"bytes"
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/common/client"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"

	"github.com/stretchr/testify/assert"
	poweroff "phoenixnap.com/pnap-cli/commands/poweroff/server"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

func powerOffSetup() {
	Body = bytes.NewBuffer([]byte{})
	URL = "servers/" + SERVERID + "/actions/power-off"
}

// Each test needs to have a name like `TestXXX`
// They also need a parameter of `*testing.T`
func TestPowerOffServerSuccess(test_framework *testing.T) {
	powerOffSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, WithBody(client.ResponseBody{Result: "OK"})), nil)

	// Run command
	err := poweroff.PowerOffServerCmd.RunE(poweroff.PowerOffServerCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPowerOffServerNotFound(test_framework *testing.T) {
	powerOffSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := poweroff.PowerOffServerCmd.RunE(poweroff.PowerOffServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestPowerOffServerError(test_framework *testing.T) {
	powerOffSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := poweroff.PowerOffServerCmd.RunE(poweroff.PowerOffServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPowerOffServerClientFailure(test_framework *testing.T) {
	powerOffSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(nil, testutil.TestError)

	// Run command
	err := poweroff.PowerOffServerCmd.RunE(poweroff.PowerOffServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "power-off server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPowerOffServerKeycloakFailure(test_framework *testing.T) {
	powerOffSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(nil, testutil.TestKeycloakError)

	// Run command
	err := poweroff.PowerOffServerCmd.RunE(poweroff.PowerOffServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
