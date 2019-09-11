package tests

import (
	"bytes"
	"errors"
	"testing"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/poweron"
)

func TestPowerOnSetup(t *testing.T) {
	Body = bytes.NewBuffer([]byte{})
	URL = "servers/" + SERVERID + "/actions/power-on"
}

func TestPowerOnServerSuccess(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, nil), nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{SERVERID})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestPowerOnServerNotFound(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, nil), nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New("Server with ID " + SERVERID + " not found")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestPowerOnServerError(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	err := poweron.P_OnCmd.RunE(poweron.P_OnCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}
