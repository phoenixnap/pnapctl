package server

import (
	"testing"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
)

// Each test needs to have a name like `TestXXX`
// They also need a parameter of `*testing.T`
func TestPowerOffServerSuccess(test_framework *testing.T) {
	actionResult := generators.Generate[bmcapisdk.ActionResult]()
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOff(RESOURCEID).
		Return(&actionResult, nil)

	// Run command
	err := PowerOffServerCmd.RunE(PowerOffServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPowerOffServerClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerPowerOff(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := PowerOffServerCmd.RunE(PowerOffServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
