package cluster

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/ranchermodels"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

func TestCreateClusterSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	cluster := ranchermodels.ClusterFromSdk(generators.GenerateCluster())

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(cluster)

	Filename = FILENAME

	// What the server should return.
	createdCluster := generators.GenerateCluster()

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(cluster.ToSdk())).
		Return(createdCluster, WithResponse(200, WithBody(createdCluster)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateClusterSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	cluster := ranchermodels.ClusterFromSdk(generators.GenerateCluster())

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(cluster)

	Filename = FILENAME

	// What the server should return.
	createdCluster := generators.GenerateCluster()

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(cluster.ToSdk())).
		Return(createdCluster, WithResponse(200, WithBody(createdCluster)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
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
		ReadFile(FILENAME).
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
		ReadFile(FILENAME).
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
	cluster := ranchermodels.ClusterFromSdk(generators.GenerateCluster())

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(cluster)

	Filename = FILENAME

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(cluster.ToSdk())).
		Return(ranchersdk.Cluster{}, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
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
	cluster := ranchermodels.ClusterFromSdk(generators.GenerateCluster())

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(cluster)

	Filename = FILENAME

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(cluster.ToSdk())).
		Return(ranchersdk.Cluster{}, nil, testutil.TestError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
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
	cluster := ranchermodels.ClusterFromSdk(generators.GenerateCluster())

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(cluster)

	Filename = FILENAME

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(cluster.ToSdk())).
		Return(ranchersdk.Cluster{}, nil, testutil.TestKeycloakError).
		Times(1)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
