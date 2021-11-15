package server

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/bmcapimodels"
	"phoenixnap.com/pnap-cli/tests/generators"
	"phoenixnap.com/pnap-cli/tests/testutil"

	"gopkg.in/yaml.v2"

	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
)

func TestTagServerSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	tagAssignmentRequests := generators.GenerateTagAssignmentRequests(2)

	tagAssignmentModel_1 := bmcapimodels.TagAssignmentRequest{
		Name:  tagAssignmentRequests[0].Name,
		Value: tagAssignmentRequests[0].Value,
	}

	tagAssignmentModel_2 := bmcapimodels.TagAssignmentRequest{
		Name:  tagAssignmentRequests[1].Name,
		Value: tagAssignmentRequests[1].Value,
	}

	tagAssignmentRequestModels := []bmcapimodels.TagAssignmentRequest{tagAssignmentModel_1, tagAssignmentModel_2}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(tagAssignmentRequestModels)

	Filename = FILENAME

	// What the server should return.
	server := generators.GenerateServer()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq(tagAssignmentRequests)).
		Return(server, WithResponse(200, WithBody(server)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestTagServerEmptyBodySuccessYAML(test_framework *testing.T) {
	//tagAssignmentRequests := []bmcapisdk.TagAssignmentRequest{}

	filecontents := []byte(``)

	Filename = FILENAME

	// What the server should return.
	server := generators.GenerateServer()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, []bmcapisdk.TagAssignmentRequest{}).
		Return(server, WithResponse(200, WithBody(server)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestTagServerSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	tagAssignmentRequests := generators.GenerateTagAssignmentRequests(2)

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(tagAssignmentRequests)

	Filename = FILENAME

	// What the server should return.
	server := generators.GenerateServer()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq(tagAssignmentRequests)).
		Return(server, WithResponse(200, WithBody(server)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestTagServerFileNotFoundFailure(test_framework *testing.T) {

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestTagServerUnmarshallingFailure(test_framework *testing.T) {
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
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "tag server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestTagServerFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'tag server' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "tag server", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestTagServerBackendErrorFailure(test_framework *testing.T) {
	// Setup
	tagAssignmentRequests := generators.GenerateTagAssignmentRequests(2)

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(tagAssignmentRequests)
	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq(tagAssignmentRequests)).
		Return(bmcapisdk.Server{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestTagServerClientFailure(test_framework *testing.T) {
	// Setup
	tagAssignmentRequests := generators.GenerateTagAssignmentRequests(2)

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(tagAssignmentRequests)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq(tagAssignmentRequests)).
		Return(bmcapisdk.Server{}, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "tag server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestTagServerKeycloakFailure(test_framework *testing.T) {
	// Setup
	tagAssignmentRequests := generators.GenerateTagAssignmentRequests(2)

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(tagAssignmentRequests)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq(tagAssignmentRequests)).
		Return(bmcapisdk.Server{}, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
