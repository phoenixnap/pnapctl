package publicnetwork

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

const deleteResult = "The server is being removed from the specified public network."

func TestDeleteServerPublicNetworkSuccess(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkDelete(RESOURCEID, RESOURCEID).
		Return(deleteResult, nil)

	// Run command
	err := DeleteServerPublicNetworkCmd.RunE(DeleteServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteServerPublicNetworkClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkDelete(RESOURCEID, RESOURCEID).
		Return("", testutil.TestError)

	// Run command
	err := DeleteServerPublicNetworkCmd.RunE(DeleteServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteServerPublicNetworkKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkDelete(RESOURCEID, RESOURCEID).
		Return("", testutil.TestKeycloakError)

	// Run command
	err := DeleteServerPublicNetworkCmd.RunE(DeleteServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
