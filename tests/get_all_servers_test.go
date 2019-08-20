package tests

import (
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/get/servers"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
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

	if err != nil {
		test_framework.Error("Expected no error, found:", err)
	}
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

	if err != nil {
		test_framework.Error("Expected no error, found:", err)
	}
}

func TestGetAllServersClientFailure(test_framework *testing.T) {
	// generate 3 long servers
	serverlist := generators.GenerateServers(3)

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(serverlist)), errors.New("client-fail"))

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	if err.Error() != "get-fail" {
		test_framework.Error("Expected client failure error, found:", err)
	}
}

func TestGetAllServersPrinterFailure(test_framework *testing.T) {
	serverlist := generators.GenerateServers(3)

	// Mocking
	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(serverlist)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(WithData(serverlist), &[]servers.LongServer{}).
		Return(3, errors.New("oops"))

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	if err.Error() != "oops" {
		test_framework.Error("Expected printer failure error, found:", err)
	}
}
