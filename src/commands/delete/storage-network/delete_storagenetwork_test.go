package storagenetwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeleteStorageNetworkSuccess(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageDelete(RESOURCEID).
		Return(WithResponse(200, WithBody(nil)), nil)

	// Run command
	err := DeleteStorageNetworkCmd.RunE(DeleteStorageNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteStorageNetworkNotFound(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageDelete(RESOURCEID).
		Return(WithResponse(404, nil), nil)

	// Run command
	err := DeleteStorageNetworkCmd.RunE(DeleteStorageNetworkCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'delete storage-network' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteStorageNetworkError(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageDelete(RESOURCEID).
		Return(WithResponse(500, nil), nil)

	// Run command
	err := DeleteStorageNetworkCmd.RunE(DeleteStorageNetworkCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'delete storage-network' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteStorageNetworkClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageDelete(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := DeleteStorageNetworkCmd.RunE(DeleteStorageNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete storage-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteStorageNetworkKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageDelete(RESOURCEID).
		Return(nil, testutil.TestKeycloakError)

	// Run command
	err := DeleteStorageNetworkCmd.RunE(DeleteStorageNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
