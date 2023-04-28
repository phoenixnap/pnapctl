package privatenetwork

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"sigs.k8s.io/yaml"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func getQueryParams() bool {
	return force
}

func patchServerPrivateNetworkSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive
	serverPrivateNetworkPatch := generators.Generate[bmcapisdk.ServerNetworkUpdate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, serverPrivateNetworkPatch)

	// What the server should return
	serverPrivateNetwork := generators.Generate[bmcapisdk.ServerPrivateNetwork]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPrivateNetworkPatch(RESOURCEID, RESOURCEID, gomock.Eq(serverPrivateNetworkPatch), force).
		Return(&serverPrivateNetwork, nil)

	// Run command
	err := PatchServerPrivateNetworkCmd.RunE(PatchServerPrivateNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchServerPrivateNetworkSuccessYAML(test_framework *testing.T) {
	patchServerPrivateNetworkSuccess(test_framework, yaml.Marshal)
}

func TestPatchServerPrivateNetworkSuccessJSON(test_framework *testing.T) {
	patchServerPrivateNetworkSuccess(test_framework, json.Marshal)
}

func TestPatchServerPrivateNetworkFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := PatchServerPrivateNetworkCmd.RunE(PatchServerPrivateNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestPatchServerPrivateNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := PatchServerPrivateNetworkCmd.RunE(PatchServerPrivateNetworkCmd, []string{RESOURCEID, RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestPatchServerClientFailure(test_framework *testing.T) {
	// Setup
	serverPrivateNetworkPatch := generators.Generate[bmcapisdk.ServerNetworkUpdate]()

	// Assumed contents of the file
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, serverPrivateNetworkPatch)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPrivateNetworkPatch(RESOURCEID, RESOURCEID, gomock.Eq(serverPrivateNetworkPatch), force).
		Return(nil, testutil.TestError)

	// Run command
	err := PatchServerPrivateNetworkCmd.RunE(PatchServerPrivateNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
