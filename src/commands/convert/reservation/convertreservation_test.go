package reservation

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func convertReservationSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	reservationConvert := generators.Generate[billingapi.ReservationRequest]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, reservationConvert)

	// What the server should return.
	createdReservation := generators.Generate[billingapi.Reservation]()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationConvert(RESOURCEID, gomock.Eq(reservationConvert)).
		Return(&createdReservation, nil)

	// Run command
	err := ConvertReservationCmd.RunE(ConvertReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestConvertReservationSuccessYAML(test_framework *testing.T) {
	convertReservationSuccess(test_framework, yaml.Marshal)
}

func TestConvertReservationSuccessJSON(test_framework *testing.T) {
	convertReservationSuccess(test_framework, json.Marshal)
}

func TestConvertReservationFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := ConvertReservationCmd.RunE(ConvertReservationCmd, []string{RESOURCEID})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestConvertReservationUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := ConvertReservationCmd.RunE(ConvertReservationCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestConvertReservationClientFailure(test_framework *testing.T) {
	// What the client should receive.
	reservationConvert := generators.Generate[billingapi.ReservationRequest]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, reservationConvert)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationConvert(RESOURCEID, gomock.Eq(reservationConvert)).
		Return(nil, testutil.TestError)

	// Run command
	err := ConvertReservationCmd.RunE(ConvertReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
