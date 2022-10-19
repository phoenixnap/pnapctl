package storagenetwork

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestPatchStorageNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatch := generators.Generate[networkstorageapi.StorageNetworkUpdate]()

	// Assumed contents of the file.
	marshalled, _ := yaml.Marshal(networkStoragePatch)

	Filename = FILENAME

	// What the networkStorageSdk should return.
	networkStorageSdk := generators.Generate[networkstorageapi.StorageNetwork]()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePatch(RESOURCEID, gomock.Eq(networkStoragePatch)).
		Return(&networkStorageSdk, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTable).
		Return(nil)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchStorageNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatch := generators.Generate[networkstorageapi.StorageNetworkUpdate]()

	// Assumed contents of the file.
	marshalled, _ := json.Marshal(networkStoragePatch)

	Filename = FILENAME

	// What the networkStorageSdk should return.
	networkStorageSdk := generators.Generate[networkstorageapi.StorageNetwork]()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePatch(RESOURCEID, gomock.Eq(networkStoragePatch)).
		Return(&networkStorageSdk, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTable).
		Return(nil)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchStorageNetworkFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, testutil.TestError)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := testutil.TestError

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchStorageNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`Invalid`)

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestPatchStorageNetworkFileReadingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchStorageNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatch := generators.Generate[networkstorageapi.StorageNetworkUpdate]()

	// Assumed contents of the file.
	marshalled, _ := yaml.Marshal(networkStoragePatch)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(marshalled, nil)

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
