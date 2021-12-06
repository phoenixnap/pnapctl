package privatenetwork

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func TestCreatePrivateNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkCreate := networkmodels.GeneratePrivateNetworkCreateCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(privateNetworkCreate)

	Filename = FILENAME

	// What the server should return.
	createdPrivateNetwork := networkmodels.GeneratePrivateNetworkSdk()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksPost(gomock.Eq(*privateNetworkCreate.ToSdk())).
		Return(createdPrivateNetwork, WithResponse(201, WithBody(createdPrivateNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePrivateNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkCreate := networkmodels.GeneratePrivateNetworkCreateCli()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(privateNetworkCreate)

	Filename = FILENAME

	// What the server should return.
	createdPrivateNetwork := networkmodels.GeneratePrivateNetworkSdk()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksPost(gomock.Eq(*privateNetworkCreate.ToSdk())).
		Return(createdPrivateNetwork, WithResponse(201, WithBody(createdPrivateNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePrivateNetworkFileNotFoundFailure(test_framework *testing.T) {

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePrivateNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`sshKeys ["1","2","3","4"]`)

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create private-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePrivateNetworkBackendErrorFailure(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkCreate := networkmodels.GeneratePrivateNetworkCreateCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(privateNetworkCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksPost(gomock.Eq(*privateNetworkCreate.ToSdk())).
		Return(networksdk.PrivateNetwork{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePrivateNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkCreate := networkmodels.GeneratePrivateNetworkCreateCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(privateNetworkCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksPost(gomock.Eq(*privateNetworkCreate.ToSdk())).
		Return(networksdk.PrivateNetwork{}, nil, testutil.TestError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create private-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePrivateNetworkKeycloakFailure(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkCreate := networkmodels.GeneratePrivateNetworkCreateCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(privateNetworkCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksPost(gomock.Eq(*privateNetworkCreate.ToSdk())).
		Return(networksdk.PrivateNetwork{}, nil, testutil.TestKeycloakError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePrivateNetworkCmd.RunE(CreatePrivateNetworkCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
