package privatenetwork

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

func TestCreatePrivateNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkCreate := generators.Generate[networkapi.PrivateNetworkCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(privateNetworkCreate)

	Filename = FILENAME

	// What the server should return.
	createdPrivateNetwork := generators.Generate[networkapi.PrivateNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksPost(gomock.Eq(privateNetworkCreate)).
		Return(&createdPrivateNetwork, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePrivateNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkCreate := generators.Generate[networkapi.PrivateNetworkCreate]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(privateNetworkCreate)

	Filename = FILENAME

	// What the server should return.
	createdPrivateNetwork := generators.Generate[networkapi.PrivateNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksPost(gomock.Eq(privateNetworkCreate)).
		Return(&createdPrivateNetwork, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePrivateNetworkFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	ExpectFromFileFailure(test_framework)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Expected error
	expectedErr := testutil.TestError

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreatePrivateNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`sshKeys ["1","2","3","4"]`)

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreatePrivateNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkCreate := generators.Generate[networkapi.PrivateNetworkCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(privateNetworkCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksPost(gomock.Eq(privateNetworkCreate)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
