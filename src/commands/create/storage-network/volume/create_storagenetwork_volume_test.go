package storagenetwork

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func createStorageNetworkVolumeSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	volumeCreate := generators.Generate[networkstorageapi.VolumeCreate]()

	// Assumed contents of the file.
	ExpectFromFileSuccess(test_framework, marshaller, volumeCreate)

	Filename = FILENAME

	// What the volumeSdk should return.
	volumeSdk := generators.Generate[networkstorageapi.Volume]()
	volumeTable := tables.VolumeTableFromSdk(volumeSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
	NetworkStoragePostVolume(RESOURCEID, gomock.Eq(volumeCreate)).
		Return(&volumeSdk, nil)

	ExpectToPrintSuccess(test_framework, volumeTable)

	// Run command
	err := CreateStorageNetworkVolumeCmd.RunE(CreateStorageNetworkVolumeCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateStorageNetworkVolumeFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateStorageNetworkVolumeCmd.RunE(CreateStorageNetworkVolumeCmd, []string{RESOURCEID})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchStorageNetworkVolumeUnmarshallingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateStorageNetworkVolumeCmd.RunE(CreateStorageNetworkVolumeCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestPatchStorageNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	volumeCreate := generators.Generate[networkstorageapi.VolumeCreate]()

	// Assumed contents of the file.
	ExpectFromFileSuccess(test_framework, yaml.Marshal, volumeCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
	NetworkStoragePostVolume(RESOURCEID, gomock.Eq(volumeCreate)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateStorageNetworkVolumeCmd.RunE(CreateStorageNetworkVolumeCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
