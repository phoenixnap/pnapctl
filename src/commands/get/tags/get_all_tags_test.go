package tags

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllTagsSuccess(test_framework *testing.T) {
	tags := generators.GenerateTagListSdk(5)

	var taglist []interface{}

	for _, x := range tags {
		taglist = append(taglist, tables.TagFromSdk(x))
	}

	// Mocking
	PrepareTagMockClient(test_framework).
		TagsGet("").
		Return(tags, WithResponse(200, WithBody(tags)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(taglist, "get tags").
		Return(nil)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllTagsClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareTagMockClient(test_framework).
		TagsGet("").
		Return(nil, WithResponse(200, nil), testutil.TestError)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get servers", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllTagsKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareTagMockClient(test_framework).
		TagsGet("").
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllTagsPrinterFailure(test_framework *testing.T) {
	tags := generators.GenerateTagListSdk(5)

	var taglist []interface{}

	for _, x := range tags {
		taglist = append(taglist, tables.TagFromSdk(x))
	}

	PrepareTagMockClient(test_framework).
		TagsGet("").
		Return(tags, WithResponse(200, WithBody(tags)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(taglist, "get tags").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetTagsCmd.RunE(GetTagsCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
