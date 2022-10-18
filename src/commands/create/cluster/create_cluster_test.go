package cluster

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreateClusterSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	clusterCreate := generators.GenerateClusterSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(clusterCreate)

	Filename = FILENAME

	// What the server should return.
	createdCluster := generators.GenerateClusterSdk()

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(clusterCreate)).
		Return(&createdCluster, WithResponse(201, WithBody(createdCluster)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateClusterSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	clusterCreate := generators.GenerateClusterSdk()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(clusterCreate)

	Filename = FILENAME

	// What the server should return.
	createdCluster := generators.GenerateClusterSdk()

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(clusterCreate)).
		Return(&createdCluster, WithResponse(201, WithBody(createdCluster)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateClusterFileNotFoundFailure(test_framework *testing.T) {

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateClusterUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`sshKeys ["1","2","3","4"]`)

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "create cluster", err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateClusterBackendErrorFailure(test_framework *testing.T) {
	// What the client should receive.
	clusterCreate := generators.GenerateClusterSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(clusterCreate)

	Filename = FILENAME

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(clusterCreate)).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Expected error
	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateClusterClientFailure(test_framework *testing.T) {
	// What the client should receive.
	clusterCreate := generators.GenerateClusterSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(clusterCreate)

	Filename = FILENAME

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(clusterCreate)).
		Return(nil, nil, testutil.TestError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "create cluster", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateClusterKeycloakFailure(test_framework *testing.T) {
	// What the client should receive.
	clusterCreate := generators.GenerateClusterSdk()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(clusterCreate)

	Filename = FILENAME

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(clusterCreate)).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
