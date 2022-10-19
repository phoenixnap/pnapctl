package server

import (
	"testing"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
)

func TestPowerOnServerSuccess(test_framework *testing.T) {
	actionResult := generators.Generate[bmcapisdk.ActionResult]()
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOn(RESOURCEID).
		Return(&actionResult, nil)

	err := PowerOnServerCmd.RunE(PowerOnServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPowerOnServerClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOn(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := PowerOnServerCmd.RunE(PowerOnServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}
