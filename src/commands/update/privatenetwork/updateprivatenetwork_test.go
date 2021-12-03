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
	. "phoenixnap.com/pnapctl/tests/mockhelp"
	"phoenixnap.com/pnapctl/tests/testutil"

	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func TestUpdatePrivateNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkUpdate := networkmodels.GeneratePrivateNetworkModifyCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(privateNetworkUpdate)

	Filename = FILENAME

	// What the server should return.
	privateNetwork := networkmodels.GeneratePrivateNetworkSdk()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(*privateNetworkUpdate.ToSdk())).
		Return(privateNetwork, WithResponse(200, WithBody(privateNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestUpdatePrivateNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkUpdate := networkmodels.GeneratePrivateNetworkModifyCli()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(privateNetworkUpdate)

	Filename = FILENAME

	// What the server should return.
	privateNetwork := networkmodels.GeneratePrivateNetworkSdk()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(*privateNetworkUpdate.ToSdk())).
		Return(privateNetwork, WithResponse(200, WithBody(privateNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestUpdatePrivateNetworkFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestUpdatePrivateNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`name this is a bad name`)

	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "update private-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestUpdatePrivateNetworkFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'update private-network' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "update private-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestUpdatePrivateNetworkBackendErrorFailure(test_framework *testing.T) {
	// Setup
	privateNetworkUpdate := networkmodels.GeneratePrivateNetworkModifyCli()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(privateNetworkUpdate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(*privateNetworkUpdate.ToSdk())).
		Return(networksdk.PrivateNetwork{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestUpdatePrivateNetworkClientFailure(test_framework *testing.T) {
	// Setup
	privateNetworkUpdate := networkmodels.GeneratePrivateNetworkModifyCli()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(privateNetworkUpdate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(*privateNetworkUpdate.ToSdk())).
		Return(networksdk.PrivateNetwork{}, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "update private-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestUpdatePrivateNetworkKeycloakFailure(test_framework *testing.T) {
	// Setup
	privateNetworkUpdate := networkmodels.GeneratePrivateNetworkModifyCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(privateNetworkUpdate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(*privateNetworkUpdate.ToSdk())).
		Return(networksdk.PrivateNetwork{}, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
