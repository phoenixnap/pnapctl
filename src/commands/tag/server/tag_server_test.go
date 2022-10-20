package server

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"sigs.k8s.io/yaml"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func tagServerSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	tagAssignmentRequests := testutil.GenN(2, generators.Generate[bmcapisdk.TagAssignmentRequest])

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, tagAssignmentRequests)

	// What the server should return.
	server := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq(tagAssignmentRequests)).
		Return(&server, nil)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestTagServerSuccessYAML(test_framework *testing.T) {
	tagServerSuccess(test_framework, yaml.Marshal)
}

func TestTagServerSuccessJSON(test_framework *testing.T) {
	tagServerSuccess(test_framework, json.Marshal)
}

func TestTagServerSuccessEmptyBody(test_framework *testing.T) {
	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, nil)

	// What the server should return.
	server := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq([]bmcapisdk.TagAssignmentRequest{})).
		Return(&server, nil)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestTagServerFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestTagServerUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestTagServerClientFailure(test_framework *testing.T) {
	// Setup
	tagAssignmentRequests := testutil.GenN(2, generators.Generate[bmcapisdk.TagAssignmentRequest])

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, tagAssignmentRequests)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTag(RESOURCEID, gomock.Eq(tagAssignmentRequests)).
		Return(nil, testutil.TestError)

	// Run command
	err := TagServerCmd.RunE(TagServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
