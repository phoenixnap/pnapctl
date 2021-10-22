package get

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/commands/get/servers"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/printer"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetServerShortSuccess(test_framework *testing.T) {

	server := generators.GenerateServer()
	var shortServer interface{}
	shortServer = printer.ToShortServer(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(SERVERID).
		Return(server, WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortServer, false, "get servers").
		Return(nil)

	servers.Full = false
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetServerLongSuccess(test_framework *testing.T) {
	server := generators.GenerateServer()
	var longServer interface{}
	longServer = printer.ToFullServer(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(SERVERID).
		Return(server, WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(longServer, false, "get servers").
		Return(nil)

	servers.Full = true
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetServerClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerGetById(SERVERID).
		Return(bmcapi.Server{}, nil, testutil.TestError)

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get servers", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetServerKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerGetById(SERVERID).
		Return(bmcapi.Server{}, nil, testutil.TestKeycloakError)

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetServerPrinterFailure(test_framework *testing.T) {
	server := generators.GenerateServer()
	shortServer := printer.ToShortServer(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(SERVERID).
		Return(server, WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortServer, false, "get servers").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	servers.Full = false
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{SERVERID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
