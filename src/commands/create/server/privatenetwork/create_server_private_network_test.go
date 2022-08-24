package privatenetwork

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"gopkg.in/yaml.v2"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestCreateServerPrivateNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	serverPrivateNetwork := servermodels.GenerateServerPrivateNetworkSdk()

	serverPrivateNetworkModel := servermodels.ServerPrivateNetwork{
		Id:                serverPrivateNetwork.Id,
		Ips:               serverPrivateNetwork.Ips,
		Dhcp:              serverPrivateNetwork.Dhcp,
		StatusDescription: serverPrivateNetwork.StatusDescription,
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverPrivateNetworkModel)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPrivateNetworkPost(RESOURCEID, gomock.Eq(*serverPrivateNetwork)).
		Return(serverPrivateNetwork, WithResponse(202, WithBody(serverPrivateNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerPrivateNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	serverPrivateNetwork := servermodels.GenerateServerPrivateNetworkSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPrivateNetwork)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPrivateNetworkPost(RESOURCEID, gomock.Eq(*serverPrivateNetwork)).
		Return(serverPrivateNetwork, WithResponse(202, WithBody(serverPrivateNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerPrivateNetworkFileNotFoundFailure(test_framework *testing.T) {

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestCreateServerPrivateNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`Name: desc`)

	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create server-private-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerPrivateNetworkFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'create server-private-network' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "create server-private-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerPrivateNetworkBackendErrorFailure(test_framework *testing.T) {
	// Setup
	serverPrivateNetwork := servermodels.GenerateServerPrivateNetworkSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPrivateNetwork)
	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPrivateNetworkPost(RESOURCEID, gomock.Eq(*serverPrivateNetwork)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerPrivateNetworkClientFailure(test_framework *testing.T) {
	// Setup
	serverPrivateNetwork := servermodels.GenerateServerPrivateNetworkSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPrivateNetwork)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPrivateNetworkPost(RESOURCEID, gomock.Eq(*serverPrivateNetwork)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create server-private-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerPrivateNetworkKeycloakFailure(test_framework *testing.T) {
	// Setup
	serverPrivateNetwork := servermodels.GenerateServerPrivateNetworkSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPrivateNetwork)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPrivateNetworkPost(RESOURCEID, gomock.Eq(*serverPrivateNetwork)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
