package server

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels"
	"phoenixnap.com/pnapctl/testsupport/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"gopkg.in/yaml.v2"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestReserveServerSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	serverReserve := generators.GenerateServerReserve()

	serverReserveModel := bmcapimodels.ServerReserve{
		PricingModel: serverReserve.PricingModel,
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverReserveModel)

	Filename = FILENAME

	// What the server should return.
	server := generators.GenerateServer()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReserve(RESOURCEID, gomock.Eq(serverReserve)).
		Return(server, WithResponse(200, WithBody(server)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestReserveServerSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	serverReserve := generators.GenerateServerReserve()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverReserve)

	Filename = FILENAME

	// What the server should return.
	server := generators.GenerateServer()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReserve(RESOURCEID, gomock.Eq(serverReserve)).
		Return(server, WithResponse(200, WithBody(server)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestReserveServerFileNotFoundFailure(test_framework *testing.T) {

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestReserveServerUnmarshallingFailure(test_framework *testing.T) {
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
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "reserve server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestReserveServerFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'reserve server' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "reserve server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestReserveServerBackendErrorFailure(test_framework *testing.T) {
	// Setup
	serverReserve := generators.GenerateServerReserve()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverReserve)
	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReserve(RESOURCEID, gomock.Eq(serverReserve)).
		Return(bmcapisdk.Server{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestReserveServerClientFailure(test_framework *testing.T) {
	// Setup
	serverReserve := generators.GenerateServerReserve()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverReserve)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReserve(RESOURCEID, gomock.Eq(serverReserve)).
		Return(bmcapisdk.Server{}, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "reserve server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestReserveServerKeycloakFailure(test_framework *testing.T) {
	// Setup
	serverReserve := generators.GenerateServerReserve()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverReserve)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReserve(RESOURCEID, gomock.Eq(serverReserve)).
		Return(bmcapisdk.Server{}, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
