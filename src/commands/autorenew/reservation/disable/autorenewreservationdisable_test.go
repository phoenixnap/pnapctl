package disable

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

func TestAutoRenewReservationDisableSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	autoRenewDisableRequest := generators.GenerateReservationAutoRenewDisableRequestSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(autoRenewDisableRequest)

	Filename = FILENAME

	// What the server should return.
	createdReservation := generators.GenerateReservation()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationDisableAutoRenew(RESOURCEID, gomock.Eq(autoRenewDisableRequest)).
		Return(createdReservation, WithResponse(201, WithBody(createdReservation)), nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestAutoRenewReservationDisableSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	autoRenewDisableRequest := generators.GenerateReservationAutoRenewDisableRequestSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(autoRenewDisableRequest)

	Filename = FILENAME

	// What the server should return.
	createdReservation := generators.GenerateReservation()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationDisableAutoRenew(RESOURCEID, gomock.Eq(autoRenewDisableRequest)).
		Return(createdReservation, WithResponse(201, WithBody(createdReservation)), nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestAutoRenewReservationDisableFileNotFoundFailure(test_framework *testing.T) {
	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestAutoRenewReservationDisableUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`reservation? ["maybe"]`)

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "auto-renew reservation disable", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestAutoRenewReservationDisableBackendErrorFailure(test_framework *testing.T) {
	// What the client should receive.
	autoRenewDisableRequest := generators.GenerateReservationAutoRenewDisableRequestSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(autoRenewDisableRequest)

	Filename = FILENAME

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationDisableAutoRenew(RESOURCEID, gomock.Eq(autoRenewDisableRequest)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestAutoRenewReservationDisableClientFailure(test_framework *testing.T) {
	// What the client should receive.
	autoRenewDisableRequest := generators.GenerateReservationAutoRenewDisableRequestSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(autoRenewDisableRequest)

	Filename = FILENAME

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationDisableAutoRenew(RESOURCEID, gomock.Eq(autoRenewDisableRequest)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "auto-renew disable reservation", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestAutoRenewReservationDisableKeycloakFailure(test_framework *testing.T) {
	// What the client should receive.
	autoRenewDisableRequest := generators.GenerateReservationAutoRenewDisableRequestSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(autoRenewDisableRequest)

	Filename = FILENAME

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationDisableAutoRenew(RESOURCEID, gomock.Eq(autoRenewDisableRequest)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
