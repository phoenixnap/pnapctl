package server

import (
	"encoding/json"
	"errors"
	"testing"

	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"

	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/tests/generators"
	"phoenixnap.com/pnap-cli/tests/testutil"

	"github.com/stretchr/testify/assert"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
)

func TestResetServerSuccessYAML(test_framework *testing.T) {
	// Setup
	serverReset := generators.GenerateServerReset()
	resetResult := generators.GenerateResetResult()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverReset)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(SERVERID, serverReset).
		Return(resetResult, WithResponse(200, WithBody(resetResult)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestResetServerSuccessJSON(test_framework *testing.T) {
	// Setup
	serverReset := generators.GenerateServerReset()
	resetResult := generators.GenerateResetResult()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverReset)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(SERVERID, serverReset).
		Return(resetResult, WithResponse(200, WithBody(resetResult)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestResetServerFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{SERVERID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	// filecontents := make([]byte, 10)
	filecontents := []byte(`sshKeys ["1","2","3","4"]`)

	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "reset server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerNotFoundFailure(test_framework *testing.T) {
	// Setup
	serverReset := generators.GenerateServerReset()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverReset)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(SERVERID, serverReset).
		Return(bmcapisdk.ResetResult{}, WithResponse(404, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestResetServerFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'reset server' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "reset server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerBackendErrorFailure(test_framework *testing.T) {
	// Setup
	serverReset := generators.GenerateServerReset()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverReset)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(SERVERID, serverReset).
		Return(bmcapisdk.ResetResult{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerClientFailure(test_framework *testing.T) {
	// Setup
	serverReset := generators.GenerateServerReset()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverReset)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(SERVERID, serverReset).
		Return(bmcapisdk.ResetResult{}, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "reset server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerKeycloakFailure(test_framework *testing.T) {
	// Setup
	serverReset := generators.GenerateServerReset()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverReset)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(SERVERID, serverReset).
		Return(bmcapisdk.ResetResult{}, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
