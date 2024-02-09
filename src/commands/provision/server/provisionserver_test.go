package server

import (
	"encoding/json"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"phoenixnap.com/pnapctl/common/models/generators"
	"sigs.k8s.io/yaml"

	"github.com/stretchr/testify/assert"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func provisionServerSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// Mocking
	requestBody := generators.Generate[bmcapisdk.ServerProvision]()
	result := generators.Generate[bmcapisdk.Server]()

	// Assumed contents of the file
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, requestBody)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerProvision(RESOURCEID, gomock.Eq(requestBody)).
		Return(&result, nil)

	// Run command
	err := ProvisionServerCmd.RunE(ProvisionServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestProvisionServerSuccessYAML(test_framework *testing.T) {
	provisionServerSuccess(test_framework, yaml.Marshal)
}

func TestProvisionServerSuccessJSON(test_framework *testing.T) {
	provisionServerSuccess(test_framework, json.Marshal)
}

func TestProvisionServerSuccessNoFile(test_framework *testing.T) {
	result := generators.Generate[bmcapisdk.Server]()

	Filename = ""

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerProvision(RESOURCEID, bmcapisdk.ServerProvision{}).
		Return(&result, nil)

	// Run command
	err := ProvisionServerCmd.RunE(ProvisionServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestProvisionServerFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := ProvisionServerCmd.RunE(ProvisionServerCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestResetServerUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := ProvisionServerCmd.RunE(ProvisionServerCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestResetServerClientFailure(test_framework *testing.T) {
	// Setup
	requestBody := generators.Generate[bmcapisdk.ServerProvision]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, requestBody)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerProvision(RESOURCEID, requestBody).
		Return(nil, testutil.TestError)

	// Run command
	err := ProvisionServerCmd.RunE(ProvisionServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
