package ip_blocks

import (
	"encoding/json"
	"errors"
	"testing"

	"phoenixnap.com/pnapctl/common/models/ipmodels"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreateIpBlockSuccessYAML(test_framework *testing.T) {
	ipBlockCreateCli := ipmodels.GenerateIpBlockCreateSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockCreateCli)

	Filename = FILENAME

	// What the server should return.
	ipBlock := ipmodels.GenerateIpBlockSdk()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlockPost(gomock.Eq(ipBlockCreateCli)).
		Return(&ipBlock, WithResponse(201, WithBody(ipBlock)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateIpBlockSuccessJSON(test_framework *testing.T) {
	ipBlockCreateCli := ipmodels.GenerateIpBlockCreateSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockCreateCli)

	Filename = FILENAME

	// What the server should return.
	ipBlock := ipmodels.GenerateIpBlockSdk()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlockPost(gomock.Eq(ipBlockCreateCli)).
		Return(&ipBlock, WithResponse(201, WithBody(ipBlock)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateIpBlockFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestCreateIpBlockUnmarshallingFailure(test_framework *testing.T) {
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
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create ip-block", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateIpBlockFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'create ip-block' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "create ip-block", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateIpBlockBackendErrorFailure(test_framework *testing.T) {
	// Setup
	ipBlockCreate := ipmodels.GenerateIpBlockCreateSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockCreate)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlockPost(gomock.Eq(ipBlockCreate)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateIpBlockClientFailure(test_framework *testing.T) {
	// Setup
	ipBlockCreate := ipmodels.GenerateIpBlockCreateSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockCreate)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlockPost(gomock.Eq(ipBlockCreate)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create ip-block", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateIpBlockKeycloakFailure(test_framework *testing.T) {
	// Setup
	ipBlockCreate := ipmodels.GenerateIpBlockCreateSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockCreate)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlockPost(gomock.Eq(ipBlockCreate)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
