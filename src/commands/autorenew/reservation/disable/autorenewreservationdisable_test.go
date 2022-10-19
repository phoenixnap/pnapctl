package disable

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

func TestAutoRenewReservationDisableSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	autoRenewDisableRequest := generators.Generate[billingapi.ReservationAutoRenewDisableRequest]()
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, autoRenewDisableRequest)

	// What the server should return.
	createdReservation := generators.Generate[billingapi.Reservation]()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationDisableAutoRenew(RESOURCEID, gomock.Eq(autoRenewDisableRequest)).
		Return(&createdReservation, nil)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestAutoRenewReservationDisableSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	autoRenewDisableRequest := generators.Generate[billingapi.ReservationAutoRenewDisableRequest]()
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, autoRenewDisableRequest)

	// What the server should return.
	createdReservation := generators.Generate[billingapi.Reservation]()

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationDisableAutoRenew(RESOURCEID, gomock.Eq(autoRenewDisableRequest)).
		Return(&createdReservation, nil)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestAutoRenewReservationDisableFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	ExpectFromFileFailure(test_framework)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := testutil.TestError

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestAutoRenewReservationDisableUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`reservation? ["maybe"]`)

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestAutoRenewReservationDisableClientFailure(test_framework *testing.T) {
	// What the client should receive.
	autoRenewDisableRequest := generators.Generate[billingapi.ReservationAutoRenewDisableRequest]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(autoRenewDisableRequest)

	Filename = FILENAME

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationDisableAutoRenew(RESOURCEID, gomock.Eq(autoRenewDisableRequest)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
