package delete

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	delete "phoenixnap.com/pnap-cli/commands/delete/server"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func deleteSetup() {
	URL = "servers/" + SERVERID
}

func TestDeleteServerSuccess(test_framework *testing.T) {
	deleteSetup()

	returnBody := "{\"result\":\"OK\",\"serverId\":123}"

	// Mocking
	PrepareMockClient(test_framework).
		PerformDelete(URL).
		Return(WithResponse(200, ioutil.NopCloser(strings.NewReader(returnBody))), nil)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteServerNotFound(test_framework *testing.T) {
	deleteSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformDelete(URL).
		Return(WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())

}

func TestDeleteServerError(test_framework *testing.T) {
	deleteSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformDelete(URL).
		Return(WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestDeleteServerClientFailure(test_framework *testing.T) {
	deleteSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformDelete(URL).
		Return(nil, testutil.TestError)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteServerKeycloakFailure(test_framework *testing.T) {
	deleteSetup()

	// Mocking
	PrepareMockClient(test_framework).
		PerformDelete(URL).
		Return(nil, testutil.TestKeycloakError)

	// Run command
	err := delete.DeleteServerCmd.RunE(delete.DeleteServerCmd, []string{SERVERID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
