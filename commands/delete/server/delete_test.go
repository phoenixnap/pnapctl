package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
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
