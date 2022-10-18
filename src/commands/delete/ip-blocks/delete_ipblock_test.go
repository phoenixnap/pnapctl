package ip_blocks

import (
	"testing"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestDeleteIpBlockSuccess(test_framework *testing.T) {
	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdDelete(RESOURCEID).
		Return(generators.GenerateDeleteIpBlockResultSdk(), WithResponse(200, nil), nil)

	// Run command
	err := DeleteIpBlockCmd.RunE(DeleteIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteIpBlockNotFound(test_framework *testing.T) {
	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdDelete(RESOURCEID).
		Return(nil, WithResponse(404, nil), nil)

	// Run command
	err := DeleteIpBlockCmd.RunE(DeleteIpBlockCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'delete ip-block' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestDeleteIpBlockError(test_framework *testing.T) {
	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdDelete(RESOURCEID).
		Return(nil, WithResponse(500, nil), nil)

	// Run command
	err := DeleteIpBlockCmd.RunE(DeleteIpBlockCmd, []string{RESOURCEID})

	expectedMessage := "Command 'delete ip-block' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteIpBlockClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdDelete(RESOURCEID).
		Return(nil, nil, testutil.TestError)

	// Run command
	err := DeleteIpBlockCmd.RunE(DeleteIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete ip-block", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteIpBlockKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdDelete(RESOURCEID).
		Return(nil, nil, testutil.TestKeycloakError)

	// Run command
	err := DeleteIpBlockCmd.RunE(DeleteIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
