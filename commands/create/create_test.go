package create

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnap-cli/tests/generators"

	create "phoenixnap.com/pnap-cli/commands/create/server"

	"phoenixnap.com/pnap-cli/common/ctlerrors"

	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/tests/testutil"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"
)

func createSetup() {
	URL = "servers"
}
func TestCreateServerSuccessYAML(test_framework *testing.T) {
	createSetup()

	// Setup
	serverCreate := create.ServerCreate{
		Name:        "name",
		Description: "description",
		Public:      true,
		Os:          "os",
		TYPE:        "type",
		Location:    "Location",
		SSHKeys:     []string{"CNDI0W92UYC480D", "HDSIODIPS9879D"},
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	// What the server should receive.
	jsonmarshal, _ := json.Marshal(serverCreate)

	create.Filename = FILENAME

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, bytes.NewBuffer(jsonmarshal)).
		Return(WithResponse(200, WithBody(Body)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)
	// Run command
	err := create.CreateServerCmd.RunE(create.CreateServerCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerSuccessJSON(test_framework *testing.T) {
	createSetup()

	// Setup
	server := generators.GenerateServer()

	serverCreate := create.ServerCreate{
		Name:        "name",
		Description: "description",
		Public:      true,
		Os:          "os",
		TYPE:        "type",
		Location:    "Location",
		SSHKeys:     []string{"CNDI0W92UYC480D", "HDSIODIPS9879D"},
	}

	// What will be sent to the server, and the assumed contents of the file.
	jsonmarshal, _ := json.Marshal(serverCreate)

	create.Filename = FILENAME

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, bytes.NewBuffer(jsonmarshal)).
		Return(WithResponse(200, WithBody(server)), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := create.CreateServerCmd.RunE(create.CreateServerCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerFileNotFoundFailure(test_framework *testing.T) {
	createSetup()

	// Setup
	create.Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, errors.New(ctlerrors.FileDoesNotExist)).
		Times(1)

	// Run command
	err := create.CreateServerCmd.RunE(create.CreateServerCmd, []string{})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestCreateServerUnmarshallingFailure(test_framework *testing.T) {
	createSetup()

	// Invalid contents of the file
	// filecontents := make([]byte, 10)
	filecontents := []byte(`sshKeys ["1","2","3","4"]`)

	create.Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := create.CreateServerCmd.RunE(create.CreateServerCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericNonRequestError(ctlerrors.UnmarshallingInFileProcessor, "create server")

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerFileReadingFailure(test_framework *testing.T) {
	createSetup()

	// Setup
	create.Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, errors.New(ctlerrors.FileReading)).
		Times(1)

	// Run command
	err := create.CreateServerCmd.RunE(create.CreateServerCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericNonRequestError(ctlerrors.FileReading, "create server")

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerBackendErrorFailure(test_framework *testing.T) {
	createSetup()

	// Setup
	serverCreate := create.ServerCreate{
		Name:        "name",
		Description: "description",
		Public:      true,
		Os:          "os",
		TYPE:        "type",
		Location:    "Location",
		SSHKeys:     []string{"CNDI0W92UYC480D", "HDSIODIPS9879D"},
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	// What the server should receive.
	jsonmarshal, _ := json.Marshal(serverCreate)

	create.Filename = FILENAME

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, bytes.NewBuffer(jsonmarshal)).
		Return(WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := create.CreateServerCmd.RunE(create.CreateServerCmd, []string{})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerClientFailure(test_framework *testing.T) {
	createSetup()

	// Setup
	serverCreate := create.ServerCreate{
		Name:        "name",
		Description: "description",
		Public:      true,
		Os:          "os",
		TYPE:        "type",
		Location:    "Location",
		SSHKeys:     []string{"CNDI0W92UYC480D", "HDSIODIPS9879D"},
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	// What the server should receive.
	jsonmarshal, _ := json.Marshal(serverCreate)

	create.Filename = FILENAME

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, bytes.NewBuffer(jsonmarshal)).
		Return(nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := create.CreateServerCmd.RunE(create.CreateServerCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create server", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateServerKeycloakFailure(test_framework *testing.T) {
	createSetup()

	// Setup
	serverCreate := create.ServerCreate{
		Name:        "name",
		Description: "description",
		Public:      true,
		Os:          "os",
		TYPE:        "type",
		Location:    "Location",
		SSHKeys:     []string{"CNDI0W92UYC480D", "HDSIODIPS9879D"},
	}

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(serverCreate)

	// What the server should receive.
	jsonmarshal, _ := json.Marshal(serverCreate)

	create.Filename = FILENAME

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, bytes.NewBuffer(jsonmarshal)).
		Return(nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := create.CreateServerCmd.RunE(create.CreateServerCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
