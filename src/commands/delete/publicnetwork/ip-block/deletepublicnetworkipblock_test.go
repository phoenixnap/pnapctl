package ipblock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

var deleteResult = "result"

func TestDeletePublicNetworkIpBlockSuccess(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockDelete(RESOURCEID, RESOURCEID).
		Return(deleteResult, nil)

	// Run command
	err := DeletePublicNetworkIpBlockCmd.RunE(DeletePublicNetworkIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeletePublicNetworkIpBlockClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockDelete(RESOURCEID, RESOURCEID).
		Return("", testutil.TestError)

	// Run command
	err := DeletePublicNetworkIpBlockCmd.RunE(DeletePublicNetworkIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
