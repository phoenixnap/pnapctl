package sshkey

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func createSshKeySuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	sshKeyCreate := generators.Generate[bmcapi.SshKeyCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, sshKeyCreate)

	// What the server should return.
	sshKey := generators.Generate[bmcapi.SshKey]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPost(gomock.Eq(sshKeyCreate)).
		Return(&sshKey, nil)

	// Run command
	err := CreateSshKeyCmd.RunE(CreateSshKeyCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateSshKeySuccessYAML(test_framework *testing.T) {
	createSshKeySuccess(test_framework, yaml.Marshal)
}

func TestCreateSshKeySuccessJSON(test_framework *testing.T) {
	createSshKeySuccess(test_framework, json.Marshal)
}

func TestCreateSshKeyFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateSshKeyCmd.RunE(CreateSshKeyCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestCreateSshKeyUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateSshKeyCmd.RunE(CreateSshKeyCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateSshKeyClientFailure(test_framework *testing.T) {
	// Setup
	sshKeyCreate := generators.Generate[bmcapi.SshKeyCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, sshKeyCreate)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPost(gomock.Eq(sshKeyCreate)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateSshKeyCmd.RunE(CreateSshKeyCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
