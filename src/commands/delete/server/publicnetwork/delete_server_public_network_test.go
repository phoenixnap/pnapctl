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
		Return(deleteResult, WithResponse(202, nil), nil)

	// Run command
	err := DeleteServerPublicNetworkCmd.RunE(DeleteServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteServerPublicNetworkNotFound(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkDelete(RESOURCEID, RESOURCEID).
		Return("", WithResponse(404, nil), nil)

	// Run command
	err := DeleteServerPublicNetworkCmd.RunE(DeleteServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	expectedMessage := "Command 'delete server-public-network' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestDeleteServerPublicNetworkError(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkDelete(RESOURCEID, RESOURCEID).
		Return("", WithResponse(500, nil), nil)

	// Run command
	err := DeleteServerPublicNetworkCmd.RunE(DeleteServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	expectedMessage := "Command 'delete server-public-network' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteServerPublicNetworkClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkDelete(RESOURCEID, RESOURCEID).
		Return("", nil, testutil.TestError)

	// Run command
	err := DeleteServerPublicNetworkCmd.RunE(DeleteServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete server-public-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteServerPublicNetworkKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkDelete(RESOURCEID, RESOURCEID).
		Return("", nil, testutil.TestKeycloakError)

	// Run command
	err := DeleteServerPublicNetworkCmd.RunE(DeleteServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}