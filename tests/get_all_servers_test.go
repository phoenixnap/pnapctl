package tests

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/commands/get/servers"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func getAllServersSetup() {
	URL = "servers"
}

func TestGetAllServersUnmarshallingError(test_framework *testing.T) {
	getAllServersSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, ioutil.NopCloser(bytes.NewBuffer([]byte{0, 5}))), nil)

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	expectedErr := ctlerrors.GenericNonRequestError(ctlerrors.UnmarshallingErrorBody, "get servers")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestGetAllServersShortSuccess(test_framework *testing.T) {
	getAllServersSetup()

	serverlist := generators.GenerateServers(3)

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(serverlist)), nil)

	shortServer := generators.ConvertLongToShortServers(serverlist)
	PrepareMockPrinter(test_framework).
		PrintOutput(&shortServer, false).
		Return(nil)

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestGetAllServersLongSuccess(test_framework *testing.T) {
	getAllServersSetup()

	serverlist := generators.GenerateServers(3)

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(&serverlist, false).
		Return(nil)

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestGetAllServersClientFailure(test_framework *testing.T) {
	getAllServersSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(nil, testutil.TestError)

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get servers", ctlerrors.ErrorSendingRequest)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestGetAllServersKeycloakFailure(test_framework *testing.T) {
	getAllServersSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(nil, testutil.TestKeycloakError)

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	testutil.AssertEqual(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllServersPrinterFailure(test_framework *testing.T) {
	getAllServersSetup()

	// generate servers
	serverlist := generators.GenerateServers(3)

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(&serverlist, false).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	testutil.AssertErrorCode(test_framework, err, ctlerrors.UnmarshallingInPrinter)
}
