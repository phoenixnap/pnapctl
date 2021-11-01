package servers

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/tables"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetServerShortSuccess(test_framework *testing.T) {

	server := generators.GenerateServer()
	var shortServer interface{}
	shortServer = tables.ToShortServerTable(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(SERVERID).
		Return(server, WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortServer, "get servers").
		Return(nil)

	Full = false
	err := GetServersCmd.RunE(GetServersCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetServerLongSuccess(test_framework *testing.T) {
	server := generators.GenerateServer()
	var longServer interface{}
	longServer = tables.ToLongServerTable(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(SERVERID).
		Return(server, WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(longServer, "get servers").
		Return(nil)

	Full = true
	err := GetServersCmd.RunE(GetServersCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetServerNotFound(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerGetById(SERVERID).
		Return(bmcapi.Server{}, WithResponse(400, nil), nil)

	err := GetServersCmd.RunE(GetServersCmd, []string{SERVERID})

	// Assertions
	expectedMessage := "Command 'get servers' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetServerClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerGetById(SERVERID).
		Return(bmcapi.Server{}, nil, testutil.TestError)

	err := GetServersCmd.RunE(GetServersCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get servers", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetServerKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerGetById(SERVERID).
		Return(bmcapi.Server{}, nil, testutil.TestKeycloakError)

	err := GetServersCmd.RunE(GetServersCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetServerPrinterFailure(test_framework *testing.T) {
	server := generators.GenerateServer()
	shortServer := tables.ToShortServerTable(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(SERVERID).
		Return(server, WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortServer, "get servers").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	Full = false
	err := GetServersCmd.RunE(GetServersCmd, []string{SERVERID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
