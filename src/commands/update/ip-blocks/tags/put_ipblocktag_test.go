package ip_blocks

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

func TestPutIpBlockTagSuccessYAML(test_framework *testing.T) {
	ipBlockPutTagCli := testutil.GenN(3, generators.Generate[ipapi.TagAssignmentRequest])

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockPutTagCli)

	Filename = FILENAME

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(&ipBlock, WithResponse(201, WithBody(ipBlock)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPutIpBlockTagSuccessJSON(test_framework *testing.T) {
	ipBlockPutTagCli := testutil.GenN(3, generators.Generate[ipapi.TagAssignmentRequest])

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockPutTagCli)

	Filename = FILENAME

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(&ipBlock, WithResponse(201, WithBody(ipBlock)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestIpBlockPutTagIdNotFound(test_framework *testing.T) {

	// Setup
	ipBlockPutTagCli := testutil.GenN(3, generators.Generate[ipapi.TagAssignmentRequest])

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockPutTagCli)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(nil, WithResponse(404, nil), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'update ip-block tag' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestIpBlockPutTagFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestIpBlockPutTagUnmarshallingFailure(test_framework *testing.T) {
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
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "update ip-block tag", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestIpBlockPutTagFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'update ip-block tag' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "update ip-block tag", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestIpBlockPutTagBackendErrorFailure(test_framework *testing.T) {
	// Setup
	ipBlockPutTagCli := testutil.GenN(3, generators.Generate[ipapi.TagAssignmentRequest])

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockPutTagCli)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestIpBlockPutTagClientFailure(test_framework *testing.T) {
	// Setup
	ipBlockPutTagCli := testutil.GenN(3, generators.Generate[ipapi.TagAssignmentRequest])

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockPutTagCli)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "update ip-block tag", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestIpBlockPutTagKeycloakFailure(test_framework *testing.T) {
	// Setup
	ipBlockPutTagCli := testutil.GenN(3, generators.Generate[ipapi.TagAssignmentRequest])

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockPutTagCli)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
