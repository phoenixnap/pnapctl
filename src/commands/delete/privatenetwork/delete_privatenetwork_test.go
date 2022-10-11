package privatenetwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeletePrivateNetworkSuccess(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkDelete(RESOURCEID).
		Return(WithResponse(204, nil), nil)

	// Run command
	err := DeletePrivateNetworkCmd.RunE(DeletePrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeletePrivateNetworkNotFound(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkDelete(RESOURCEID).
		Return(WithResponse(404, nil), nil)

	// Run command
	err := DeletePrivateNetworkCmd.RunE(DeletePrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestDeletePrivateNetworkError(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkDelete(RESOURCEID).
		Return(WithResponse(500, nil), nil)

	// Run command
	err := DeletePrivateNetworkCmd.RunE(DeletePrivateNetworkCmd, []string{RESOURCEID})

	expectedMessage := "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeletePrivateNetworkClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkDelete(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := DeletePrivateNetworkCmd.RunE(DeletePrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeletePrivateNetworkKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkDelete(RESOURCEID).
		Return(nil, testutil.TestKeycloakError)

	// Run command
	err := DeletePrivateNetworkCmd.RunE(DeletePrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
