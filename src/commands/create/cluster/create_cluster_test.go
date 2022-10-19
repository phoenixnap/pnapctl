package cluster

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreateClusterSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	clusterCreate := generators.Generate[ranchersolutionapi.Cluster]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(clusterCreate)

	Filename = FILENAME

	// What the server should return.
	createdCluster := generators.Generate[ranchersolutionapi.Cluster]()

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(clusterCreate)).
		Return(&createdCluster, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateClusterSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	clusterCreate := generators.Generate[ranchersolutionapi.Cluster]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(clusterCreate)

	Filename = FILENAME

	// What the server should return.
	createdCluster := generators.Generate[ranchersolutionapi.Cluster]()

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(clusterCreate)).
		Return(&createdCluster, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateClusterFileNotFoundFailure(test_framework *testing.T) {

	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."})

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
		Return(filecontents, nil)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestCreateClusterClientFailure(test_framework *testing.T) {
	// What the client should receive.
	clusterCreate := generators.Generate[ranchersolutionapi.Cluster]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(clusterCreate)

	Filename = FILENAME

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(clusterCreate)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}
