package servers

import (
	"errors"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/testsupport/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetServerShortSuccess(test_framework *testing.T) {

	server := generators.GenerateServer()
	var shortServer = tables.ToShortServerTable(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(RESOURCEID).
		Return(server, WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortServer, "get servers").
		Return(nil)

	Full = false
	err := GetServersCmd.RunE(GetServersCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetServerLongSuccess(test_framework *testing.T) {
	server := generators.GenerateServer()
	var longServer = tables.ToLongServerTable(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(RESOURCEID).
		Return(server, WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(longServer, "get servers").
		Return(nil)

	Full = true
	err := GetServersCmd.RunE(GetServersCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetServerNotFound(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerGetById(RESOURCEID).
		Return(bmcapisdk.Server{}, WithResponse(400, nil), nil)

	err := GetServersCmd.RunE(GetServersCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'get servers' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetServerClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerGetById(RESOURCEID).
		Return(bmcapisdk.Server{}, nil, testutil.TestError)

	err := GetServersCmd.RunE(GetServersCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get servers", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetServerKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerGetById(RESOURCEID).
		Return(bmcapisdk.Server{}, nil, testutil.TestKeycloakError)

	err := GetServersCmd.RunE(GetServersCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetServerPrinterFailure(test_framework *testing.T) {
	server := generators.GenerateServer()
	shortServer := tables.ToShortServerTable(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(RESOURCEID).
		Return(server, WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortServer, "get servers").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	Full = false
	err := GetServersCmd.RunE(GetServersCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
