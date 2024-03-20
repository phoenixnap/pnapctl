package tag

import (
	"testing"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeleteTagSuccess(test_framework *testing.T) {
	// Mocking
	PrepareTagMockClient(test_framework).
		TagDelete(RESOURCEID).
		Return(testutil.AsPointer(generators.Generate[tagapisdk.DeleteResult]()), nil)

	// Run command
	err := DeleteTagCmd.RunE(DeleteTagCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteTagClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareTagMockClient(test_framework).
		TagDelete(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := DeleteTagCmd.RunE(DeleteTagCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
