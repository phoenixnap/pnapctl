package ssh_key

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func TestUpdateSshKeySuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	sshKeyUpdate := generators.GenerateSshKeyUpdate()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// What the server should return.
	sshKey := generators.GenerateSshKey()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(sshKeyUpdate)).
		Return(sshKey, WithResponse(200, WithBody(sshKey)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestUpdateSshKeySuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	sshKeyUpdate := generators.GenerateSshKeyUpdate()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// What the server should return.
	sshKey := generators.GenerateSshKey()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(sshKeyUpdate)).
		Return(sshKey, WithResponse(200, WithBody(sshKey)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
		ReadFile(FILENAME).
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
		ReadFile(FILENAME).
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
		ReadFile(FILENAME).
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
	sshKeyUpdate := generators.GenerateSshKeyUpdate()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(sshKeyUpdate)).
		Return(bmcapisdk.SshKey{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
	sshKeyUpdate := generators.GenerateSshKeyUpdate()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(sshKeyUpdate)).
		Return(bmcapisdk.SshKey{}, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
	sshKeyUpdate := generators.GenerateSshKeyUpdate()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(sshKeyUpdate)).
		Return(bmcapisdk.SshKey{}, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
