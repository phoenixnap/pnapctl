package server

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"

	"github.com/stretchr/testify/assert"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestDeprovisionServerSuccessYAML(test_framework *testing.T) {
	// Mocking
	result := "Server Deprovisioned"
	requestBody := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(requestBody)

	Filename = FILENAME

	PrepareBmcApiMockClient(test_framework).
		ServerDeprovision(RESOURCEID, gomock.Eq(requestBody)).
		Return(result, WithResponse(200, WithBody(result)), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeprovisionServerSuccessJSON(test_framework *testing.T) {
	// Mocking
	result := "Server Deprovisioned"
	requestBody := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(requestBody)

	Filename = FILENAME

	PrepareBmcApiMockClient(test_framework).
		ServerDeprovision(RESOURCEID, gomock.Eq(requestBody)).
		Return(result, WithResponse(200, WithBody(result)), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeprovisionServerFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeprovisionServerUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`deleteIpBlocks negative`)

	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "deprovision server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeprovisionServerFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'deprovision server' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "deprovision server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeprovisionServerBackendErrorFailure(test_framework *testing.T) {
	// Setup
	// Mocking
	requestBody := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(requestBody)

	Filename = FILENAME

	PrepareBmcApiMockClient(test_framework).
		ServerDeprovision(RESOURCEID, gomock.Eq(requestBody)).
		Return("", WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestDeprovisionServerClientFailure(test_framework *testing.T) {
	// Setup
	requestBody := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(requestBody)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDeprovision(RESOURCEID, gomock.Eq(requestBody)).
		Return("", nil, testutil.TestError)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "deprovision server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeprovisionServerKeycloakFailure(test_framework *testing.T) {
	// Setup
	requestBody := generators.Generate[bmcapisdk.RelinquishIpBlock]()
	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(requestBody)

	Filename = FILENAME

	// Mocking

	PrepareBmcApiMockClient(test_framework).
		ServerDeprovision(RESOURCEID, gomock.Eq(requestBody)).
		Return("", nil, testutil.TestKeycloakError)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
