package tests

import (
	"bytes"
	"net/http"
	"os"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl"

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
	defer func() {
		if r := recover(); r != nil {
			if r != "409-conflict" {
				t.Errorf("Panicked with error '%s'. Expected 409-conflict.", r)
			}
		} else {
			t.Errorf("The code did not panic - it should have.")
		}
	}()

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
}
