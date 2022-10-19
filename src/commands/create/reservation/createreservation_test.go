package reservation

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreateReservationSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	reservationCreate := generators.Generate[billingapi.ReservationRequest]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(reservationCreate)

	Filename = FILENAME

	// What the server should return.
	createdReservation := generators.Generate[billingapi.Reservation]()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsPost(gomock.Eq(reservationCreate)).
		Return(&createdReservation, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateReservationSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	reservationCreate := generators.Generate[billingapi.ReservationRequest]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(reservationCreate)

	Filename = FILENAME

	// What the server should return.
	createdReservation := generators.Generate[billingapi.Reservation]()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsPost(gomock.Eq(reservationCreate)).
		Return(&createdReservation, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateReservationFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Expected error
	expectedErr := testutil.TestError

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateReservationUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`reservation? ["maybe"]`)

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateReservationClientFailure(test_framework *testing.T) {
	// What the client should receive.
	reservationCreate := generators.Generate[billingapi.ReservationRequest]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(reservationCreate)

	Filename = FILENAME

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsPost(gomock.Eq(reservationCreate)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
