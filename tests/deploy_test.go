package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/tests/generators"

	"phoenixnap.com/pnap-cli/pnapctl/bmc/deploy"

	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"

	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/tests/testutil"

	. "phoenixnap.com/pnap-cli/tests/mockhelp"
)

func deploySetup() {
	URL = "servers"
}
func TestDeployServerSuccessYAML(test_framework *testing.T) {
	deploySetup()

	// Setup
	serverCreate := deploy.ServerCreate{
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

	deploy.Filename = FILENAME

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
	err := deploy.DeployCmd.RunE(deploy.DeployCmd, []string{})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestDeployServerSuccessJSON(test_framework *testing.T) {
	deploySetup()

	// Setup
	server := generators.GenerateServer()

	serverCreate := deploy.ServerCreate{
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

	deploy.Filename = FILENAME

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
	err := deploy.DeployCmd.RunE(deploy.DeployCmd, []string{})

	// Assertions
	testutil.AssertNoError(test_framework, err)
}

func TestDeployServerFileNotFoundFailure(test_framework *testing.T) {
	deploySetup()

	// Setup
	deploy.Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, errors.New(ctlerrors.FileDoesNotExist)).
		Times(1)

	// Run command
	err := deploy.DeployCmd.RunE(deploy.DeployCmd, []string{})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())

}

func TestDeployServerUnmarshallingFailure(test_framework *testing.T) {
	deploySetup()

	// Invalid contents of the file
	// filecontents := make([]byte, 10)
	filecontents := []byte(`sshKeys ["1","2","3","4"]`)

	deploy.Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := deploy.DeployCmd.RunE(deploy.DeployCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericNonRequestError(ctlerrors.UnmarshallingInFileProcessor, "deploy")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestDeployServerFileReadingFailure(test_framework *testing.T) {
	deploySetup()

	// Setup
	deploy.Filename = FILENAME

	// Mocking
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, errors.New(ctlerrors.FileReading)).
		Times(1)

	// Run command
	err := deploy.DeployCmd.RunE(deploy.DeployCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericNonRequestError(ctlerrors.FileReading, "deploy")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestDeployServerBackendErrorFailure(test_framework *testing.T) {
	deploySetup()

	// Setup
	serverCreate := deploy.ServerCreate{
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

	deploy.Filename = FILENAME

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
	err := deploy.DeployCmd.RunE(deploy.DeployCmd, []string{})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}

func TestDeployServerClientFailure(test_framework *testing.T) {
	deploySetup()

	// Setup
	serverCreate := deploy.ServerCreate{
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

	deploy.Filename = FILENAME

	// Mocking
	PrepareMockClient(test_framework).
		PerformPost(URL, bytes.NewBuffer(jsonmarshal)).
		Return(WithResponse(500, nil), testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := deploy.DeployCmd.RunE(deploy.DeployCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError("deploy")

	// Assertions
	testutil.AssertEqual(test_framework, expectedErr.Error(), err.Error())
}
