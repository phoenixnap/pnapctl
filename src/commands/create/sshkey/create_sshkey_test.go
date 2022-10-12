package sshkey

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreateSshKeySuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	sshKeyCreate := generators.Generate[bmcapi.SshKeyCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(sshKeyCreate)

	Filename = FILENAME

	// What the server should return.
	sshKey := generators.Generate[bmcapi.SshKey]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPost(gomock.Eq(sshKeyCreate)).
		Return(&sshKey, nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateSshKeyCmd.RunE(CreateSshKeyCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateSshKeySuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	sshKeyCreate := generators.Generate[bmcapi.SshKeyCreate]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(sshKeyCreate)

	Filename = FILENAME

	// What the server should return.
	sshKey := generators.Generate[bmcapi.SshKey]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPost(gomock.Eq(sshKeyCreate)).
		Return(&sshKey, nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateSshKeyCmd.RunE(CreateSshKeyCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateSshKeyFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreateSshKeyCmd.RunE(CreateSshKeyCmd, []string{})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestCreateSshKeyUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`name this is a bad name`)

	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := CreateSshKeyCmd.RunE(CreateSshKeyCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateSshKeyFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := CreateSshKeyCmd.RunE(CreateSshKeyCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateSshKeyClientFailure(test_framework *testing.T) {
	// Setup
	sshKeyCreate := generators.Generate[bmcapi.SshKeyCreate]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(sshKeyCreate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPost(gomock.Eq(sshKeyCreate)).
		Return(nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateSshKeyCmd.RunE(CreateSshKeyCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateSshKeyKeycloakFailure(test_framework *testing.T) {
	// Setup
	sshKeyCreate := generators.Generate[bmcapi.SshKeyCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(sshKeyCreate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPost(gomock.Eq(sshKeyCreate)).
		Return(nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateSshKeyCmd.RunE(CreateSshKeyCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
