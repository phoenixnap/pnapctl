package reset

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/common/client"
	"phoenixnap.com/pnap-cli/common/ctlerrors"

	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/tests/testutil"

	"github.com/stretchr/testify/assert"
	reset "phoenixnap.com/pnap-cli/commands/reset/server"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
)

func resetSetup() {
	URL = "servers/" + SERVERID + "/actions/reset"
}

func TestResetServerSuccessYAML(test_framework *testing.T) {
	resetSetup()

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
		Return(WithResponse(200, WithBody(client.ResponseBody{Result: "OK"})), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := reset.ResetServerCmd.RunE(reset.ResetServerCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestResetServerSuccessJSON(test_framework *testing.T) {
	resetSetup()

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
		Return(WithResponse(200, WithBody(client.ResponseBody{Result: "OK"})), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := reset.ResetServerCmd.RunE(reset.ResetServerCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestResetServerFileNotFoundFailure(test_framework *testing.T) {
	resetSetup()

	// Setup
	reset.Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, errors.New(ctlerrors.FileDoesNotExist)).
		Times(1)

	// Run command
	err := reset.ResetServerCmd.RunE(reset.ResetServerCmd, []string{SERVERID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestResetServerUnmarshallingFailure(test_framework *testing.T) {
	resetSetup()

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
	err := reset.ResetServerCmd.RunE(reset.ResetServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericNonRequestError(ctlerrors.UnmarshallingInFileProcessor, "reset server")

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerNotFoundFailure(test_framework *testing.T) {
	resetSetup()

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
		Return(WithResponse(404, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := reset.ResetServerCmd.RunE(reset.ResetServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestResetServerFileReadingFailure(test_framework *testing.T) {
	resetSetup()

	// Setup
	reset.Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, errors.New(ctlerrors.FileReading)).
		Times(1)

	// Run command
	err := reset.ResetServerCmd.RunE(reset.ResetServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericNonRequestError(ctlerrors.FileReading, "reset server")

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerBackendErrorFailure(test_framework *testing.T) {
	resetSetup()

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
	err := reset.ResetServerCmd.RunE(reset.ResetServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerClientFailure(test_framework *testing.T) {
	resetSetup()

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
		Return(nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := reset.ResetServerCmd.RunE(reset.ResetServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "reset server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerKeycloakFailure(test_framework *testing.T) {
	resetSetup()

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
		Return(nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := reset.ResetServerCmd.RunE(reset.ResetServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
