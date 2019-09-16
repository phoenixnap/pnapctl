package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"

	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/tests/testutil"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/reset"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
)

func TestResetSetup(t *testing.T) {
	URL = "servers/" + SERVERID + "/actions/reset"
}

func TestResetServerSuccessYAML(test_framework *testing.T) {
	// Setup
	serverReset := reset.ServerReset{
		SSHKeys: []string{"CNDI0W92UYC480D", "HDSIODIPS9879D"},
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverReset)

	// What the server should receive.
	jsonmarshal, _ := json.Marshal(serverReset)

	reset.Filename = FILENAME

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, bytes.NewBuffer(jsonmarshal)).
		Return(WithResponse(200, nil), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := reset.ResetCmd.RunE(reset.ResetCmd, []string{SERVERID})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestResetServerSuccessJSON(test_framework *testing.T) {
	// Setup
	serverReset := reset.ServerReset{
		SSHKeys: []string{"CNDI0W92UYC480D", "HDSIODIPS9879D"},
	}

	// What will be sent to the server, and the assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverReset)

	reset.Filename = FILENAME

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, bytes.NewBuffer(jsonmarshal)).
		Return(WithResponse(200, nil), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := reset.ResetCmd.RunE(reset.ResetCmd, []string{SERVERID})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestResetServerFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	reset.Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, errors.New("FileDoesNotExist")).
		Times(1)

	// Run command
	err := reset.ResetCmd.RunE(reset.ResetCmd, []string{SERVERID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())

}

func TestResetServerUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	// filecontents := make([]byte, 10)
	filecontents := []byte(`sshKeys ["1","2","3","4"]`)

	reset.Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := reset.ResetCmd.RunE(reset.ResetCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericNonRequestError(ctlerrors.UnmarshallingInFileProcessor, "reset")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestResetServerNotFoundFailure(test_framework *testing.T) {
	// Setup
	serverReset := reset.ServerReset{
		SSHKeys: []string{"CNDI0W92UYC480D", "HDSIODIPS9879D"},
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverReset)

	// What the server should receive.
	jsonmarshal, _ := json.Marshal(serverReset)

	reset.Filename = FILENAME

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, bytes.NewBuffer(jsonmarshal)).
		Return(WithResponse(404, nil), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := reset.ResetCmd.RunE(reset.ResetCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New("Server with ID " + SERVERID + " not found")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestResetServerFileReadingFailure(test_framework *testing.T) {
	// Setup
	reset.Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, errors.New("FileReading")).
		Times(1)

	// Run command
	err := reset.ResetCmd.RunE(reset.ResetCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericNonRequestError("FileReading", "reset")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestResetServerBackendErrorFailure(test_framework *testing.T) {
	// Setup
	serverReset := reset.ServerReset{
		SSHKeys: []string{"CNDI0W92UYC480D", "HDSIODIPS9879D"},
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverReset)

	// What the server should receive.
	jsonmarshal, _ := json.Marshal(serverReset)

	reset.Filename = FILENAME

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, bytes.NewBuffer(jsonmarshal)).
		Return(WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := reset.ResetCmd.RunE(reset.ResetCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestResetServerClientFailure(test_framework *testing.T) {
	// Setup
	serverReset := reset.ServerReset{
		SSHKeys: []string{"CNDI0W92UYC480D", "HDSIODIPS9879D"},
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverReset)

	// What the server should receive.
	jsonmarshal, _ := json.Marshal(serverReset)

	reset.Filename = FILENAME

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, bytes.NewBuffer(jsonmarshal)).
		Return(WithResponse(500, nil), testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := reset.ResetCmd.RunE(reset.ResetCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError("reset")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}
