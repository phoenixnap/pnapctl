package publicnetwork

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

func patchServerPublicNetworkSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive
	serverPublicNetworkPatch := generators.Generate[bmcapisdk.ServerNetworkUpdate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, serverPublicNetworkPatch)

	// What the server should return
	serverPublicNetwork := generators.Generate[bmcapisdk.ServerPublicNetwork]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPatch(RESOURCEID, RESOURCEID, gomock.Eq(serverPublicNetworkPatch)).
		Return(&serverPublicNetwork, nil)

	// Run command
	err := PatchServerPublicNetworkCmd.RunE(PatchServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchServerPublicNetworkSuccessYAML(test_framework *testing.T) {
	patchServerPublicNetworkSuccess(test_framework, yaml.Marshal)
}

func TestPatchServerPublicNetworkSuccessJSON(test_framework *testing.T) {
	patchServerPublicNetworkSuccess(test_framework, json.Marshal)
}

func TestPatchServerPublicNetworkFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := PatchServerPublicNetworkCmd.RunE(PatchServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestPatchServerPublicNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := PatchServerPublicNetworkCmd.RunE(PatchServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestPatchServerClientFailure(test_framework *testing.T) {
	// Setup
	serverPublicNetworkPatch := generators.Generate[bmcapisdk.ServerNetworkUpdate]()

	// Assumed contents of the file
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, serverPublicNetworkPatch)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPatch(RESOURCEID, RESOURCEID, gomock.Eq(serverPublicNetworkPatch)).
		Return(nil, testutil.TestError)

	// Run command
	err := PatchServerPublicNetworkCmd.RunE(PatchServerPublicNetworkCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
