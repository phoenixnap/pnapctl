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
		Return(testutil.AsPointer(generators.Generate[bmcapi.DeleteSshKeyResult]()), nil)

	// Run command
	err := DeleteSshKeyCmd.RunE(DeleteSshKeyCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteSshKeyClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeyDelete(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := DeleteSshKeyCmd.RunE(DeleteSshKeyCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
