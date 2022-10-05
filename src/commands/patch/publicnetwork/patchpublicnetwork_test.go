package publicnetwork

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

func TestPatchPublicNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkModifyCli := networkmodels.GeneratePublicNetworkModifyCli()
	publicNetworkModifySdk := publicNetworkModifyCli.ToSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(publicNetworkModifyCli)

	Filename = FILENAME

	// What the server should return.
	publicNetwork := networkmodels.GeneratePublicNetworkSdk()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkPatch(RESOURCEID, gomock.Eq(*publicNetworkModifySdk)).
		Return(&publicNetwork, WithResponse(201, WithBody(&publicNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchPublicNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkModifyCli := networkmodels.GeneratePublicNetworkModifyCli()
	publicNetworkModifySdk := publicNetworkModifyCli.ToSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(publicNetworkModifySdk)

	Filename = FILENAME

	// What the server should return.
	publicNetwork := networkmodels.GeneratePublicNetworkSdk()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkPatch(RESOURCEID, gomock.Eq(*publicNetworkModifySdk)).
		Return(&publicNetwork, WithResponse(201, WithBody(&publicNetwork)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchPublicNetworkFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchPublicNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`Invalid`)

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "patch public-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchPublicNetworkFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'patch public-network' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "patch public-network", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchPublicNetworkBackendErrorFailure(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkModifyCli := networkmodels.GeneratePublicNetworkModifyCli()
	publicNetworkModifySdk := publicNetworkModifyCli.ToSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(publicNetworkModifySdk)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkPatch(RESOURCEID, gomock.Eq(*publicNetworkModifySdk)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchPublicNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkModifyCli := networkmodels.GeneratePublicNetworkModifyCli()
	publicNetworkModifySdk := publicNetworkModifyCli.ToSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(publicNetworkModifySdk)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkPatch(RESOURCEID, gomock.Eq(*publicNetworkModifySdk)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "patch public-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchPublicNetworkKeycloakFailure(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkModifyCli := networkmodels.GeneratePublicNetworkModifyCli()
	publicNetworkModifySdk := publicNetworkModifyCli.ToSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(publicNetworkModifySdk)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkPatch(RESOURCEID, gomock.Eq(*publicNetworkModifySdk)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
