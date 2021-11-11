package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ranchersdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/ranchersolutionapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestDeleteClusterSuccess(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(generators.GenerateRancherDeleteResult(), WithResponse(200, nil), nil)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteServerNotFound(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(ranchersdk.DeleteResult{}, WithResponse(404, nil), nil)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'delete cluster' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestDeleteServerError(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(ranchersdk.DeleteResult{}, WithResponse(500, nil), nil)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	expectedMessage := "Command 'delete cluster' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteServerClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(ranchersdk.DeleteResult{}, nil, testutil.TestError)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete cluster", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteServerKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(ranchersdk.DeleteResult{}, nil, testutil.TestKeycloakError)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
