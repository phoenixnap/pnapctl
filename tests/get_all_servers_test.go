package tests

import (
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/get/servers"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetAllSetup(test_framework *testing.T) {
	URL = "servers"
}

func TestGetAllServersShortSuccess(test_framework *testing.T) {
	serverlist := generators.GenerateServers(3)

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(WithData(serverlist), &[]servers.ShortServer{}).
		Return(3, nil)

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestGetAllServersLongSuccess(test_framework *testing.T) {
	serverlist := generators.GenerateServers(3)

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(WithData(serverlist), &[]servers.LongServer{}).
		Return(3, nil)

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestGetAllServersClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(nil, testutil.TestError)

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError("get servers")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestGetAllServersPrinterFailure(test_framework *testing.T) {
	// generate servers
	serverlist := generators.GenerateServers(3)

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(WithData(serverlist), &[]servers.LongServer{}).
		Return(-1, errors.New(ctlerrors.UnmarshallingInPrinter))

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	// Assertions
	testutil.AssertErrorCode(test_framework, err, ctlerrors.UnmarshallingInPrinter)
}
