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

func TestCreateStorageNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	networkStorageCreate := generators.Generate[networkstorageapi.StorageNetworkCreate]()

	// Assumed contents of the file.
	marshalled, _ := yaml.Marshal(networkStorageCreate)

	Filename = FILENAME

	// What the networkStorageSdk should return.
	networkStorageSdk := generators.Generate[networkstorageapi.StorageNetwork]()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePost(gomock.Eq(networkStorageCreate)).
		Return(&networkStorageSdk, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTable).
		Return(nil)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateStorageNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	networkStorageCreate := generators.Generate[networkstorageapi.StorageNetworkCreate]()

	// Assumed contents of the file.
	marshalled, _ := json.Marshal(networkStorageCreate)

	Filename = FILENAME

	// What the networkStorageSdk should return.
	networkStorageSdk := generators.Generate[networkstorageapi.StorageNetwork]()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePost(gomock.Eq(networkStorageCreate)).
		Return(&networkStorageSdk, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(networkStorageTable).
		Return(nil)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateStorageNetworkFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Expected error
	expectedErr := testutil.TestError

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateStorageNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`Invalid`)

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateStorageNetworkFileReadingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateStorageNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	networkStorageCreate := generators.Generate[networkstorageapi.StorageNetworkCreate]()

	// Assumed contents of the file.
	marshalled, _ := yaml.Marshal(networkStorageCreate)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(marshalled, nil)

	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePost(gomock.Eq(networkStorageCreate)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
