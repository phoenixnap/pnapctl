package server

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"

	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestCreateServerSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	serverCreate := generators.Generate[bmcapisdk.ServerCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	Filename = FILENAME

	// What the server should return.
	createdServer := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(serverCreate)).
		Return(&createdServer, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	serverCreate := generators.Generate[bmcapisdk.ServerCreate]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverCreate)

	Filename = FILENAME

	// What the server should return.
	createdServer := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(serverCreate)).
		Return(&createdServer, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Expected command
	expectedErr := testutil.TestError // TODO remove this from tests. We should give plain text here, not compare it.

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestCreateServerUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`sshKeys ["1","2","3","4"]`)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateServerFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateServerClientFailure(test_framework *testing.T) {
	// Setup
	serverCreate := generators.Generate[bmcapisdk.ServerCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(serverCreate)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
