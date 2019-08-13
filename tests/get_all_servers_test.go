package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	mocks "phoenixnap.com/pnap-cli/pnapctl/mocks"

	"phoenixnap.com/pnap-cli/pnapctl/printer"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/get/servers"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
)

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
	// init controller
	controller := gomock.NewController(test_framework)

	// mocks
	mockClient := mocks.NewMockWebClient(controller)
	mockPrinter := mocks.NewMockPrinter(controller)

	// generate 3 long servers
	serverlist := []servers.LongServer{
		generators.GenerateServer(),
		generators.GenerateServer(),
		generators.GenerateServer(),
	}

	// marshal array into bytes
	data, _ := json.Marshal(serverlist)

	// manufacture response
	resp := http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer(data)),
	}

	// swap Mains with Mocks
	client.MainClient = mockClient
	printer.MainPrinter = mockPrinter

	// mock client
	mockClient.
		EXPECT().
		PerformGet("servers"). // should hit GET /servers
		Return(&resp, nil)     // returns &resp and nil

	mockPrinter.
		EXPECT().
		PrintOutput(data, &[]servers.LongServer{}). // should print short server
		Return(-1, errors.New("test-error"))

	// to display full output
	servers.Full = true

	err := servers.GetServersCmd.RunE(servers.GetServersCmd, []string{})

	if err.Error() != "test-error" {
		test_framework.Error("Expected printer failure error, found:", err)
	}
}
