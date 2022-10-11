package ipblock

import (
	"encoding/json"
	"errors"
	"testing"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
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
		Return(&ipBlock, WithResponse(201, WithBody(ipBlock)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

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
		Return(&ipBlock, WithResponse(201, WithBody(ipBlock)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchIpBlockIdNotFound(test_framework *testing.T) {

	// Setup
	ipBlockPatchCli := generators.Generate[ipapi.IpBlockPatch]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockPatchCli)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdPatch(RESOURCEID, gomock.Eq(ipBlockPatchCli)).
		Return(nil, WithResponse(404, nil), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'patch ip-block' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestPatchIpBlockFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestPatchIpBlockUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`error error`)

	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "patch ip-block", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchIpBlockFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'patch ip-block' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "patch ip-block", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchIpBlockBackendErrorFailure(test_framework *testing.T) {
	// Setup
	ipBlockPatchCli := generators.Generate[ipapi.IpBlockPatch]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockPatchCli)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdPatch(RESOURCEID, gomock.Eq(ipBlockPatchCli)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
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
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "patch ip-block", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestPatchIpBlockKeycloakFailure(test_framework *testing.T) {
	// Setup
	ipBlockPatchCli := generators.Generate[ipapi.IpBlockPatch]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockPatchCli)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdPatch(RESOURCEID, gomock.Eq(ipBlockPatchCli)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
