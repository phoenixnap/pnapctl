package server

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"sigs.k8s.io/yaml"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestPatchServerSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	serverPatch := generators.Generate[bmcapisdk.ServerPatch]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, serverPatch)

	// What the server should return.
	server := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(serverPatch)).
		Return(&server, nil)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchServerSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	serverPatch := generators.Generate[bmcapisdk.ServerPatch]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, serverPatch)

	// What the server should return.
	server := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(serverPatch)).
		Return(&server, nil)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchServerFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestPatchServerUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestPatchServerClientFailure(test_framework *testing.T) {
	// Setup
	serverPatch := generators.Generate[bmcapisdk.ServerPatch]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, serverPatch)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPatch(RESOURCEID, gomock.Eq(serverPatch)).
		Return(nil, testutil.TestError)

	// Run command
	err := PatchServerCmd.RunE(PatchServerCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
