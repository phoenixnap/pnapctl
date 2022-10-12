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
		Return(testutil.AsPointer(generators.Generate[ranchersolutionapi.DeleteResult]()), nil)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteClusterClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteClusterKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareRancherMockClient(test_framework).
		ClusterDelete(RESOURCEID).
		Return(nil, testutil.TestKeycloakError)

	// Run command
	err := DeleteClusterCmd.RunE(DeleteClusterCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
