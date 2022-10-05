package volumes

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/networkstoragemodels"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllVolumesSuccess(test_framework *testing.T) {
	// What the server should return.
	volumeSdk := testutil.GenN(2, networkstoragemodels.GenerateVolumeSdk)
	volumeTables := iterutils.MapInterface(volumeSdk, tables.VolumeTableFromSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetVolumes(RESOURCEID).
		Return(volumeSdk, WithResponse(200, WithBody(volumeSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(volumeTables, "get storage-network volumes").
		Return(nil)

	// Run command
	err := GetStorageNetworkVolumesCmd.RunE(GetStorageNetworkVolumesCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllVolumesClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetVolumes(RESOURCEID).
		Return(nil, nil, testutil.TestError)

	// Run command
	err := GetStorageNetworkVolumesCmd.RunE(GetStorageNetworkVolumesCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "get storage-network volumes", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllVolumesKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetVolumes(RESOURCEID).
		Return(nil, nil, testutil.TestKeycloakError)

	// Run command
	err := GetStorageNetworkVolumesCmd.RunE(GetStorageNetworkVolumesCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllVolumesPrinterFailure(test_framework *testing.T) {
	// What the server should return.
	volumeSdk := testutil.GenN(2, networkstoragemodels.GenerateVolumeSdk)
	volumeTables := iterutils.MapInterface(volumeSdk, tables.VolumeTableFromSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageGetVolumes(RESOURCEID).
		Return(volumeSdk, WithResponse(200, WithBody(volumeSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(volumeTables, "get storage-network volumes").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	// Run command
	err := GetStorageNetworkVolumesCmd.RunE(GetStorageNetworkVolumesCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
