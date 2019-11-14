package tests

import (
	"io/ioutil"
	"strings"
	"testing"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"

	delete "phoenixnap.com/pnap-cli/pnapctl/commands/delete/server"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

func deleteSetup() {
	URL = "servers/" + SERVERID
}

func TestDeleteServerSuccess(test_framework *testing.T) {
	deleteSetup()

	returnBody := "{\"result\":\"OK\",\"serverId\":123}"

	// Mocking
	PrepareMockClient(test_framework).
		PerformDelete(URL).
		Return(WithResponse(200, ioutil.NopCloser(strings.NewReader(returnBody))), nil)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestDeleteServerNotFound(test_framework *testing.T) {
	deleteSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformDelete(URL).
		Return(WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	testutil.AssertEqual(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestDeleteServerError(test_framework *testing.T) {
	deleteSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformDelete(URL).
		Return(WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	testutil.AssertEqual(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestDeleteServerClientFailure(test_framework *testing.T) {
	deleteSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformDelete(URL).
		Return(nil, testutil.TestError)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete server", ctlerrors.IncorrectRequestStructure)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestDeleteServerKeycloakFailure(test_framework *testing.T) {
	deleteSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformDelete(URL).
		Return(nil, testutil.TestKeycloakError)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	testutil.AssertEqual(test_framework, testutil.TestKeycloakError, err)
}
