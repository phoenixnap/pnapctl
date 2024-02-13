package reservations

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func getQueryParams() string {
	return productCategory
}

func TestGetAllReservationsShortSuccess(test_framework *testing.T) {
	reservationList := testutil.GenN(5, generators.Generate[billingapi.Reservation])

	shortReservations := iterutils.MapInterface(
		reservationList,
		tables.ShortReservationTableFromSdk,
	)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsGet(getQueryParams()).
		Return(reservationList, nil)

	ExpectToPrintSuccess(test_framework, shortReservations)

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllReservationsFullSuccess(test_framework *testing.T) {
	reservationList := testutil.GenN(5, generators.Generate[billingapi.Reservation])

	reservationTables := iterutils.MapInterface(
		reservationList,
		tables.ReservationTableFromSdk,
	)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsGet(getQueryParams()).
		Return(reservationList, nil)

	ExpectToPrintSuccess(test_framework, reservationTables)

	// to display full output
	Full = true

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllReservationsClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsGet(getQueryParams()).
		Return(nil, testutil.TestError)

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetAllReservationsPrinterFailure(test_framework *testing.T) {
	reservationList := testutil.GenN(5, generators.Generate[billingapi.Reservation])

	shortReservations := iterutils.MapInterface(
		reservationList,
		tables.ShortReservationTableFromSdk,
	)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsGet(getQueryParams()).
		Return(reservationList, nil)

	expectedErr := ExpectToPrintFailure(test_framework, shortReservations)

	Full = false

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
