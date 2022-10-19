package publicnetwork

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreatePublicNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkCreate := generators.Generate[networkapi.PublicNetworkCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, publicNetworkCreate)

	// What the server should return.
	createdPublicNetwork := generators.Generate[networkapi.PublicNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksPost(gomock.Eq(publicNetworkCreate)).
		Return(&createdPublicNetwork, nil)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePublicNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkCreate := generators.Generate[networkapi.PublicNetworkCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, publicNetworkCreate)

	// What the server should return.
	createdPublicNetwork := generators.Generate[networkapi.PublicNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksPost(gomock.Eq(publicNetworkCreate)).
		Return(&createdPublicNetwork, nil)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePublicNetworkFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreatePublicNetworkUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreatePublicNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkCreate := generators.Generate[networkapi.PublicNetworkCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, publicNetworkCreate)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksPost(gomock.Eq(publicNetworkCreate)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
