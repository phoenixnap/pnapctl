package tests

import (
	"bytes"
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/reboot"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
)

func TestRebootSetup(t *testing.T) {
	Body = bytes.NewBuffer([]byte{})
	URL = "servers/" + SERVERID + "/actions/reboot"
}

func TestRebootServerSuccess(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, nil), nil)

	// Run command
	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})
	if err != nil {
		test_framework.Errorf("Error detected: %s", err)
	}
}

func TestRebootServerArgFail(test_framework *testing.T) {
	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID, "extra"})
	if err.Error() != "args" {
		test_framework.Errorf("Expected invalid args error - found %s", err)
	}
}

func TestRebootServerClientFail(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(200, nil), errors.New("oops"))

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})
	if err.Error() != "client-fail" {
		test_framework.Errorf("Error: Expected client failure error, found %s", err)
	}
}

func TestRebootServerNotFoundFail(test_framework *testing.T) {
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(404, nil), nil)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})
	if err.Error() != "404" {
		test_framework.Errorf("Error: not found error, found %s", err)
	}
}

func TestRebootServerInternalServerErrorFail(test_framework *testing.T) {
	bmcErr := ctlerrors.BMCError{
		Message:          "Something went wrong!",
		ValidationErrors: []string{},
	}
	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, Body).
		Return(WithResponse(500, WithBody(bmcErr)), nil)

	err := reboot.RebootCmd.RunE(reboot.RebootCmd, []string{SERVERID})
	if err.Error() != "500" {
		test_framework.Errorf("Error: Expected internal server error, found %s", err)
	}
}
