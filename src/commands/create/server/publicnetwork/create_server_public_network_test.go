package publicnetwork

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"

	"sigs.k8s.io/yaml"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestCreateServerPublicNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	serverPublicNetwork := generators.Generate[bmcapisdk.ServerPublicNetwork]()

	serverPublicNetworkModel := bmcapisdk.ServerPublicNetwork{
		Id:                serverPublicNetwork.Id,
		Ips:               serverPublicNetwork.Ips,
		StatusDescription: serverPublicNetwork.StatusDescription,
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverPublicNetworkModel)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPost(RESOURCEID, gomock.Eq(serverPublicNetwork)).
		Return(&serverPublicNetwork, nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerPublicNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	serverPublicNetwork := generators.Generate[bmcapisdk.ServerPublicNetwork]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPublicNetwork)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPost(RESOURCEID, gomock.Eq(serverPublicNetwork)).
		Return(&serverPublicNetwork, nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
		ReadFile(FILENAME).
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
		ReadFile(FILENAME).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerPublicNetworkFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerPublicNetworkClientFailure(test_framework *testing.T) {
	// Setup
	serverPublicNetwork := generators.Generate[bmcapisdk.ServerPublicNetwork]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPublicNetwork)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPost(RESOURCEID, gomock.Eq(serverPublicNetwork)).
		Return(nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}
