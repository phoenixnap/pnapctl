package tags

import (
	"errors"
	"testing"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetTagSuccess(test_framework *testing.T) {

	tag := generators.Generate[tagapisdk.Tag]()
	tagTable := tables.TagFromSdk(tag)

	PrepareTagMockClient(test_framework).
		TagGetById(RESOURCEID).
		Return(&tag, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(tagTable).
		Return(nil)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetTagClientFailure(test_framework *testing.T) {
	PrepareTagMockClient(test_framework).
		TagGetById(RESOURCEID).
		Return(nil, testutil.TestError)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetTagPrinterFailure(test_framework *testing.T) {
	tag := generators.Generate[tagapisdk.Tag]()
	tagTable := tables.TagFromSdk(tag)

	PrepareTagMockClient(test_framework).
		TagGetById(RESOURCEID).
		Return(&tag, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(tagTable).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetTagsCmd.RunE(GetTagsCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
