package reservation

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/billingapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func createReservationSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	reservationCreate := generators.Generate[billingapi.ReservationRequest]()
	reservationCreate.AdditionalProperties = map[string]interface{}{}

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, reservationCreate)

	// What the server should return.
	createdReservation := generators.Generate[billingapi.Reservation]()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsPost(gomock.Eq(reservationCreate)).
		Return(&createdReservation, nil)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateReservationSuccessYAML(test_framework *testing.T) {
	createReservationSuccess(test_framework, yaml.Marshal)
}

func TestCreateReservationSuccessJSON(test_framework *testing.T) {
	createReservationSuccess(test_framework, json.Marshal)
}

func TestCreateReservationFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateReservationUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateReservationClientFailure(test_framework *testing.T) {
	// What the client should receive.
	reservationCreate := generators.Generate[billingapi.ReservationRequest]()
	reservationCreate.AdditionalProperties = map[string]interface{}{}

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, reservationCreate)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsPost(gomock.Eq(reservationCreate)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
