package reservations

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetReservationShortSuccess(test_framework *testing.T) {
	reservation := generators.Generate[billingapi.Reservation]()
	var shortReservation = tables.ShortReservationTableFromSdk(reservation)

	PrepareBillingMockClient(test_framework).
		ReservationGetById(RESOURCEID).
		Return(&reservation, nil)

	ExpectToPrintSuccess(test_framework, shortReservation)

	Full = false
	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetReservationFullSuccess(test_framework *testing.T) {
	reservation := generators.Generate[billingapi.Reservation]()
	var reservationTable = tables.ReservationTableFromSdk(reservation)

	PrepareBillingMockClient(test_framework).
		ReservationGetById(RESOURCEID).
		Return(&reservation, nil)

	ExpectToPrintSuccess(test_framework, reservationTable)

	Full = true
	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetReservationClientFailure(test_framework *testing.T) {
	PrepareBillingMockClient(test_framework).
		ReservationGetById(RESOURCEID).
		Return(nil, testutil.TestError)

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetReservationPrinterFailure(test_framework *testing.T) {
	reservation := generators.Generate[billingapi.Reservation]()
	var shortReservation = tables.ShortReservationTableFromSdk(reservation)

	PrepareBillingMockClient(test_framework).
		ReservationGetById(RESOURCEID).
		Return(&reservation, nil)

	expectedErr := ExpectToPrintFailure(test_framework, shortReservation)

	Full = false
	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
