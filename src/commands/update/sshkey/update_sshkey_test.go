package sshkey

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func updateSshKeySuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	sshKeyUpdate := generators.Generate[bmcapi.SshKeyUpdate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, sshKeyUpdate)

	// What the server should return.
	sshKey := generators.Generate[bmcapi.SshKey]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(sshKeyUpdate)).
		Return(&sshKey, nil)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestUpdateSshKeySuccessYAML(test_framework *testing.T) {
	updateSshKeySuccess(test_framework, yaml.Marshal)
}

func TestUpdateSshKeySuccessJSON(test_framework *testing.T) {
	updateSshKeySuccess(test_framework, json.Marshal)
}

func TestUpdateSshKeyFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestUpdateSshKeyUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestUpdateSshKeyClientFailure(test_framework *testing.T) {
	// Setup
	sshKeyUpdate := generators.Generate[bmcapi.SshKeyUpdate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, sshKeyUpdate)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyPut(RESOURCEID, gomock.Eq(sshKeyUpdate)).
		Return(nil, testutil.TestError)

	// Run command
	err := UpdateSshKeyCmd.RunE(UpdateSshKeyCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
