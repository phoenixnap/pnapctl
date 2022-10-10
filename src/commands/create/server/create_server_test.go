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

func TestCreateServerSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	serverCreate := generators.GenerateServerCreateSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	Filename = FILENAME

	// What the server should return.
	createdServer := generators.GenerateServerSdk()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(serverCreate)).
		Return(&createdServer, WithResponse(200, WithBody(createdServer)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	serverCreate := generators.GenerateServerCreateSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverCreate)

	Filename = FILENAME

	// What the server should return.
	createdServer := generators.GenerateServerSdk()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(serverCreate)).
		Return(&createdServer, WithResponse(200, WithBody(createdServer)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerFileNotFoundFailure(test_framework *testing.T) {

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME) // TODO remove this from tests. We should give plain text here, not compare it.

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestCreateServerUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`sshKeys ["1","2","3","4"]`)

	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'create server' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "create server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerBackendErrorFailure(test_framework *testing.T) {
	// Setup
	serverCreate := generators.GenerateServerCreateSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(serverCreate)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerClientFailure(test_framework *testing.T) {

	// Setup
	serverCreate := generators.GenerateServerCreateSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(serverCreate)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerKeycloakFailure(test_framework *testing.T) {
	// Setup
	serverCreate := generators.GenerateServerCreateSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(serverCreate)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
