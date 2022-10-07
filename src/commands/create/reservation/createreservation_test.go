package reservation

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreateReservationSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	reservationCreate := generators.GenerateReservationRequestSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(reservationCreate)

	Filename = FILENAME

	// What the server should return.
	createdReservation := generators.GenerateReservation()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsPost(gomock.Eq(reservationCreate)).
		Return(&createdReservation, WithResponse(201, WithBody(createdReservation)), nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateReservationSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	reservationCreate := generators.GenerateReservationRequestSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(reservationCreate)

	Filename = FILENAME

	// What the server should return.
	createdReservation := generators.GenerateReservation()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsPost(gomock.Eq(reservationCreate)).
		Return(&createdReservation, WithResponse(201, WithBody(createdReservation)), nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateReservationFileNotFoundFailure(test_framework *testing.T) {
	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateReservationUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`reservation? ["maybe"]`)

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create reservation", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateReservationBackendErrorFailure(test_framework *testing.T) {
	// What the client should receive.
	reservationCreate := generators.GenerateReservationRequestSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(reservationCreate)

	Filename = FILENAME

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsPost(gomock.Eq(reservationCreate)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateReservationClientFailure(test_framework *testing.T) {
	// What the client should receive.
	reservationCreate := generators.GenerateReservationRequestSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(reservationCreate)

	Filename = FILENAME

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsPost(gomock.Eq(reservationCreate)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create reservation", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateReservationKeycloakFailure(test_framework *testing.T) {
	// What the client should receive.
	reservationCreate := generators.GenerateReservationRequestSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(reservationCreate)

	Filename = FILENAME

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationsPost(gomock.Eq(reservationCreate)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateReservationCmd.RunE(CreateReservationCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
