package publicnetwork

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestPatchPublicNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkModifyCli := generators.Generate[networkapi.PublicNetworkModify]()
	publicNetworkModifySdk := publicNetworkModifyCli

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(publicNetworkModifyCli)

	Filename = FILENAME

	// What the server should return.
	publicNetwork := generators.Generate[networkapi.PublicNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkPatch(RESOURCEID, gomock.Eq(publicNetworkModifySdk)).
		Return(&publicNetwork, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchPublicNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkModifyCli := generators.Generate[networkapi.PublicNetworkModify]()
	publicNetworkModifySdk := publicNetworkModifyCli

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(publicNetworkModifySdk)

	Filename = FILENAME

	// What the server should return.
	publicNetwork := generators.Generate[networkapi.PublicNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkPatch(RESOURCEID, gomock.Eq(publicNetworkModifySdk)).
		Return(&publicNetwork, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchPublicNetworkFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."})

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchPublicNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`Invalid`)

	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchPublicNetworkFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchPublicNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkModifyCli := generators.Generate[networkapi.PublicNetworkModify]()
	publicNetworkModifySdk := publicNetworkModifyCli

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(publicNetworkModifySdk)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkPatch(RESOURCEID, gomock.Eq(publicNetworkModifySdk)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := PatchPublicNetworkCmd.RunE(PatchPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
