package volumes

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetVolumeByIdSuccess(test_framework *testing.T) {
	// What the server should return.
	volumeSdk := generators.Generate[networkstorageapi.Volume]()
	volumeTable := tables.ShortVolumeTableFromSdk(volumeSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetVolumeById(RESOURCEID, RESOURCEID).
		Return(&volumeSdk, nil)

	ExpectToPrintSuccess(test_framework, volumeTable)

	// Run command
	err := GetStorageNetworkVolumesCmd.RunE(GetStorageNetworkVolumesCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetVolumeByIdClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetVolumeById(RESOURCEID, RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := GetStorageNetworkVolumesCmd.RunE(GetStorageNetworkVolumesCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetVolumeByIdPrinterFailure(test_framework *testing.T) {
	// What the server should return.
	volumeSdk := generators.Generate[networkstorageapi.Volume]()
	volumeTable := tables.ShortVolumeTableFromSdk(volumeSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetVolumeById(RESOURCEID, RESOURCEID).
		Return(&volumeSdk, nil)

	expectedErr := ExpectToPrintFailure(test_framework, volumeTable)

	// Run command
	err := GetStorageNetworkVolumesCmd.RunE(GetStorageNetworkVolumesCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
