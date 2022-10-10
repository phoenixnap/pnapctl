package cluster

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeleteClusterSuccess(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(testutil.AsPointer(generators.Generate[ranchersolutionapi.DeleteResult]()), WithResponse(200, nil), nil)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteClusterNotFound(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(nil, WithResponse(404, nil), nil)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'delete cluster' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestDeleteClusterError(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(nil, WithResponse(500, nil), nil)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	expectedMessage := "Command 'delete cluster' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteClusterClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(nil, nil, testutil.TestError)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete cluster", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteClusterKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(nil, nil, testutil.TestKeycloakError)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
