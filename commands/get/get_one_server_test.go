package get

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnap-cli/commands/get/servers"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
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
	assert.NoError(test_framework, err)
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
	assert.NoError(test_framework, err)
}

func TestGetServerClientFailure(test_framework *testing.T) {
	getOneServerSetup()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(nil, testutil.TestError)

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get servers", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetServerKeycloakFailure(test_framework *testing.T) {
	getOneServerSetup()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(nil, testutil.TestKeycloakError)

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
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
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
