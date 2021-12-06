package sshkey

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/sshkeymodels"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeleteSshKeySuccess(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyDelete(RESOURCEID).
		Return(sshkeymodels.GenerateSshKeyDeleteResult(), WithResponse(200, nil), nil)

	// Run command
	err := DeleteSshKeyCmd.RunE(DeleteSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteSshKeyNotFound(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyDelete(RESOURCEID).
		Return(bmcapisdk.DeleteSshKeyResult{}, WithResponse(404, nil), nil)

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
		Return(bmcapisdk.DeleteSshKeyResult{}, WithResponse(500, nil), nil)

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
		Return(bmcapisdk.DeleteSshKeyResult{}, nil, testutil.TestError)

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
		Return(bmcapisdk.DeleteSshKeyResult{}, nil, testutil.TestKeycloakError)

	// Run command
	err := DeleteSshKeyCmd.RunE(DeleteSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
