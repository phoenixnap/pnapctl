package tests

import (
	"bytes"
	"errors"
	"testing"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/poweroff"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
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
		Return(WithResponse(200, nil), nil)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{SERVERID})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestPowerOffServerNotFound(test_framework *testing.T) {
	powerOffSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, nil), nil)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New("Server with ID " + SERVERID + " not found")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestPowerOffServerError(test_framework *testing.T) {
	powerOffSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestPowerOffServerClientFailure(test_framework *testing.T) {
	powerOffSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, nil), testutil.TestError)

	// Run command
	err := poweroff.P_OffCmd.RunE(poweroff.P_OffCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError("power-off")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}
