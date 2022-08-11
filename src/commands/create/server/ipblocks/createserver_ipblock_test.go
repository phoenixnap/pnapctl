package ipblocks

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestCreateServerIpBlockSuccessYAML(test_framework *testing.T) {
	serverIpBlockCli := servermodels.GenerateServerIpBlockCli()
	serverIpBlockSdk := serverIpBlockCli.ToSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverIpBlockCli)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockPost(RESOURCEID, gomock.Eq(serverIpBlockSdk)).
		Return(&serverIpBlockSdk, WithResponse(202, WithBody(serverIpBlockSdk)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerIpBlockSuccessJSON(test_framework *testing.T) {
	serverIpBlockCli := servermodels.GenerateServerIpBlockCli()
	serverIpBlockSdk := serverIpBlockCli.ToSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverIpBlockCli)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockPost(RESOURCEID, gomock.Eq(serverIpBlockSdk)).
		Return(&serverIpBlockSdk, WithResponse(202, WithBody(serverIpBlockSdk)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerIpBlockFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestCreateServerIpBlockUnmarshallingFailure(test_framework *testing.T) {
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
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create server-ip-block", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerIpBlockFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'create server-private-network' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "create server-private-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerIpBlockBackendErrorFailure(test_framework *testing.T) {
	// Setup
	serverIpBlockSdk := servermodels.GenerateServerIpBlockSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverIpBlockSdk)
	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockPost(RESOURCEID, gomock.Eq(serverIpBlockSdk)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerIpBlockClientFailure(test_framework *testing.T) {
	// Setup
	serverIpBlockSdk := servermodels.GenerateServerIpBlockSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverIpBlockSdk)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockPost(RESOURCEID, gomock.Eq(serverIpBlockSdk)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create server-ip-block", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerIpBlockKeycloakFailure(test_framework *testing.T) {
	// Setup
	serverIpBlockSdk := servermodels.GenerateServerIpBlockSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverIpBlockSdk)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockPost(RESOURCEID, gomock.Eq(serverIpBlockSdk)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
