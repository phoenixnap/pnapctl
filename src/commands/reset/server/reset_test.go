package server

import (
	"encoding/json"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"

	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"

	"github.com/stretchr/testify/assert"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func resetServerSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// Setup
	serverReset := generators.Generate[bmcapisdk.ServerReset]()

	resetResult := generators.Generate[bmcapisdk.ResetResult]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, serverReset)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(RESOURCEID, serverReset).
		Return(&resetResult, nil)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestResetServerSuccessYAML(test_framework *testing.T) {
	resetServerSuccess(test_framework, yaml.Marshal)
}

func TestResetServerSuccessJSON(test_framework *testing.T) {
	resetServerSuccess(test_framework, json.Marshal)
}

func TestResetServerSuccessNoFile(test_framework *testing.T) {
	// Setup
	resetResult := generators.Generate[bmcapisdk.ResetResult]()

	Filename = ""

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(RESOURCEID, bmcapisdk.ServerReset{}).
		Return(&resetResult, nil)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestResetServerFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestResetServerUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestResetServerClientFailure(test_framework *testing.T) {
	// Setup
	serverReset := generators.Generate[bmcapisdk.ServerReset]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, serverReset)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerReset(RESOURCEID, serverReset).
		Return(nil, testutil.TestError)

	// Run command
	err := ResetServerCmd.RunE(ResetServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
