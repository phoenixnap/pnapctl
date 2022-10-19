package privatenetwork

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

const deleteResult = "The server is being removed from the specified private network."

func TestDeleteServerPrivateNetworkSuccess(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPrivateNetworkDelete(RESOURCEID, RESOURCEID).
		Return(deleteResult, nil)

	// Run command
	err := DeleteServerPrivateNetworkCmd.RunE(DeleteServerPrivateNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteServerPrivateNetworkClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPrivateNetworkDelete(RESOURCEID, RESOURCEID).
		Return("", testutil.TestError)

	// Run command
	err := DeleteServerPrivateNetworkCmd.RunE(DeleteServerPrivateNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}
