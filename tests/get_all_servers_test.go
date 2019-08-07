package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"phoenixnap.com/pnap-cli/pnapctl"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	mocks "phoenixnap.com/pnap-cli/pnapctl/mocks"
	"phoenixnap.com/pnap-cli/pnapctl/printer"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/get/servers"
	"phoenixnap.com/pnap-cli/tests/generators"
)

func TestGetAllServersSuccess(test_framework *testing.T) {
	controller := gomock.NewController(test_framework)

	mockClient := mocks.NewMockWebClient(controller)
	mockPrinter := mocks.NewMockPrinter(controller)

	serverlist := []servers.LongServer{
		generators.GenerateServer(),
		generators.GenerateServer(),
		generators.GenerateServer(),
	}

	data, _ := json.Marshal(serverlist)

	resp := http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer(data)),
	}

	client.MainClient = mockClient
	printer.MainPrinter = mockPrinter

	mockClient.
		EXPECT().
		PerformGet("servers").
		Return(&resp, nil)

	mockPrinter.
		EXPECT().
		PrintOutput(data, &[]servers.ShortServer{}).
		Return(3, nil)

	os.Args = []string{"pnapctl", "bmc", "get", "servers"}
	pnapctl.Execute()
}
