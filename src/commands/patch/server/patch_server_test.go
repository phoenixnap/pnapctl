package server

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels"
	"phoenixnap.com/pnapctl/testsupport/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"gopkg.in/yaml.v2"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestPatchServerSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	serverPatch := generators.GenerateServerPatch()

	serverPatchModel := bmcapimodels.ServerPatch{
		Hostname:    serverPatch.Hostname,
		Description: serverPatch.Description,
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverPatchModel)

	Filename = FILENAME

	// What the server should return.
	server := generators.GenerateServer()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(serverPatch)).
		Return(server, WithResponse(200, WithBody(server)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchServerSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	serverPatch := generators.GenerateServerPatch()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPatch)

	Filename = FILENAME

	// What the server should return.
	server := generators.GenerateServer()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(serverPatch)).
		Return(server, WithResponse(200, WithBody(server)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
		ReadFile(FILENAME).
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
	filecontents := []byte(`Description: desc`)

	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
		ReadFile(FILENAME).
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
	serverPatch := generators.GenerateServerPatch()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPatch)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(serverPatch)).
		Return(bmcapisdk.Server{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
	serverPatch := generators.GenerateServerPatch()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPatch)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(serverPatch)).
		Return(bmcapisdk.Server{}, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
	serverPatch := generators.GenerateServerPatch()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPatch)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(serverPatch)).
		Return(bmcapisdk.Server{}, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
