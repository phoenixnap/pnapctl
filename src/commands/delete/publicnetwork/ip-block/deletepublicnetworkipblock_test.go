package ipblock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

var deleteResult = "result"

func TestDeletePublicNetworkIpBlockSuccess(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockDelete(RESOURCEID, RESOURCEID).
		Return(deleteResult, WithResponse(204, WithBody("response")), nil)

	// Run command
	err := DeletePublicNetworkIpBlockCmd.RunE(DeletePublicNetworkIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeletePublicNetworkIpBlockNotFound(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockDelete(RESOURCEID, RESOURCEID).
		Return("", WithResponse(404, nil), nil)

	// Run command
	err := DeletePublicNetworkIpBlockCmd.RunE(DeletePublicNetworkIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	expectedMessage := "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeletePublicNetworkIpBlockError(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockDelete(RESOURCEID, RESOURCEID).
		Return("", WithResponse(500, nil), nil)

	// Run command
	err := DeletePublicNetworkIpBlockCmd.RunE(DeletePublicNetworkIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	expectedMessage := "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeletePublicNetworkIpBlockClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockDelete(RESOURCEID, RESOURCEID).
		Return("", nil, testutil.TestError)

	// Run command
	err := DeletePublicNetworkIpBlockCmd.RunE(DeletePublicNetworkIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeletePublicNetworkIpBlockKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockDelete(RESOURCEID, RESOURCEID).
		Return("", nil, testutil.TestKeycloakError)

	// Run command
	err := DeletePublicNetworkIpBlockCmd.RunE(DeletePublicNetworkIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
