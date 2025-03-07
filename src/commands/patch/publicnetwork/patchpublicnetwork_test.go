package publicnetwork

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func patchPublicNetworkSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	publicNetworkModifyCli := generators.Generate[networkapi.PublicNetworkModify]()
	publicNetworkModifySdk := publicNetworkModifyCli

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, publicNetworkModifyCli)

	// What the server should return.
	publicNetwork := generators.Generate[networkapi.PublicNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkPatch(RESOURCEID, gomock.Eq(publicNetworkModifySdk)).
		Return(&publicNetwork, nil)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchPublicNetworkSuccessYAML(test_framework *testing.T) {
	patchPublicNetworkSuccess(test_framework, yaml.Marshal)
}

func TestPatchPublicNetworkSuccessJSON(test_framework *testing.T) {
	patchPublicNetworkSuccess(test_framework, json.Marshal)
}

func TestPatchPublicNetworkFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchPublicNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestPatchPublicNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkModifyCli := generators.Generate[networkapi.PublicNetworkModify]()
	publicNetworkModifySdk := publicNetworkModifyCli

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, publicNetworkModifySdk)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkPatch(RESOURCEID, gomock.Eq(publicNetworkModifySdk)).
		Return(nil, testutil.TestError)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
