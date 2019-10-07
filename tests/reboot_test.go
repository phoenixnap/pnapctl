package tests

import (
	"bytes"
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/reboot"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func rebootSetup() {
	Body = bytes.NewBuffer([]byte{})
	URL = "servers/" + SERVERID + "/actions/reboot"
}

func TestRebootServerSuccess(test_framework *testing.T) {
	rebootSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, nil), nil)

	// Run command
	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestRebootServerClientFail(test_framework *testing.T) {
	rebootSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, nil), testutil.TestError)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError("reboot")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestRebootServerNotFoundFail(test_framework *testing.T) {
	rebootSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, nil), nil)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New("Server with ID " + SERVERID + " not found.")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestRebootServerErrorFail(test_framework *testing.T) {
	rebootSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}
