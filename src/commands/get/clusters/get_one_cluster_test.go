package clusters

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetServerShortSuccess(test_framework *testing.T) {

	cluster := generators.Generate[ranchersolutionapi.Cluster]()
	var clusterTable = tables.ClusterFromSdk(cluster)

	PrepareRancherMockClient(test_framework).
		ClusterGetById(RESOURCEID).
		Return(&cluster, WithResponse(200, WithBody(cluster)), nil)

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
		Return(nil, WithResponse(400, nil), nil)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'get clusters' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetServerClientFailure(test_framework *testing.T) {
	PrepareRancherMockClient(test_framework).
		ClusterGetById(RESOURCEID).
		Return(nil, nil, testutil.TestError)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get clusters", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetServerKeycloakFailure(test_framework *testing.T) {
	PrepareRancherMockClient(test_framework).
		ClusterGetById(RESOURCEID).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetServerPrinterFailure(test_framework *testing.T) {
	cluster := generators.Generate[ranchersolutionapi.Cluster]()
	clusterTable := tables.ClusterFromSdk(cluster)

	PrepareRancherMockClient(test_framework).
		ClusterGetById(RESOURCEID).
		Return(&cluster, WithResponse(200, WithBody(cluster)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(clusterTable, "get clusters").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
