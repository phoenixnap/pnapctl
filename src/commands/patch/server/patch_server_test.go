package server

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"sigs.k8s.io/yaml"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestPatchServerSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	serverPatch := generators.Generate[bmcapisdk.ServerPatch]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverPatch)

	Filename = FILENAME

	// What the server should return.
	server := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(serverPatch)).
		Return(&server, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchServerSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	serverPatch := generators.Generate[bmcapisdk.ServerPatch]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPatch)

	Filename = FILENAME

	// What the server should return.
	server := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(serverPatch)).
		Return(&server, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchServerFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	ExpectFromFileFailure(test_framework)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := testutil.TestError

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestPatchServerUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	// filecontents := make([]byte, 10)
	filecontents := []byte(`notproperty: desc`)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestPatchServerFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchServerClientFailure(test_framework *testing.T) {
	// Setup
	serverPatch := generators.Generate[bmcapisdk.ServerPatch]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverPatch)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(serverPatch)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
