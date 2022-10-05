package publicnetwork

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"

	"sigs.k8s.io/yaml"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestCreateServerPublicNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	serverPublicNetwork := servermodels.GenerateServerPublicNetworkSdk()

	serverPublicNetworkModel := bmcapi.ServerPublicNetwork{
		Id:                serverPublicNetwork.Id,
		Ips:               serverPublicNetwork.Ips,
		StatusDescription: serverPublicNetwork.StatusDescription,
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverPublicNetworkModel)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPost(RESOURCEID, gomock.Eq(*serverPublicNetwork)).
		Return(serverPublicNetwork, WithResponse(202, WithBody(serverPublicNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerPublicNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	serverPublicNetwork := servermodels.GenerateServerPublicNetworkSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPublicNetwork)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPost(RESOURCEID, gomock.Eq(*serverPublicNetwork)).
		Return(serverPublicNetwork, WithResponse(202, WithBody(serverPublicNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerPublicNetworkFileNotFoundFailure(test_framework *testing.T) {

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestCreateServerPublicNetworkUnmarshallingFailure(test_framework *testing.T) {
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
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create server-public-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerPublicNetworkFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'create server-public-network' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "create server-public-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerPublicNetworkBackendErrorFailure(test_framework *testing.T) {
	// Setup
	serverPublicNetwork := servermodels.GenerateServerPublicNetworkSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPublicNetwork)
	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPost(RESOURCEID, gomock.Eq(*serverPublicNetwork)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerPublicNetworkClientFailure(test_framework *testing.T) {
	// Setup
	serverPublicNetwork := servermodels.GenerateServerPublicNetworkSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPublicNetwork)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPost(RESOURCEID, gomock.Eq(*serverPublicNetwork)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create server-private-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerPublicNetworkKeycloakFailure(test_framework *testing.T) {
	// Setup
	serverPublicNetwork := servermodels.GenerateServerPublicNetworkSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPublicNetwork)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPost(RESOURCEID, gomock.Eq(*serverPublicNetwork)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
