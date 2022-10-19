package tags

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllTagsSuccess(test_framework *testing.T) {
	tags := testutil.GenN(1, generators.Generate[tagapi.Tag])
	taglist := iterutils.MapInterface(tags, tables.TagFromSdk)

	// Mocking
	PrepareTagMockClient(test_framework).
		TagsGet("").
		Return(tags, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(taglist).
		Return(nil)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllTagsClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareTagMockClient(test_framework).
		TagsGet("").
		Return(nil, testutil.TestError)

	err := GetTagsCmd.RunE(GetTagsCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetAllTagsPrinterFailure(test_framework *testing.T) {
	tags := testutil.GenN(1, generators.Generate[tagapi.Tag])
	taglist := iterutils.MapInterface(tags, tables.TagFromSdk)

	PrepareTagMockClient(test_framework).
		TagsGet("").
		Return(tags, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(taglist).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetTagsCmd.RunE(GetTagsCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
