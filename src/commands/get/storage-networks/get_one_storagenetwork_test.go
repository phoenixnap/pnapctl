package storagenetworks

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetStorageNetworkByIdSuccess(test_framework *testing.T) {
	// What the server should return.
	networkStorageSdk := generators.Generate[networkstorageapi.StorageNetwork]()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetById(RESOURCEID).
		Return(&networkStorageSdk, nil)

	ExpectToPrintSuccess(test_framework, networkStorageTable)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetStorageNetworkByIdClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetById(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetStorageNetworkByIdPrinterFailure(test_framework *testing.T) {
	// What the server should return.
	networkStorageSdk := generators.Generate[networkstorageapi.StorageNetwork]()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetById(RESOURCEID).
		Return(&networkStorageSdk, nil)

	expectedErr := ExpectToPrintFailure(test_framework, networkStorageTable)

	// Run command
	err := GetStorageNetworksCmd.RunE(GetStorageNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
