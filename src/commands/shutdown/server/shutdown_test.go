package server

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestShutdownServerSuccess(test_framework *testing.T) {
	actionResult := generators.Generate[bmcapisdk.ActionResult]()
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(RESOURCEID).
		Return(&actionResult, nil)

	// Run command
	err := ShutdownCmd.RunE(ShutdownCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestShutdownServerClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := ShutdownCmd.RunE(ShutdownCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestShutdownServerKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerShutdown(RESOURCEID).
		Return(nil, testutil.TestKeycloakError)

	// Run command
	err := ShutdownCmd.RunE(ShutdownCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
