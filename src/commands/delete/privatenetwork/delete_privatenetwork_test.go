package privatenetwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeletePrivateNetworkSuccess(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkDelete(RESOURCEID).
		Return(nil)

	// Run command
	err := DeletePrivateNetworkCmd.RunE(DeletePrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeletePrivateNetworkClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkDelete(RESOURCEID).
		Return(testutil.TestError)

	// Run command
	err := DeletePrivateNetworkCmd.RunE(DeletePrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
