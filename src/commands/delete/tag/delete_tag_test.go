package tag

import (
	"testing"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/tagmodels"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeleteTagSuccess(test_framework *testing.T) {
	// Mocking
	PrepareTagMockClient(test_framework).
		TagDelete(RESOURCEID).
		Return(tagmodels.GenerateTagsDeleteResultSdk(), WithResponse(200, nil), nil)

	// Run command
	err := DeleteTagCmd.RunE(DeleteTagCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteTagNotFound(test_framework *testing.T) {
	// Mocking
	PrepareTagMockClient(test_framework).
		TagDelete(RESOURCEID).
		Return(tagapisdk.DeleteResult{}, WithResponse(404, nil), nil)

	// Run command
	err := DeleteTagCmd.RunE(DeleteTagCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'delete tag' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestDeleteTagError(test_framework *testing.T) {
	// Mocking
	PrepareTagMockClient(test_framework).
		TagDelete(RESOURCEID).
		Return(tagapisdk.DeleteResult{}, WithResponse(500, nil), nil)

	// Run command
	err := DeleteTagCmd.RunE(DeleteTagCmd, []string{RESOURCEID})

	expectedMessage := "Command 'delete tag' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteTagClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareTagMockClient(test_framework).
		TagDelete(RESOURCEID).
		Return(tagapisdk.DeleteResult{}, nil, testutil.TestError)

	// Run command
	err := DeleteTagCmd.RunE(DeleteTagCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete tag", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteTagKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareTagMockClient(test_framework).
		TagDelete(RESOURCEID).
		Return(tagapisdk.DeleteResult{}, nil, testutil.TestKeycloakError)

	// Run command
	err := DeleteTagCmd.RunE(DeleteTagCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
