package privatenetwork

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func createPrivateNetworkSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	privateNetworkCreate := generators.Generate[networkapi.PrivateNetworkCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, privateNetworkCreate)

	// What the server should return.
	createdPrivateNetwork := generators.Generate[networkapi.PrivateNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksPost(gomock.Eq(privateNetworkCreate), gomock.Eq(false)).
		Return(&createdPrivateNetwork, nil)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePrivateNetworkSuccessYAML(test_framework *testing.T) {
	createPrivateNetworkSuccess(test_framework, yaml.Marshal)
}

func TestCreatePrivateNetworkSuccessJSON(test_framework *testing.T) {
	createPrivateNetworkSuccess(test_framework, json.Marshal)
}

func TestCreatePrivateNetworkFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreatePrivateNetworkUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreatePrivateNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkCreate := generators.Generate[networkapi.PrivateNetworkCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, privateNetworkCreate)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksPost(gomock.Eq(privateNetworkCreate), gomock.Eq(false)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
