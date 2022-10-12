package ip_blocks

import (
	"testing"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"github.com/stretchr/testify/assert"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestDeleteIpBlockSuccess(test_framework *testing.T) {
	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdDelete(RESOURCEID).
		Return(testutil.AsPointer(generators.Generate[ipapi.DeleteIpBlockResult]()), nil)

	// Run command
	err := DeleteIpBlockCmd.RunE(DeleteIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteIpBlockClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdDelete(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := DeleteIpBlockCmd.RunE(DeleteIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteIpBlockKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdDelete(RESOURCEID).
		Return(nil, testutil.TestKeycloakError)

	// Run command
	err := DeleteIpBlockCmd.RunE(DeleteIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
