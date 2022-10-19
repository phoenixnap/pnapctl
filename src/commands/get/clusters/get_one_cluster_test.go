package clusters

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
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
		Return(&cluster, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(clusterTable).
		Return(nil)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetServerClientFailure(test_framework *testing.T) {
	PrepareRancherMockClient(test_framework).
		ClusterGetById(RESOURCEID).
		Return(nil, testutil.TestError)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetServerPrinterFailure(test_framework *testing.T) {
	cluster := generators.Generate[ranchersolutionapi.Cluster]()
	clusterTable := tables.ClusterFromSdk(cluster)

	PrepareRancherMockClient(test_framework).
		ClusterGetById(RESOURCEID).
		Return(&cluster, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(clusterTable).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetClustersCmd.RunE(GetClustersCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
