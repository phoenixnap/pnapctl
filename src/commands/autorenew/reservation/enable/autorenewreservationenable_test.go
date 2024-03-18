package enable

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestAutoRenewReservationEnableSuccess(test_framework *testing.T) {
	// Mocking
	reservation := generators.Generate[billingapi.Reservation]()
	PrepareBillingMockClient(test_framework).
		ReservationEnableAutoRenew(RESOURCEID).
		Return(&reservation, nil)

	ExpectToPrintSuccess(test_framework, tables.ShortReservationTableFromSdk(reservation))

	// Run command
	err := AutoRenewEnableReservationCmd.RunE(AutoRenewEnableReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestAutoRenewReservationEnableClientFailure(test_framework *testing.T) {
	PrepareBillingMockClient(test_framework).
		ReservationEnableAutoRenew(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := AutoRenewEnableReservationCmd.RunE(AutoRenewEnableReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
