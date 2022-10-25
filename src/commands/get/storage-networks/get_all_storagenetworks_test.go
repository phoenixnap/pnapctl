package storagenetworks

import (
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
		Return(networkStorageSdk, nil)

	ExpectToPrintSuccess(test_framework, networkStorageTables)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllStorageNetworksClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGet().
		Return(nil, testutil.TestError)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetAllStorageNetworksPrinterFailure(test_framework *testing.T) {
	// What the server should return.
	networkStorageSdk := testutil.GenN(2, generators.Generate[networkstorageapi.StorageNetwork])
	networkStorageTables := iterutils.MapInterface(networkStorageSdk, tables.StorageNetworkTableFromSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGet().
		Return(networkStorageSdk, nil)

	expectedErr := ExpectToPrintFailure(test_framework, networkStorageTables)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
