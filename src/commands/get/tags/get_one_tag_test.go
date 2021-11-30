package tags

import (
	"errors"
	"testing"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/models/tagmodels"
	. "phoenixnap.com/pnapctl/tests/mockhelp"
	"phoenixnap.com/pnapctl/tests/testutil"
)

func TestGetTagSuccess(test_framework *testing.T) {

	tag := tagmodels.GenerateTag()
	var tagTable interface{}
	tagTable = tables.TagFromSdk(*tag)

	PrepareTagMockClient(test_framework).
		TagGetById(RESOURCEID).
		Return(*tag, WithResponse(200, WithBody(tag)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(tagTable, "get tags").
		Return(nil)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetTagNotFound(test_framework *testing.T) {
	PrepareTagMockClient(test_framework).
		TagGetById(RESOURCEID).
		Return(tagapisdk.Tag{}, WithResponse(400, nil), nil)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'get tags' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetTagClientFailure(test_framework *testing.T) {
	PrepareTagMockClient(test_framework).
		TagGetById(RESOURCEID).
		Return(tagapisdk.Tag{}, nil, testutil.TestError)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get tags", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetTagKeycloakFailure(test_framework *testing.T) {
	PrepareTagMockClient(test_framework).
		TagGetById(RESOURCEID).
		Return(tagapisdk.Tag{}, nil, testutil.TestKeycloakError)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetTagPrinterFailure(test_framework *testing.T) {
	tag := tagmodels.GenerateTag()
	tagTable := tables.TagFromSdk(*tag)

	PrepareTagMockClient(test_framework).
		TagGetById(RESOURCEID).
		Return(*tag, WithResponse(200, WithBody(tag)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(tagTable, "get tags").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetTagsCmd.RunE(GetTagsCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
