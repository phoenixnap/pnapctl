package storagenetwork

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func patchStorageNetworkSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	networkStoragePatch := generators.Generate[networkstorageapi.StorageNetworkUpdate]()

	// Assumed contents of the file.
	ExpectFromFileSuccess(test_framework, marshaller, networkStoragePatch)

	Filename = FILENAME

	// What the networkStorageSdk should return.
	networkStorageSdk := generators.Generate[networkstorageapi.StorageNetwork]()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePatch(RESOURCEID, gomock.Eq(networkStoragePatch)).
		Return(&networkStorageSdk, nil)

	ExpectToPrintSuccess(test_framework, networkStorageTable)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchStorageNetworkSuccessYAML(test_framework *testing.T) {
	patchStorageNetworkSuccess(test_framework, yaml.Marshal)
}

func TestPatchStorageNetworkSuccessJSON(test_framework *testing.T) {
	patchStorageNetworkSuccess(test_framework, json.Marshal)
}

func TestPatchStorageNetworkFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchStorageNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestPatchStorageNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatch := generators.Generate[networkstorageapi.StorageNetworkUpdate]()

	// Assumed contents of the file.
	ExpectFromFileSuccess(test_framework, yaml.Marshal, networkStoragePatch)

	Filename = FILENAME

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePatch(RESOURCEID, gomock.Eq(networkStoragePatch)).
		Return(nil, testutil.TestError)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
