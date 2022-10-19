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

func TestRebootServerSuccess(test_framework *testing.T) {
	// Mocking
	actionResult := generators.Generate[bmcapisdk.ActionResult]()
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(RESOURCEID).
		Return(&actionResult, nil)

	// Run command
	err := RebootCmd.RunE(RebootCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestRebootServerClientFail(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerReboot(RESOURCEID).
		Return(nil, testutil.TestError)

	err := RebootCmd.RunE(RebootCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}
