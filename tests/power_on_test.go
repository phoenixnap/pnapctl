package tests

import (
	"bytes"
	"net/http"
	"os"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl"

	poweron "phoenixnap.com/pnap-cli/pnapctl/bmc/power_on"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	mocks "phoenixnap.com/pnap-cli/pnapctl/mocks"

	// "phoenixnap.com/pnap-cli/pnapctl/mocks"
	"github.com/golang/mock/gomock"
)

func TestPowerOnServerSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	serverID := "fake_id"

	// Init mock client
	m := mocks.NewMockWebClient(ctrl)

	resp := http.Response{
		StatusCode: 200,
	}

	client.MainClient = m

	m.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-on", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	os.Args = []string{"pnapctl", "bmc", "power-on", serverID}
	pnapctl.Execute()
}

func TestPowerOnServerConflict(t *testing.T) {
	ctrl := gomock.NewController(t)
	serverID := "fake_id"

	// Init mock client
	m := mocks.NewMockWebClient(ctrl)

	resp := http.Response{
		StatusCode: 409,
	}

	client.MainClient = m

	m.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-on", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	os.Args = []string{"pnapctl", "bmc", "power-on", serverID}
	pnapctl.Execute()

	if poweron.ErrorCode != "409" {
		t.Errorf("Expected '409 CONFLICT' error. Instead got %s", poweron.ErrorCode)
	}
}

func TestPowerOnServerNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	serverID := "fake_id"

	// Init mock client
	m := mocks.NewMockWebClient(ctrl)

	resp := http.Response{
		StatusCode: 404,
	}

	client.MainClient = m

	m.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-on", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	os.Args = []string{"pnapctl", "bmc", "power-on", serverID}
	pnapctl.Execute()

	if poweron.ErrorCode != "404" {
		t.Errorf("Expected '404 NOT FOUND' error. Instead got %s", poweron.ErrorCode)
	}
}

func TestPowerOnServerInternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	serverID := "fake_id"

	// Init mock client
	m := mocks.NewMockWebClient(ctrl)

	resp := http.Response{
		StatusCode: 500,
	}

	client.MainClient = m

	m.
		EXPECT().
		PerformPost("servers/"+serverID+"/actions/power-on", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	os.Args = []string{"pnapctl", "bmc", "power-on", serverID}
	pnapctl.Execute()

	if poweron.ErrorCode != "500" {
		t.Errorf("Expected '500 INTERNAL SERVER ERROR' error. Instead got %s", poweron.ErrorCode)
	}
}
