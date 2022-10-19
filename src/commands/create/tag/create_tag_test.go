package tag

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreateTagSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	tagCreate := generators.Generate[tagapisdk.TagCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(tagCreate)

	Filename = FILENAME

	// What the server should return.
	createdTag := generators.Generate[tagapisdk.Tag]()

	// Mocking
	PrepareTagMockClient(test_framework).
		TagPost(gomock.Eq(tagCreate)).
		Return(&createdTag, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateTagSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	tagCreate := generators.Generate[tagapisdk.TagCreate]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(tagCreate)

	Filename = FILENAME

	// What the server should return.
	createdTag := generators.Generate[tagapisdk.Tag]()

	// Mocking
	PrepareTagMockClient(test_framework).
		TagPost(gomock.Eq(tagCreate)).
		Return(&createdTag, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateTagFileNotFoundFailure(test_framework *testing.T) {

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."})

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateTagUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`sshKeys ["1","2","3","4"]`)

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateTagClientFailure(test_framework *testing.T) {
	// What the client should receive.
	tagCreate := generators.Generate[tagapisdk.TagCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(tagCreate)

	Filename = FILENAME

	// Mocking
	PrepareTagMockClient(test_framework).
		TagPost(gomock.Eq(tagCreate)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
