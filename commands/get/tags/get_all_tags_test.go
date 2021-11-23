package tags

import (
	"errors"
	"testing"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/tables"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetAllTagsSuccess(test_framework *testing.T) {
	tags := generators.GenerateTags(5)

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
		Return([]tagapisdk.Tag{}, WithResponse(200, nil), testutil.TestError)

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
		Return([]tagapisdk.Tag{}, nil, testutil.TestKeycloakError)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllTagsPrinterFailure(test_framework *testing.T) {
	tags := generators.GenerateTags(5)

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