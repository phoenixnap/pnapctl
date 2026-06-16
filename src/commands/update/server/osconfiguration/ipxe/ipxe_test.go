package ipxe

import (
	"encoding/json"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func updateServerIpxeSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	serverIpxeUpdate := generators.Generate[bmcapisdk.OsConfigurationIPXE]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, serverIpxeUpdate)

	// What the server should return.
	serverIpxeResponse := generators.Generate[bmcapisdk.OsConfigurationIPXE]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerOsConfigurationIpxePut(RESOURCEID, serverIpxeUpdate).
		Return(&serverIpxeResponse, nil)

	// Run command
	err := PutServerIpxeCmd.RunE(PutServerIpxeCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestUpdateServerIpxeSuccessYAML(test_framework *testing.T) {
	updateServerIpxeSuccess(test_framework, yaml.Marshal)
}

func TestUpdateServerIpxeSuccessJSON(test_framework *testing.T) {
	updateServerIpxeSuccess(test_framework, json.Marshal)
}

func TestUpdateServerIpxeFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := PutServerIpxeCmd.RunE(PutServerIpxeCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestUpdateServerIpxeUnmarshallingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := PutServerIpxeCmd.RunE(PutServerIpxeCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestUpdateServerIpxeClientFailure(test_framework *testing.T) {
	// What the client should receive.
	serverIpxeUpdate := generators.Generate[bmcapisdk.OsConfigurationIPXE]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, serverIpxeUpdate)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerOsConfigurationIpxePut(RESOURCEID, serverIpxeUpdate).
		Return(nil, testutil.TestError)

	// Run command
	err := PutServerIpxeCmd.RunE(PutServerIpxeCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
