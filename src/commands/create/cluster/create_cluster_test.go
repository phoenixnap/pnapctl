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
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, clusterCreate)

	// What the server should return.
	createdCluster := generators.Generate[ranchersolutionapi.Cluster]()

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(clusterCreate)).
		Return(&createdCluster, nil)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateClusterSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	clusterCreate := generators.Generate[ranchersolutionapi.Cluster]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, clusterCreate)

	// What the server should return.
	createdCluster := generators.Generate[ranchersolutionapi.Cluster]()

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(clusterCreate)).
		Return(&createdCluster, nil)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateClusterFileProcessorFailure(test_framework *testing.T) {
	Filename = FILENAME

	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateClusterUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateClusterClientFailure(test_framework *testing.T) {
	// What the client should receive.
	clusterCreate := generators.Generate[ranchersolutionapi.Cluster]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, clusterCreate)

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterPost(gomock.Eq(clusterCreate)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateClusterCmd.RunE(CreateClusterCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
