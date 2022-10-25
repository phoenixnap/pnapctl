package volumes

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

func TestGetAllVolumesSuccess(test_framework *testing.T) {
	// What the server should return.
	volumeSdk := testutil.GenN(2, generators.Generate[networkstorageapi.Volume])
	volumeTables := iterutils.MapInterface(volumeSdk, tables.ShortVolumeTableFromSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetVolumes(RESOURCEID).
		Return(volumeSdk, nil)

	ExpectToPrintSuccess(test_framework, volumeTables)

	// Run command
	err := GetStorageNetworkVolumesCmd.RunE(GetStorageNetworkVolumesCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllVolumesClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetVolumes(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := GetStorageNetworkVolumesCmd.RunE(GetStorageNetworkVolumesCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetAllVolumesPrinterFailure(test_framework *testing.T) {
	// What the server should return.
	volumeSdk := testutil.GenN(2, generators.Generate[networkstorageapi.Volume])
	volumeTables := iterutils.MapInterface(volumeSdk, tables.ShortVolumeTableFromSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetVolumes(RESOURCEID).
		Return(volumeSdk, nil)

	expectedErr := ExpectToPrintFailure(test_framework, volumeTables)

	// Run command
	err := GetStorageNetworkVolumesCmd.RunE(GetStorageNetworkVolumesCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
