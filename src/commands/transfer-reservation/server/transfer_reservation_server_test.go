package server

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func transferReservationServerSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	serverTransferReservation := generators.Generate[bmcapisdk.ReservationTransferDetails]()

	// Assumed contents of the file
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, serverTransferReservation)

	// What the server should return
	server := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTransferReservation(RESOURCEID, gomock.Eq(serverTransferReservation)).
		Return(&server, nil)

	// Run command
	err := TransferReservationServerCmd.RunE(TransferReservationServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestTransferReservationServerSuccessYAML(test_framework *testing.T) {
	transferReservationServerSuccess(test_framework, yaml.Marshal)
}

func TestTransferReservationServerSuccessJSON(test_framework *testing.T) {
	transferReservationServerSuccess(test_framework, json.Marshal)
}

func TestTransferReservationServerFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := TransferReservationServerCmd.RunE(TransferReservationServerCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestTransferReservationServerUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := TransferReservationServerCmd.RunE(TransferReservationServerCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestTransferReservationServerClientFailure(test_framework *testing.T) {
	// Setup
	serverTransferReservation := generators.Generate[bmcapisdk.ReservationTransferDetails]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, serverTransferReservation)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerTransferReservation(RESOURCEID, gomock.Eq(serverTransferReservation)).
		Return(nil, testutil.TestError)

	// Run command
	err := TransferReservationServerCmd.RunE(TransferReservationServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
