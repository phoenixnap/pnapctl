package tests

import (
	"bytes"
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/client"
	reboot "phoenixnap.com/pnap-cli/pnapctl/commands/reboot/server"
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
		Return(WithResponse(200, WithBody(client.ResponseBody{Result: "OK"})), nil)

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
		Return(nil, testutil.TestError)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "reboot server", ctlerrors.IncorrectRequestStructure)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestRebootServerKeycloakFailure(test_framework *testing.T) {
	rebootSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(nil, testutil.TestKeycloakError)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Assertions
	testutil.AssertEqual(test_framework, testutil.TestKeycloakError, err)
}

func TestRebootServerNotFoundFail(test_framework *testing.T) {
	rebootSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})

	// Assertions
	testutil.AssertEqual(test_framework, testutil.GenericBMCError.Message, err.Error())
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
