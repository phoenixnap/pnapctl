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

func TestUpdateSshKeySuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	sshKeyUpdate := generators.Generate[bmcapi.SshKeyUpdate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// What the server should return.
	sshKey := generators.Generate[bmcapi.SshKey]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(sshKeyUpdate)).
		Return(&sshKey, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestUpdateSshKeySuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	sshKeyUpdate := generators.Generate[bmcapi.SshKeyUpdate]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// What the server should return.
	sshKey := generators.Generate[bmcapi.SshKey]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(sshKeyUpdate)).
		Return(&sshKey, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

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
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."})

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
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestUpdateSshKeyFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestUpdateSshKeyClientFailure(test_framework *testing.T) {
	// Setup
	sshKeyUpdate := generators.Generate[bmcapi.SshKeyUpdate]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(sshKeyUpdate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(sshKeyUpdate)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}
