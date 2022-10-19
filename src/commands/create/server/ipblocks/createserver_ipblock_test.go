package ipblocks

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreateServerIpBlockSuccessYAML(test_framework *testing.T) {
	serverIpBlockSdk := generators.Generate[bmcapisdk.ServerIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverIpBlockSdk)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockPost(RESOURCEID, gomock.Eq(serverIpBlockSdk)).
		Return(&serverIpBlockSdk, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerIpBlockSuccessJSON(test_framework *testing.T) {
	serverIpBlockSdk := generators.Generate[bmcapisdk.ServerIpBlock]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverIpBlockSdk)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockPost(RESOURCEID, gomock.Eq(serverIpBlockSdk)).
		Return(&serverIpBlockSdk, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerIpBlockFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."})

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestCreateServerIpBlockUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`Name: desc`)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateServerIpBlockFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateServerIpBlockClientFailure(test_framework *testing.T) {
	// Setup
	serverIpBlockSdk := generators.Generate[bmcapisdk.ServerIpBlock]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverIpBlockSdk)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockPost(RESOURCEID, gomock.Eq(serverIpBlockSdk)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
