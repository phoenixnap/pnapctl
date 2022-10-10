package storagenetworks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/networkstoragemodels"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetStorageNetworkByIdSuccess(test_framework *testing.T) {
	// What the server should return.
	networkStorageSdk := networkstoragemodels.GenerateStorageNetworkSdk()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetById(RESOURCEID).
		Return(&networkStorageSdk, WithResponse(200, WithBody(networkStorageSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTable, "get storage-networks").
		Return(nil)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetStorageNetworkByIdNotFound(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetById(RESOURCEID).
		Return(nil, WithResponse(404, nil), nil)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'get storage-networks' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetStorageNetworkByIdClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetById(RESOURCEID).
		Return(nil, nil, testutil.TestError)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "get storage-networks", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetStorageNetworkByIdKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetById(RESOURCEID).
		Return(nil, nil, testutil.TestKeycloakError)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetStorageNetworkByIdPrinterFailure(test_framework *testing.T) {
	// What the server should return.
	networkStorageSdk := networkstoragemodels.GenerateStorageNetworkSdk()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetById(RESOURCEID).
		Return(&networkStorageSdk, WithResponse(200, WithBody(networkStorageSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTable, "get storage-networks").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
