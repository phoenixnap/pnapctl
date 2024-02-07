package disable

import (
	"encoding/json"
	"fmt"
	"github.com/go-faker/faker/v4/pkg/options"
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

func autoRenewReservationDisableSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	autoRenewDisableRequest := generators.Generate[billingapi.ReservationAutoRenewDisableRequest](options.WithFieldsToIgnore("AdditionalProperties"))
	fmt.Printf("AutoRenewDisableReason 1: %v\n", autoRenewDisableRequest)
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, autoRenewDisableRequest)
	fmt.Printf("AutoRenewDisableReason 2: %v\n", autoRenewDisableRequest)

	// What the server should return.
	createdReservation := generators.Generate[billingapi.Reservation](options.WithFieldsToIgnore("AdditionalProperties"))
	fmt.Printf("AutoRenewDisableReason 3: %v\n", autoRenewDisableRequest)
	fmt.Printf("Response generator: %+v\n", createdReservation)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationDisableAutoRenew(RESOURCEID, gomock.Eq(autoRenewDisableRequest)).
		Return(&createdReservation, nil)

	fmt.Printf("AutoRenewDisableReason 4: %v\n", autoRenewDisableRequest)
	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	fmt.Printf("AutoRenewDisableReason 5: %v\n", autoRenewDisableRequest)
	// Assertionsioutil
	assert.NoError(test_framework, err)
}

func TestAutoRenewReservationDisableSuccessYAML(test_framework *testing.T) {
	autoRenewReservationDisableSuccess(test_framework, yaml.Marshal)
}

func TestAutoRenewReservationDisableSuccessJSON(test_framework *testing.T) {
	autoRenewReservationDisableSuccess(test_framework, json.Marshal)
}

func TestAutoRenewReservationDisableFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestAutoRenewReservationDisableUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestAutoRenewReservationDisableClientFailure(test_framework *testing.T) {
	// What the client should receive.
	autoRenewDisableRequest := generators.Generate[billingapi.ReservationAutoRenewDisableRequest]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, autoRenewDisableRequest)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ReservationDisableAutoRenew(RESOURCEID, gomock.Eq(autoRenewDisableRequest)).
		Return(nil, testutil.TestError)

	// Run command
	err := AutoRenewDisableReservationCmd.RunE(AutoRenewDisableReservationCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
