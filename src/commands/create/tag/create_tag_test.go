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

func createTagSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	tagCreate := generators.Generate[tagapisdk.TagCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, tagCreate)

	// What the server should return.
	createdTag := generators.Generate[tagapisdk.Tag]()

	// Mocking
	PrepareTagMockClient(test_framework).
		TagPost(gomock.Eq(tagCreate)).
		Return(&createdTag, nil)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateTagSuccessYAML(test_framework *testing.T) {
	createTagSuccess(test_framework, yaml.Marshal)
}

func TestCreateTagSuccessJSON(test_framework *testing.T) {
	createTagSuccess(test_framework, json.Marshal)
}

func TestCreateTagFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateTagUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateTagClientFailure(test_framework *testing.T) {
	// What the client should receive.
	tagCreate := generators.Generate[tagapisdk.TagCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, tagCreate)

	// Mocking
	PrepareTagMockClient(test_framework).
		TagPost(gomock.Eq(tagCreate)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateTagCmd.RunE(CreateTagCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
