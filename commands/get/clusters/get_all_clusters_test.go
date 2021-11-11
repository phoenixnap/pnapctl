package clusters

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	ranchersdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/ranchersolutionapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/tables"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetAllServersShortSuccess(test_framework *testing.T) {
	clusters := generators.GenerateClusters(5)

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

func TestGetAllServersClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClustersGet().
		Return([]ranchersdk.Cluster{}, WithResponse(200, nil), testutil.TestError)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get servers", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllServersKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClustersGet().
		Return([]ranchersdk.Cluster{}, nil, testutil.TestKeycloakError)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllServersPrinterFailure(test_framework *testing.T) {
	clusters := generators.GenerateClusters(5)

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
