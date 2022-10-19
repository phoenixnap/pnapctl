package server

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"

	"github.com/stretchr/testify/assert"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestDeprovisionServerSuccessYAML(test_framework *testing.T) {
	// Mocking
	result := "Server Deprovisioned"
	requestBody := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, requestBody)

	PrepareBmcApiMockClient(test_framework).
		ServerDeprovision(RESOURCEID, gomock.Eq(requestBody)).
		Return(result, nil)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeprovisionServerSuccessJSON(test_framework *testing.T) {
	// Mocking
	result := "Server Deprovisioned"
	requestBody := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, requestBody)

	PrepareBmcApiMockClient(test_framework).
		ServerDeprovision(RESOURCEID, gomock.Eq(requestBody)).
		Return(result, nil)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeprovisionServerFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestDeprovisionServerUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestDeprovisionServerClientFailure(test_framework *testing.T) {
	// Setup
	requestBody := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, requestBody)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerDeprovision(RESOURCEID, gomock.Eq(requestBody)).
		Return("", testutil.TestError)

	// Run command
	err := DeprovisionServerCmd.RunE(DeprovisionServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
