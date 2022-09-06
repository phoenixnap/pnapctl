package reservation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestAutoRenewEnableReservationSuccess(test_framework *testing.T) {
	// Mocking
	reservation := billingmodels.GenerateReservation()
	PrepareBillingMockClient(test_framework).
		ReservationEnableAutoRenew(RESOURCEID).
		Return(reservation, WithResponse(200, WithBody(reservation)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(tables.ShortReservationTableFromSdk(*reservation), "auto-renew enable reservation").
		Return(nil)

	// Run command
	err := AutoRenewEnableReservationCmd.RunE(AutoRenewEnableReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestAutoRenewEnableReservationClientFailure(test_framework *testing.T) {
	PrepareBillingMockClient(test_framework).
		ReservationEnableAutoRenew(RESOURCEID).
		Return(nil, nil, testutil.TestError)

	// Run command
	err := AutoRenewEnableReservationCmd.RunE(AutoRenewEnableReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "auto-renew enable reservation", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestAutoRenewEnableReservationKeycloakFailure(test_framework *testing.T) {
	PrepareBillingMockClient(test_framework).
		ReservationEnableAutoRenew(RESOURCEID).
		Return(nil, nil, testutil.TestKeycloakError)

	// Run command
	err := AutoRenewEnableReservationCmd.RunE(AutoRenewEnableReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, testutil.TestKeycloakError, err.Error())
}

func TestAutoRenewEnableReservationNotFoundFailure(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationEnableAutoRenew(RESOURCEID).
		Return(nil, WithResponse(404, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := AutoRenewEnableReservationCmd.RunE(AutoRenewEnableReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.GenericBMCError.Message, err.Error())
}

func TestAutoRenewEnableReservationBackendErrorFailure(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationEnableAutoRenew(RESOURCEID).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil)

	// Run command
	err := AutoRenewEnableReservationCmd.RunE(AutoRenewEnableReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}
