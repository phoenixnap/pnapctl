package sshkey

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeleteSshKeySuccess(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyDelete(RESOURCEID).
		Return(testutil.AsPointer(generators.Generate[bmcapi.DeleteSshKeyResult]()), WithResponse(200, nil), nil)

	// Run command
	err := DeleteSshKeyCmd.RunE(DeleteSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteSshKeyNotFound(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyDelete(RESOURCEID).
		Return(nil, WithResponse(404, nil), nil)

	// Run command
	err := DeleteSshKeyCmd.RunE(DeleteSshKeyCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'delete ssh-key' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestDeleteSshKeyError(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyDelete(RESOURCEID).
		Return(nil, WithResponse(500, nil), nil)

	// Run command
	err := DeleteSshKeyCmd.RunE(DeleteSshKeyCmd, []string{RESOURCEID})

	expectedMessage := "Command 'delete ssh-key' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteSshKeyClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyDelete(RESOURCEID).
		Return(nil, nil, testutil.TestError)

	// Run command
	err := DeleteSshKeyCmd.RunE(DeleteSshKeyCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete ssh-key", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteSshKeyKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyDelete(RESOURCEID).
		Return(nil, nil, testutil.TestKeycloakError)

	// Run command
	err := DeleteSshKeyCmd.RunE(DeleteSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
