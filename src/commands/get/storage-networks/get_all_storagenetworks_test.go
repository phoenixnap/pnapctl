package storagenetworks

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllStorageNetworksSuccess(test_framework *testing.T) {
	// What the server should return.
	networkStorageSdk := testutil.GenN(2, generators.Generate[networkstorageapi.StorageNetwork])
	networkStorageTables := iterutils.MapInterface(networkStorageSdk, tables.StorageNetworkTableFromSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGet().
		Return(networkStorageSdk, WithResponse(200, WithBody(networkStorageSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTables, "get storage-networks").
		Return(nil)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllStorageNetworksClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGet().
		Return(nil, nil, testutil.TestError)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "get storage-networks", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllStorageNetworksKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGet().
		Return(nil, nil, testutil.TestKeycloakError)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllStorageNetworksPrinterFailure(test_framework *testing.T) {
	// What the server should return.
	networkStorageSdk := testutil.GenN(2, generators.Generate[networkstorageapi.StorageNetwork])
	networkStorageTables := iterutils.MapInterface(networkStorageSdk, tables.StorageNetworkTableFromSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGet().
		Return(networkStorageSdk, WithResponse(200, WithBody(networkStorageSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTables, "get storage-networks").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
