package server

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func reserveServerSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	serverReserve := generators.Generate[bmcapisdk.ServerReserve]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, serverReserve)

	// What the server should return.
	server := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReserve(RESOURCEID, gomock.Eq(serverReserve)).
		Return(&server, nil)

	// Run command
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestReserveServerSuccessYAML(test_framework *testing.T) {
	reserveServerSuccess(test_framework, yaml.Marshal)
}

func TestReserveServerSuccessJSON(test_framework *testing.T) {
	reserveServerSuccess(test_framework, json.Marshal)
}

func TestReserveServerFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestReserveServerUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestReserveServerClientFailure(test_framework *testing.T) {
	// Setup
	serverReserve := generators.Generate[bmcapisdk.ServerReserve]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, serverReserve)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReserve(RESOURCEID, gomock.Eq(serverReserve)).
		Return(nil, testutil.TestError)

	// Run command
	err := ReserveServerCmd.RunE(ReserveServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
