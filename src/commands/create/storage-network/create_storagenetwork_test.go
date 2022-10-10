package storagenetwork

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/networkstoragemodels"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestCreateStorageNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatchCli := networkstoragemodels.GenerateStorageNetworkCreateCli()
	networkStoragePatchSdk := networkStoragePatchCli.ToSdk()

	// Assumed contents of the file.
	marshalled, _ := yaml.Marshal(networkStoragePatchCli)

	Filename = FILENAME

	// What the networkStorageSdk should return.
	networkStorageSdk := networkstoragemodels.GenerateStorageNetworkSdk()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePost(gomock.Eq(networkStoragePatchSdk)).
		Return(&networkStorageSdk, WithResponse(200, WithBody(networkStorageSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTable, "create storage-network").
		Return(nil)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateStorageNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatchCli := networkstoragemodels.GenerateStorageNetworkCreateCli()
	networkStoragePatchSdk := networkStoragePatchCli.ToSdk()

	// Assumed contents of the file.
	marshalled, _ := json.Marshal(networkStoragePatchCli)

	Filename = FILENAME

	// What the networkStorageSdk should return.
	networkStorageSdk := networkstoragemodels.GenerateStorageNetworkSdk()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePost(gomock.Eq(networkStoragePatchSdk)).
		Return(&networkStorageSdk, WithResponse(200, WithBody(networkStorageSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTable, "create storage-network").
		Return(nil)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateStorageNetworkFileNotFoundFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."})

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateStorageNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`Invalid`)

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create storage-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateStorageNetworkFileReadingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'create storage-network' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "create storage-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateStorageNetworkBackendErrorFailure(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatchCli := networkstoragemodels.GenerateStorageNetworkCreateCli()
	networkStoragePatchSdk := networkStoragePatchCli.ToSdk()

	// Assumed contents of the file.
	marshalled, _ := yaml.Marshal(networkStoragePatchCli)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePost(gomock.Eq(networkStoragePatchSdk)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateStorageNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatchCli := networkstoragemodels.GenerateStorageNetworkCreateCli()
	networkStoragePatchSdk := networkStoragePatchCli.ToSdk()

	// Assumed contents of the file.
	marshalled, _ := yaml.Marshal(networkStoragePatchCli)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePost(gomock.Eq(networkStoragePatchSdk)).
		Return(nil, nil, testutil.TestError)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create storage-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateStorageNetworkKeycloakFailure(test_framework *testing.T) {
	// What the client should receive.
	networkStoragePatchCli := networkstoragemodels.GenerateStorageNetworkCreateCli()
	networkStoragePatchSdk := networkStoragePatchCli.ToSdk()

	// Assumed contents of the file.
	marshalled, _ := yaml.Marshal(networkStoragePatchCli)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePost(gomock.Eq(networkStoragePatchSdk)).
		Return(nil, nil, testutil.TestKeycloakError)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
