package storagenetwork

import (
	"encoding/json"
	"errors"
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
		ReadFile(FILENAME, commandName).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePatch(RESOURCEID, gomock.Eq(networkStoragePatch)).
		Return(&networkStorageSdk, WithResponse(200, WithBody(networkStorageSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTable, "patch storage-network").
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
		ReadFile(FILENAME, commandName).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePatch(RESOURCEID, gomock.Eq(networkStoragePatch)).
		Return(&networkStorageSdk, WithResponse(200, WithBody(networkStorageSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTable, "patch storage-network").
		Return(nil)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchStorageNetworkFileNotFoundFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."})

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchStorageNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`Invalid`)

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "patch storage-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchStorageNetworkFileReadingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'patch storage-network' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "patch storage-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchStorageNetworkBackendErrorFailure(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatch := generators.Generate[networkstorageapi.StorageNetworkUpdate]()

	// Assumed contents of the file.
	marshalled, _ := yaml.Marshal(networkStoragePatch)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePatch(RESOURCEID, gomock.Eq(networkStoragePatch)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchStorageNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatch := generators.Generate[networkstorageapi.StorageNetworkUpdate]()

	// Assumed contents of the file.
	marshalled, _ := yaml.Marshal(networkStoragePatch)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePatch(RESOURCEID, gomock.Eq(networkStoragePatch)).
		Return(nil, nil, testutil.TestError)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "patch storage-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchStorageNetworkKeycloakFailure(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatch := generators.Generate[networkstorageapi.StorageNetworkUpdate]()

	// Assumed contents of the file.
	marshalled, _ := yaml.Marshal(networkStoragePatch)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePatch(RESOURCEID, gomock.Eq(networkStoragePatch)).
		Return(nil, nil, testutil.TestKeycloakError)

	// Run command
	err := PatchStorageNetworkCmd.RunE(PatchStorageNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
