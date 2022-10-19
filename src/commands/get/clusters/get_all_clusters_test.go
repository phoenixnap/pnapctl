package clusters

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllClustersShortSuccess(test_framework *testing.T) {
	clusters := testutil.GenN(2, generators.Generate[ranchersolutionapi.Cluster])
	clusterlist := iterutils.Map(clusters, tables.ClusterFromSdk)

	// Mocking
	PrepareRancherMockClient(test_framework).
		ClustersGet().
		Return(clusters, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(clusterlist).
		Return(nil)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllClustersClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClustersGet().
		Return(nil, testutil.TestError)

	err := GetClustersCmd.RunE(GetClustersCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetAllClustersPrinterFailure(test_framework *testing.T) {
	clusters := testutil.GenN(2, generators.Generate[ranchersolutionapi.Cluster])
	clusterlist := iterutils.Map(clusters, tables.ClusterFromSdk)

	PrepareRancherMockClient(test_framework).
		ClustersGet().
		Return(clusters, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(clusterlist).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetClustersCmd.RunE(GetClustersCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
