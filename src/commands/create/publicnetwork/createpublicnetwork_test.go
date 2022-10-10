package publicnetwork

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkapi"
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
	yamlmarshal, _ := yaml.Marshal(publicNetworkCreate)

	Filename = FILENAME

	// What the server should return.
	createdPublicNetwork := generators.Generate[networkapi.PublicNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksPost(gomock.Eq(publicNetworkCreate)).
		Return(&createdPublicNetwork, WithResponse(200, WithBody(createdPublicNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePublicNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkCreate := generators.Generate[networkapi.PublicNetworkCreate]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(publicNetworkCreate)

	Filename = FILENAME

	// What the server should return.
	createdPublicNetwork := generators.Generate[networkapi.PublicNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksPost(gomock.Eq(publicNetworkCreate)).
		Return(&createdPublicNetwork, WithResponse(200, WithBody(createdPublicNetwork)), nil).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePublicNetworkFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePublicNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Setup
	filecontents := []byte(`invalid`)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create public-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePublicNetworkFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'create public-network' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "create public-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePublicNetworkBackendErrorFailure(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkCreate := generators.Generate[networkapi.PublicNetworkCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(publicNetworkCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksPost(gomock.Eq(publicNetworkCreate)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePublicNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkCreate := generators.Generate[networkapi.PublicNetworkCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(publicNetworkCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksPost(gomock.Eq(publicNetworkCreate)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create public-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePublicNetworkKeycloakFailure(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkCreate := generators.Generate[networkapi.PublicNetworkCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(publicNetworkCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksPost(gomock.Eq(publicNetworkCreate)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
