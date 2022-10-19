package ipblock

import (
	"encoding/json"
	"testing"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"github.com/stretchr/testify/assert"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"sigs.k8s.io/yaml"
)

func TestPatchIpBlockSuccessYAML(test_framework *testing.T) {
	ipBlockPatchCli := generators.Generate[ipapi.IpBlockPatch]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockPatchCli)

	Filename = FILENAME

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdPatch(RESOURCEID, gomock.Eq(ipBlockPatchCli)).
		Return(&ipBlock, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchIpBlockSuccessJSON(test_framework *testing.T) {
	ipBlockPatchCli := generators.Generate[ipapi.IpBlockPatch]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockPatchCli)

	Filename = FILENAME

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdPatch(RESOURCEID, gomock.Eq(ipBlockPatchCli)).
		Return(&ipBlock, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchIpBlockFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."})

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestPatchIpBlockUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`error error`)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchIpBlockFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchIpBlockClientFailure(test_framework *testing.T) {
	// Setup
	ipBlockPatchCli := generators.Generate[ipapi.IpBlockPatch]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockPatchCli)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdPatch(RESOURCEID, gomock.Eq(ipBlockPatchCli)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
