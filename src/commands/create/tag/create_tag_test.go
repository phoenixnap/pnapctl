package tag

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/tagmodels"
	. "phoenixnap.com/pnapctl/tests/mockhelp"
	"phoenixnap.com/pnapctl/tests/testutil"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
)

func TestCreateTagSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	tagCreate := *tagmodels.TagCreateFromSdk(tagmodels.GenerateTagCreate())

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(tagCreate)

	Filename = FILENAME

	// What the server should return.
	createdTag := *tagmodels.GenerateTag()

	// Mocking
	PrepareTagMockClient(test_framework).
		TagPost(gomock.Eq(*tagCreate.ToSdk())).
		Return(createdTag, WithResponse(201, WithBody(createdTag)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateTagSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	tagCreate := *tagmodels.TagCreateFromSdk(tagmodels.GenerateTagCreate())

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(tagCreate)

	Filename = FILENAME

	// What the server should return.
	createdTag := *tagmodels.GenerateTag()

	// Mocking
	PrepareTagMockClient(test_framework).
		TagPost(gomock.Eq(*tagCreate.ToSdk())).
		Return(createdTag, WithResponse(201, WithBody(createdTag)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateTagFileNotFoundFailure(test_framework *testing.T) {

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateTagUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`sshKeys ["1","2","3","4"]`)

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create tag", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateTagBackendErrorFailure(test_framework *testing.T) {
	// What the client should receive.
	tagCreate := *tagmodels.TagCreateFromSdk(tagmodels.GenerateTagCreate())

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(tagCreate)

	Filename = FILENAME

	// Mocking
	PrepareTagMockClient(test_framework).
		TagPost(gomock.Eq(*tagCreate.ToSdk())).
		Return(tagapisdk.Tag{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateTagClientFailure(test_framework *testing.T) {
	// What the client should receive.
	tagCreate := *tagmodels.TagCreateFromSdk(tagmodels.GenerateTagCreate())

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(tagCreate)

	Filename = FILENAME

	// Mocking
	PrepareTagMockClient(test_framework).
		TagPost(gomock.Eq(*tagCreate.ToSdk())).
		Return(tagapisdk.Tag{}, nil, testutil.TestError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create tag", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateTagKeycloakFailure(test_framework *testing.T) {
	// What the client should receive.
	tagCreate := *tagmodels.TagCreateFromSdk(tagmodels.GenerateTagCreate())

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(tagCreate)

	Filename = FILENAME

	// Mocking
	PrepareTagMockClient(test_framework).
		TagPost(gomock.Eq(*tagCreate.ToSdk())).
		Return(tagapisdk.Tag{}, nil, testutil.TestKeycloakError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
