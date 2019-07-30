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

func TestPowerOffServer(t *testing.T) {
	t.Run("Power off server - Success", PowerOffServerSuccess())
	t.Run("Power off server - Conflict", PowerOffServerConflict())
}

func PowerOffServerSuccess() func(*testing.T) {
	return func(t *testing.T) {
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
}

func PowerOffServerConflict() func(*testing.T) {
	return func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Log("RECOVERED ", r)
			} else {
				t.Errorf("The code did not panic - it should have.")
			}
		}()

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
}
