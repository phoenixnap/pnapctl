package servers

import (
	"errors"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/tables"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetAllServersShortSuccess(test_framework *testing.T) {
	serverlist := generators.GenerateServers(5)

	var shortServers []interface{}

	for _, x := range serverlist {
		shortServers = append(shortServers, tables.ToShortServerTable(x))
	}

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortServers, "get servers").
		Return(nil)

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllServersLongSuccess(test_framework *testing.T) {
	serverlist := generators.GenerateServers(5)

	var longServers []interface{}

	for _, x := range serverlist {
		longServers = append(longServers, tables.ToLongServerTable(x))
	}

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(longServers, "get servers").
		Return(nil)

	// to display full output
	Full = true

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestFilteredServersLongSuccess(test_framework *testing.T) {
	serverlist := generators.GenerateServers(5)

	var longServers []interface{}

	for _, x := range serverlist {
		longServers = append(longServers, tables.ToLongServerTable(x))
	}

	// to display full output
	Full = true
	tags = []string{generators.RandSeq(10), generators.RandSeq(10)}

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(longServers, "get servers").
		Return(nil)

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllServersClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return([]bmcapisdk.Server{}, WithResponse(200, nil), testutil.TestError)

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get servers", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllServersKeycloakFailure(test_framework *testing.T) {
	server := []bmcapisdk.Server{generators.GenerateServer()}
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(server, nil, testutil.TestKeycloakError)

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllServersPrinterFailure(test_framework *testing.T) {
	serverlist := generators.GenerateServers(5)

	var shortServers []interface{}

	for _, x := range serverlist {
		shortServers = append(shortServers, tables.ToShortServerTable(x))
	}

	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortServers, "get servers").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	Full = false

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
