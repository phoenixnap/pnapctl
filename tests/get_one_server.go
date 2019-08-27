package tests

import (
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/get/servers"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
)

func TestGetServerSetup(test_framework *testing.T) {
	URL = "servers/" + SERVERID
}

func TestGetServerShortSuccess(test_framework *testing.T) {
	server := generators.GenerateServer()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(WithData(server), &servers.ShortServer{}).
		Return(1, nil)

	servers.ID = SERVERID
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	if err != nil {
		test_framework.Error("Expected no error, found:", err)
	}
}

func TestGetServerLongSuccess(test_framework *testing.T) {
	server := generators.GenerateServer()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(WithData(server), &servers.ShortServer{}).
		Return(1, nil)

	servers.ID = SERVERID
	servers.Full = true
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	if err != nil {
		test_framework.Error("Expected no error, found:", err)
	}
}

func TestGetServerClientFailure(test_framework *testing.T) {
	server := generators.GenerateServer()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(server)), errors.New("client-fail"))

	servers.ID = SERVERID
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	if err.Error() != "get-fail" {
		test_framework.Error("Expected client failure error, found:", err)
	}
}

func TestGetServerPrinterFailure(test_framework *testing.T) {
	server := generators.GenerateServer()

	PrepareMockClient(test_framework).
		PerformGet(URL).
		Return(WithResponse(200, WithBody(server)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(WithData(server), &servers.ShortServer{}).
		Return(1, errors.New("printer-not-connected?"))

	servers.ID = SERVERID
	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	if err.Error() != "printer-not-connected?" {
		test_framework.Error("Expected no error, found:", err)
	}
}
