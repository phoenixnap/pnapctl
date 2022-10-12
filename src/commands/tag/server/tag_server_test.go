package server

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"sigs.k8s.io/yaml"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestTagServerSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	tagAssignmentRequests := testutil.GenN(2, generators.Generate[bmcapisdk.TagAssignmentRequest])

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(tagAssignmentRequests)

	Filename = FILENAME

	// What the server should return.
	server := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq(tagAssignmentRequests)).
		Return(&server, nil).
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
	filecontents := []byte(``)

	Filename = FILENAME

	// What the server should return.
	server := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, []bmcapisdk.TagAssignmentRequest{}).
		Return(&server, nil).
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
	tagAssignmentRequests := testutil.GenN(2, generators.Generate[bmcapisdk.TagAssignmentRequest])

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(tagAssignmentRequests)

	Filename = FILENAME

	// What the server should return.
	server := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq(tagAssignmentRequests)).
		Return(&server, nil).
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
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

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
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestTagServerClientFailure(test_framework *testing.T) {
	// Setup
	tagAssignmentRequests := testutil.GenN(2, generators.Generate[bmcapisdk.TagAssignmentRequest])

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(tagAssignmentRequests)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq(tagAssignmentRequests)).
		Return(nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestTagServerKeycloakFailure(test_framework *testing.T) {
	// Setup
	tagAssignmentRequests := testutil.GenN(2, generators.Generate[bmcapisdk.TagAssignmentRequest])

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(tagAssignmentRequests)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq(tagAssignmentRequests)).
		Return(nil, testutil.TestKeycloakError).
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
