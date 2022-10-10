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

func TestGetAllClustersShortSuccess(test_framework *testing.T) {
	clusters := testutil.GenN(2, generators.Generate[ranchersolutionapi.Cluster])

	var clusterlist []interface{}

	for _, x := range clusters {
		clusterlist = append(clusterlist, tables.ClusterFromSdk(x))
	}

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClustersGet().
		Return(clusters, WithResponse(200, WithBody(clusters)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(clusterlist, "get clusters").
		Return(nil)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllClustersClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClustersGet().
		Return(nil, WithResponse(200, nil), testutil.TestError)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get Clusters", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllClustersKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClustersGet().
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllClustersPrinterFailure(test_framework *testing.T) {
	clusters := testutil.GenN(2, generators.Generate[ranchersolutionapi.Cluster])

	var clusterlist []interface{}

	for _, x := range clusters {
		clusterlist = append(clusterlist, tables.ClusterFromSdk(x))
	}

	PrepareRancherMockClient(test_framework).
		ClustersGet().
		Return(clusters, WithResponse(200, WithBody(clusters)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(clusterlist, "get clusters").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetClustersCmd.RunE(GetClustersCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
