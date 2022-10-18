package sshkey

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestUpdateSshKeySuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	sshKeyUpdate := generators.GenerateSshKeyUpdateSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// What the server should return.
	sshKey := generators.GenerateSshKeySdk()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(*sshKeyUpdate)).
		Return(&sshKey, WithResponse(200, WithBody(sshKey)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestUpdateSshKeySuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	sshKeyUpdate := generators.GenerateSshKeyUpdateSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// What the server should return.
	sshKey := generators.GenerateSshKeySdk()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(*sshKeyUpdate)).
		Return(&sshKey, WithResponse(200, WithBody(sshKey)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestUpdateSshKeyFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestUpdateSshKeyUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`name this is a bad name`)

	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "update ssh-key", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestUpdateSshKeyFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'update ssh-key' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "update ssh-key", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestUpdateSshKeyBackendErrorFailure(test_framework *testing.T) {
	// Setup
	sshKeyUpdate := generators.GenerateSshKeyUpdateSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(*sshKeyUpdate)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestUpdateSshKeyClientFailure(test_framework *testing.T) {
	// Setup
	sshKeyUpdate := generators.GenerateSshKeyUpdateSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(*sshKeyUpdate)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "update ssh-key", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestUpdateSshKeyKeycloakFailure(test_framework *testing.T) {
	// Setup
	sshKeyUpdate := generators.GenerateSshKeyUpdateSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(*sshKeyUpdate)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
