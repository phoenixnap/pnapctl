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

func createStorageNetworkSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	networkStorageCreate := generators.Generate[networkstorageapi.StorageNetworkCreate]()

	// Assumed contents of the file.
	ExpectFromFileSuccess(test_framework, marshaller, networkStorageCreate)

	Filename = FILENAME

	// What the networkStorageSdk should return.
	networkStorageSdk := generators.Generate[networkstorageapi.StorageNetwork]()
	networkStorageTable := tables.StorageNetworkTableFromSdk(networkStorageSdk)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStoragePost(gomock.Eq(networkStorageCreate)).
		Return(&networkStorageSdk, nil)

	ExpectToPrintSuccess(test_framework, networkStorageTable)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateStorageNetworkSuccessYAML(test_framework *testing.T) {
	createStorageNetworkSuccess(test_framework, yaml.Marshal)
}

func TestCreateStorageNetworkSuccessJSON(test_framework *testing.T) {
	createStorageNetworkSuccess(test_framework, json.Marshal)
}

func TestCreateStorageNetworkFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateStorageNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateStorageNetworkCmd.RunE(CreateStorageNetworkCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateStorageNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	networkStorageCreate := generators.Generate[networkstorageapi.StorageNetworkCreate]()

	// Assumed contents of the file.
	ExpectFromFileSuccess(test_framework, yaml.Marshal, networkStorageCreate)

	Filename = FILENAME

	// Mocking
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
