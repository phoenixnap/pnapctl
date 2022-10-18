package server

import (
	"errors"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestRebootServerSuccess(test_framework *testing.T) {
	// Mocking
	actionResult := generators.Generate[bmcapisdk.ActionResult]()
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(RESOURCEID).
		Return(&actionResult, WithResponse(200, WithBody(actionResult)), nil)

	// Run command
	err := RebootCmd.RunE(RebootCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestRebootServerClientFail(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(RESOURCEID).
		Return(nil, nil, testutil.TestError)

	err := RebootCmd.RunE(RebootCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "reboot server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestRebootServerKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(RESOURCEID).
		Return(nil, nil, testutil.TestKeycloakError)

	err := RebootCmd.RunE(RebootCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestRebootServerNotFoundFail(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(RESOURCEID).
		Return(nil, WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	err := RebootCmd.RunE(RebootCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestRebootServerErrorFail(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(RESOURCEID).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	err := RebootCmd.RunE(RebootCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}
