package ipblock

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreatePublicNetworkIpBlockSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	ipBlockCreate := networkmodels.GeneratePublicNetworkIpBlockCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockCreate)

	Filename = FILENAME

	// What the server should return.
	createdIpBlock := networkmodels.GeneratePublicNetworkIpBlockSdk()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockPost(RESOURCEID, gomock.Eq(*ipBlockCreate.ToSdk())).
		Return(&createdIpBlock, WithResponse(200, WithBody(createdIpBlock)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePublicNetworkIpBlockSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	ipBlockCreate := networkmodels.GeneratePublicNetworkIpBlockCli()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockCreate)

	Filename = FILENAME

	// What the server should return.
	createdIpBlock := networkmodels.GeneratePublicNetworkIpBlockSdk()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockPost(RESOURCEID, gomock.Eq(*ipBlockCreate.ToSdk())).
		Return(&createdIpBlock, WithResponse(200, WithBody(createdIpBlock)), nil).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePublicNetworkIpBlockFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePublicNetworkIpBlockUnmarshallingFailure(test_framework *testing.T) {
	// Setup
	filecontents := []byte(`invalid`)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create public-network ip-block", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePublicNetworkIpBlockFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'create public-network ip-block' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "create public-network ip-block", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePublicNetworkIpBlockBackendErrorFailure(test_framework *testing.T) {
	// What the client should receive.
	ipBlockCreate := networkmodels.GeneratePublicNetworkIpBlockCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockPost(RESOURCEID, gomock.Eq(*ipBlockCreate.ToSdk())).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePublicNetworkIpBlockClientFailure(test_framework *testing.T) {
	// What the client should receive.
	ipBlockCreate := networkmodels.GeneratePublicNetworkIpBlockCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockPost(RESOURCEID, gomock.Eq(*ipBlockCreate.ToSdk())).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create public-network ip-block", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreatePublicNetworkIpBlockKeycloakFailure(test_framework *testing.T) {
	// What the client should receive.
	ipBlockCreate := networkmodels.GeneratePublicNetworkIpBlockCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockPost(RESOURCEID, gomock.Eq(*ipBlockCreate.ToSdk())).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
