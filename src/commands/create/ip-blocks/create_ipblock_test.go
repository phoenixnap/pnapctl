package ip_blocks

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreateIpBlockSuccessYAML(test_framework *testing.T) {
	ipBlockCreate := generators.Generate[ipapi.IpBlockCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockCreate)

	Filename = FILENAME

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlockPost(gomock.Eq(ipBlockCreate)).
		Return(&ipBlock, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateIpBlockSuccessJSON(test_framework *testing.T) {
	ipBlockCreateCli := generators.Generate[ipapi.IpBlockCreate]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockCreateCli)

	Filename = FILENAME

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlockPost(gomock.Eq(ipBlockCreateCli)).
		Return(&ipBlock, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateIpBlockFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Expected command
	expectedErr := testutil.TestError

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestCreateIpBlockUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`error error`)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateIpBlockFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateIpBlockClientFailure(test_framework *testing.T) {
	// Setup
	ipBlockCreate := generators.Generate[ipapi.IpBlockCreate]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockCreate)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlockPost(gomock.Eq(ipBlockCreate)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
