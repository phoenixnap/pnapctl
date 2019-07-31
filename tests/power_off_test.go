package tests

import (
	"bytes"
	"net/http"
	"os"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl"
	poweroff "phoenixnap.com/pnap-cli/pnapctl/bmc/power_off"

	"phoenixnap.com/pnap-cli/pnapctl/client"
	mocks "phoenixnap.com/pnap-cli/pnapctl/mocks"

	// "phoenixnap.com/pnap-cli/pnapctl/mocks"
	"github.com/golang/mock/gomock"
)

func TestPowerOffServerSuccess(t *testing.T) {
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
		PerformPost("servers/"+serverID+"/actions/power-off", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	os.Args = []string{"pnapctl", "bmc", "power-off", serverID}
	pnapctl.Execute()
}

func TestPowerOffServerConflict(t *testing.T) {
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
		PerformPost("servers/"+serverID+"/actions/power-off", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	os.Args = []string{"pnapctl", "bmc", "power-off", serverID}
	pnapctl.Execute()

	if poweroff.ErrorCode != "409" {
		t.Errorf("Expected '409 CONFLICT' error. Instead got %s", poweroff.ErrorCode)
	}
}

func TestPowerOffServerNotFound(t *testing.T) {
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
		PerformPost("servers/"+serverID+"/actions/power-off", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	os.Args = []string{"pnapctl", "bmc", "power-off", serverID}
	pnapctl.Execute()

	if poweroff.ErrorCode != "404" {
		t.Errorf("Expected '404 NOT FOUND' error. Instead got %s", poweroff.ErrorCode)
	}
}

func TestPowerOffServerInternalServerError(t *testing.T) {
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
		PerformPost("servers/"+serverID+"/actions/power-off", bytes.NewBuffer([]byte{})).
		Return(&resp, nil)

	os.Args = []string{"pnapctl", "bmc", "power-off", serverID}
	pnapctl.Execute()

	if poweroff.ErrorCode != "500" {
		t.Errorf("Expected '500 INTERNAL SERVER ERROR' error. Instead got %s", poweroff.ErrorCode)
	}
}
