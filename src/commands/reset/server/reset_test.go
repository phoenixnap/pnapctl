package server

import (
	"encoding/json"
	"errors"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"

	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestResetServerSuccessYAML(test_framework *testing.T) {
	// Setup
	serverReset := servermodels.GenerateServerResetSdk()
	resetResult := servermodels.GenerateResetResultSdk()

	// to be used for marshaling only
	serverResetModel := servermodels.ServerReset{
		InstallDefaultSshKeys: nil,
		SshKeys:               nil,
		SshKeyIds:             nil,
		OsConfiguration:       nil,
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverResetModel)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(RESOURCEID, serverReset).
		Return(resetResult, WithResponse(200, WithBody(resetResult)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestResetServerSuccessJSON(test_framework *testing.T) {
	// Setup
	serverReset := servermodels.GenerateServerResetSdk()
	resetResult := servermodels.GenerateResetResultSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverReset)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(RESOURCEID, serverReset).
		Return(resetResult, WithResponse(200, WithBody(resetResult)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestResetServerSuccessNoFile(test_framework *testing.T) {
	// Setup
	resetResult := servermodels.GenerateResetResultSdk()

	Filename = ""

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(RESOURCEID, bmcapisdk.ServerReset{}).
		Return(resetResult, WithResponse(200, WithBody(resetResult)), nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

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
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

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
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "reset server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerNotFoundFailure(test_framework *testing.T) {
	// Setup
	serverReset := servermodels.GenerateServerResetSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverReset)
	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(RESOURCEID, serverReset).
		Return(bmcapisdk.ResetResult{}, WithResponse(404, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

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
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "reset server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerBackendErrorFailure(test_framework *testing.T) {
	// Setup
	serverReset := servermodels.GenerateServerResetSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverReset)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(RESOURCEID, serverReset).
		Return(bmcapisdk.ResetResult{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerClientFailure(test_framework *testing.T) {
	// Setup
	serverReset := servermodels.GenerateServerResetSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverReset)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(RESOURCEID, serverReset).
		Return(bmcapisdk.ResetResult{}, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "reset server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestResetServerKeycloakFailure(test_framework *testing.T) {
	// Setup
	serverReset := servermodels.GenerateServerResetSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverReset)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(RESOURCEID, serverReset).
		Return(bmcapisdk.ResetResult{}, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
