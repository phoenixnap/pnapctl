package reservations

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

func TestGetReservationShortSuccess(test_framework *testing.T) {
	reservation := billingmodels.GenerateReservation()
	var shortReservation = tables.ShortReservationTableFromSdk(*reservation)

	PrepareBillingMockClient(test_framework).
		ReservationGetById(RESOURCEID).
		Return(reservation, WithResponse(200, WithBody(*reservation)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortReservation, "get reservations").
		Return(nil)

	Full = false
	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetReservationFullSuccess(test_framework *testing.T) {
	reservation := billingmodels.GenerateReservation()
	var reservationTable = tables.ReservationTableFromSdk(*reservation)

	PrepareBillingMockClient(test_framework).
		ReservationGetById(RESOURCEID).
		Return(reservation, WithResponse(200, WithBody(*reservation)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(reservationTable, "get reservations").
		Return(nil)

	Full = true
	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetReservationNotFound(test_framework *testing.T) {
	PrepareBillingMockClient(test_framework).
		ReservationGetById(RESOURCEID).
		Return(nil, WithResponse(400, nil), nil)

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{RESOURCEID})

	// Expected error
	expectedMessage := "Command 'get reservations' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetReservationClientFailure(test_framework *testing.T) {
	PrepareBillingMockClient(test_framework).
		ReservationGetById(RESOURCEID).
		Return(nil, nil, testutil.TestError)

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get reservations", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetReservationKeycloakFailure(test_framework *testing.T) {
	PrepareBillingMockClient(test_framework).
		ReservationGetById(RESOURCEID).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetReservationPrinterFailure(test_framework *testing.T) {
	reservation := billingmodels.GenerateReservation()
	var shortReservation = tables.ShortReservationTableFromSdk(*reservation)

	PrepareBillingMockClient(test_framework).
		ReservationGetById(RESOURCEID).
		Return(reservation, WithResponse(200, WithBody(*reservation)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortReservation, "get reservations").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	Full = false
	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
