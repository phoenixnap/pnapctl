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
		Return(WithResponse(204, WithBody("response")), nil)

	// Run command
	err := DeletePublicNetworkCmd.RunE(DeletePublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeletePublicNetworkIpBlockNotFound(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkDelete(RESOURCEID).
		Return(WithResponse(404, nil), nil)

	// Run command
	err := DeletePublicNetworkCmd.RunE(DeletePublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'delete public-network' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeletePublicNetworkIpBlockError(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkDelete(RESOURCEID).
		Return(WithResponse(500, nil), nil)

	// Run command
	err := DeletePublicNetworkCmd.RunE(DeletePublicNetworkCmd, []string{RESOURCEID})

	expectedMessage := "Command 'delete public-network' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeletePublicNetworkIpBlockClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkDelete(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := DeletePublicNetworkCmd.RunE(DeletePublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete public-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeletePublicNetworkIpBlockKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkDelete(RESOURCEID).
		Return(nil, testutil.TestKeycloakError)

	// Run command
	err := DeletePublicNetworkCmd.RunE(DeletePublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
