package privatenetwork

import (
	"encoding/json"
	"errors"
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

func TestUpdatePrivateNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkUpdate := generators.Generate[networkapi.PrivateNetworkModify]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(privateNetworkUpdate)

	Filename = FILENAME

	// What the server should return.
	privateNetwork := generators.Generate[networkapi.PrivateNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(privateNetworkUpdate)).
		Return(&privateNetwork, WithResponse(200, WithBody(privateNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestUpdatePrivateNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	privateNetworkUpdate := generators.Generate[networkapi.PrivateNetworkModify]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(privateNetworkUpdate)

	Filename = FILENAME

	// What the server should return.
	privateNetwork := generators.Generate[networkapi.PrivateNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(privateNetworkUpdate)).
		Return(&privateNetwork, WithResponse(200, WithBody(privateNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
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
		ReadFile(FILENAME, commandName).
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
		ReadFile(FILENAME, commandName).
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
		ReadFile(FILENAME, commandName).
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
	privateNetworkUpdate := generators.Generate[networkapi.PrivateNetworkModify]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(privateNetworkUpdate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(privateNetworkUpdate)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
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
	privateNetworkUpdate := generators.Generate[networkapi.PrivateNetworkModify]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(privateNetworkUpdate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(privateNetworkUpdate)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
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
	privateNetworkUpdate := generators.Generate[networkapi.PrivateNetworkModify]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(privateNetworkUpdate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(privateNetworkUpdate)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
