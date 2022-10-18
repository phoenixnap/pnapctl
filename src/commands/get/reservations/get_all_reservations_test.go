package reservations

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/queryparams/billing"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllReservationsShortSuccess(test_framework *testing.T) {
	reservationList := testutil.GenNDeref(5, generators.GenerateReservation)
	queryParams := generators.GenerateReservationGetQueryParams()
	setQueryParams(*queryParams)

	shortReservations := iterutils.MapInterface(
		reservationList,
		tables.ShortReservationTableFromSdk,
	)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsGet(*queryParams).
		Return(reservationList, WithResponse(200, WithBody(reservationList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortReservations, "get reservations").
		Return(nil)

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllReservationsFullSuccess(test_framework *testing.T) {
	reservationList := testutil.GenNDeref(5, generators.GenerateReservation)
	queryParams := generators.GenerateReservationGetQueryParams()
	setQueryParams(*queryParams)

	reservationTables := iterutils.MapInterface(
		reservationList,
		tables.ReservationTableFromSdk,
	)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsGet(*queryParams).
		Return(reservationList, WithResponse(200, WithBody(reservationList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(reservationTables, "get reservations").
		Return(nil)

	// to display full output
	Full = true

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllReservationsClientFailure(test_framework *testing.T) {
	queryParams := generators.GenerateReservationGetQueryParams()
	setQueryParams(*queryParams)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsGet(*queryParams).
		Return(nil, WithResponse(400, nil), testutil.TestError)

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get reservations", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllReservationsKeycloakFailure(test_framework *testing.T) {
	queryParams := generators.GenerateReservationGetQueryParams()
	setQueryParams(*queryParams)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsGet(*queryParams).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllReservationsPrinterFailure(test_framework *testing.T) {
	reservationList := testutil.GenNDeref(5, generators.GenerateReservation)
	queryParams := generators.GenerateReservationGetQueryParams()
	setQueryParams(*queryParams)

	shortReservations := iterutils.MapInterface(
		reservationList,
		tables.ShortReservationTableFromSdk,
	)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsGet(*queryParams).
		Return(reservationList, WithResponse(200, WithBody(reservationList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortReservations, "get reservations").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	Full = false

	err := GetReservationsCmd.RunE(GetReservationsCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}

func setQueryParams(queryparams billing.ReservationsGetQueryParams) {
	productCategory = string(*queryparams.ProductCategory)
}
