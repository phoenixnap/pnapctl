package clusters

import (
	"errors"
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/tables"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetServerShortSuccess(test_framework *testing.T) {

	cluster := generators.GenerateCluster()
	var clusterTable interface{}
	clusterTable = tables.ClusterFromSdk(cluster)

	PrepareRancherMockClient(test_framework).
		ClusterGetById(RESOURCEID).
		Return(cluster, WithResponse(200, WithBody(cluster)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(clusterTable, "get clusters").
		Return(nil)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetServerNotFound(test_framework *testing.T) {
	PrepareRancherMockClient(test_framework).
		ClusterGetById(RESOURCEID).
		Return(ranchersdk.Cluster{}, WithResponse(400, nil), nil)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'get clusters' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetServerClientFailure(test_framework *testing.T) {
	PrepareRancherMockClient(test_framework).
		ClusterGetById(RESOURCEID).
		Return(ranchersdk.Cluster{}, nil, testutil.TestError)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get clusters", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetServerKeycloakFailure(test_framework *testing.T) {
	PrepareRancherMockClient(test_framework).
		ClusterGetById(RESOURCEID).
		Return(ranchersdk.Cluster{}, nil, testutil.TestKeycloakError)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetServerPrinterFailure(test_framework *testing.T) {
	cluster := generators.GenerateCluster()
	clusterTable := tables.ClusterFromSdk(cluster)

	PrepareRancherMockClient(test_framework).
		ClusterGetById(RESOURCEID).
		Return(cluster, WithResponse(200, WithBody(cluster)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(clusterTable, "get clusters").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
