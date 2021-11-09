package server

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/tests/generators"

	"phoenixnap.com/pnap-cli/common/ctlerrors"

	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/tests/testutil"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"
)

func TestCreateServerSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	serverCreate := generators.GenerateServerCreate()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	Filename = FILENAME

	// What the server should return.
	createdServer := generators.GenerateServer()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(*serverCreate.ToSdk())).
		Return(createdServer, WithResponse(200, WithBody(createdServer)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	serverCreate := generators.GenerateServerCreate()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverCreate)

	Filename = FILENAME

	// What the server should return.
	createdServer := generators.GenerateServer()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(*serverCreate.ToSdk())).
		Return(createdServer, WithResponse(200, WithBody(createdServer)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
		ReadFile(FILENAME).
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
		ReadFile(FILENAME).
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
	serverCreate := generators.GenerateServerCreate()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(*serverCreate.ToSdk())).
		Return(bmcapisdk.Server{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
	serverCreate := generators.GenerateServerCreate()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(*serverCreate.ToSdk())).
		Return(bmcapisdk.Server{}, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
	serverCreate := generators.GenerateServerCreate()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(*serverCreate.ToSdk())).
		Return(bmcapisdk.Server{}, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
