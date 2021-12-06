package server

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/testsupport/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeleteServerSuccess(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDelete(RESOURCEID).
		Return(generators.GenerateBmcApiDeleteResult(), WithResponse(200, nil), nil)

	// Run command
	err := DeleteServerCmd.RunE(DeleteServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteServerNotFound(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDelete(RESOURCEID).
		Return(bmcapisdk.DeleteResult{}, WithResponse(404, nil), nil)

	// Run command
	err := DeleteServerCmd.RunE(DeleteServerCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'delete server' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestDeleteServerError(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDelete(RESOURCEID).
		Return(bmcapisdk.DeleteResult{}, WithResponse(500, nil), nil)

	// Run command
	err := DeleteServerCmd.RunE(DeleteServerCmd, []string{RESOURCEID})

	expectedMessage := "Command 'delete server' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteServerClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDelete(RESOURCEID).
		Return(bmcapisdk.DeleteResult{}, nil, testutil.TestError)

	// Run command
	err := DeleteServerCmd.RunE(DeleteServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteServerKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDelete(RESOURCEID).
		Return(bmcapisdk.DeleteResult{}, nil, testutil.TestKeycloakError)

	// Run command
	err := DeleteServerCmd.RunE(DeleteServerCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
