package publicnetwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeletePublicNetworkIpBlockSuccess(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkDelete(RESOURCEID).
		Return(nil)

	// Run command
	err := DeletePublicNetworkCmd.RunE(DeletePublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeletePublicNetworkIpBlockClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkDelete(RESOURCEID).
		Return(testutil.TestError)

	// Run command
	err := DeletePublicNetworkCmd.RunE(DeletePublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}
