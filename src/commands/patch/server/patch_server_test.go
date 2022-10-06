package server

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"sigs.k8s.io/yaml"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestPatchServerSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	serverPatch := generators.GenerateServerPatchSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverPatch)

	Filename = FILENAME

	// What the server should return.
	server := generators.GenerateServerSdk()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(*serverPatch)).
		Return(&server, WithResponse(200, WithBody(server)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchServerSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	serverPatch := generators.GenerateServerPatchSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPatch)

	Filename = FILENAME

	// What the server should return.
	server := generators.GenerateServerSdk()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(*serverPatch)).
		Return(&server, WithResponse(200, WithBody(server)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchServerFileNotFoundFailure(test_framework *testing.T) {

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestPatchServerUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	// filecontents := make([]byte, 10)
	filecontents := []byte(`notproperty: desc`)

	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "patch server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchServerFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'patch server' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "patch server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchServerBackendErrorFailure(test_framework *testing.T) {
	// Setup
	serverPatch := generators.GenerateServerPatchSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPatch)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(*serverPatch)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchServerClientFailure(test_framework *testing.T) {
	// Setup
	serverPatch := generators.GenerateServerPatchSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPatch)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(*serverPatch)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "patch server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchServerKeycloakFailure(test_framework *testing.T) {
	// Setup
	serverPatch := generators.GenerateServerPatchSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPatch)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(*serverPatch)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
