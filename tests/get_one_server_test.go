package tests

import (
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/commands/get/servers"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func getOneServerSetup() {
	URL = "servers/" + SERVERID
}

func TestGetServerShortSuccess(test_framework *testing.T) {
	getOneServerSetup()

	server := generators.GenerateServer()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(server)), nil)

	shortServer := generators.ConvertLongToShortServer(server)
	PrepareMockPrinter(test_framework).
		PrintOutput(&shortServer, false).
		Return(nil)

	servers.Full = false
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestGetServerLongSuccess(test_framework *testing.T) {
	getOneServerSetup()

	server := generators.GenerateServer()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(&server, false).
		Return(nil)

	servers.Full = true
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestGetServerClientFailure(test_framework *testing.T) {
	getOneServerSetup()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(nil, testutil.TestError)

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(nil, "get servers")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestGetServerKeycloakFailure(test_framework *testing.T) {
	getOneServerSetup()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(nil, testutil.TestKeycloakError)

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Assertions
	testutil.AssertEqual(test_framework, testutil.TestKeycloakError, err)
}

func TestGetServerPrinterFailure(test_framework *testing.T) {
	getOneServerSetup()

	server := generators.GenerateServer()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(server)), nil)

	shortServer := generators.ConvertLongToShortServer(server)
	PrepareMockPrinter(test_framework).
		PrintOutput(&shortServer, false).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	servers.Full = false
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Assertions
	testutil.AssertErrorCode(test_framework, err, ctlerrors.UnmarshallingInPrinter)
}
