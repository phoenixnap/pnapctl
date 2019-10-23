package tests

import (
	"testing"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"

	delete "phoenixnap.com/pnap-cli/pnapctl/commands/delete/server"
)

func deleteSetup() {
	URL = "servers/" + SERVERID
}

func TestDeleteServerSuccess(test_framework *testing.T) {
	deleteSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformDelete(URL).
		Return(WithResponse(200, nil), nil)

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
		Return(WithResponse(404, nil), nil)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	testutil.AssertEqual(test_framework, "Server with ID "+SERVERID+" not found", err.Error())
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
		Return(WithResponse(404, nil), testutil.TestError)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	testutil.AssertEqual(test_framework, "Command 'delete server' could not be performed. Please try again later.", err.Error())
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
