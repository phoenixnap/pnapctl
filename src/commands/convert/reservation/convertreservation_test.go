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

func TestConvertReservationSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	reservationConvert := generators.Generate[billingapi.ReservationRequest]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(reservationConvert)

	Filename = FILENAME

	// What the server should return.
	createdReservation := generators.Generate[billingapi.Reservation]()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationConvert(RESOURCEID, gomock.Eq(reservationConvert)).
		Return(&createdReservation, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := ConvertReservationCmd.RunE(ConvertReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestConvertReservationSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	reservationConvert := generators.Generate[billingapi.ReservationRequest]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(reservationConvert)

	Filename = FILENAME

	// What the server should return.
	createdReservation := generators.Generate[billingapi.Reservation]()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationConvert(RESOURCEID, gomock.Eq(reservationConvert)).
		Return(&createdReservation, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := ConvertReservationCmd.RunE(ConvertReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestConvertReservationFileNotFoundFailure(test_framework *testing.T) {
	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."})

	// Run command
	err := ConvertReservationCmd.RunE(ConvertReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestConvertReservationUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`reservation? ["maybe"]`)

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := ConvertReservationCmd.RunE(ConvertReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestConvertReservationClientFailure(test_framework *testing.T) {
	// What the client should receive.
	reservationConvert := generators.Generate[billingapi.ReservationRequest]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(reservationConvert)

	Filename = FILENAME

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationConvert(RESOURCEID, gomock.Eq(reservationConvert)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := ConvertReservationCmd.RunE(ConvertReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
