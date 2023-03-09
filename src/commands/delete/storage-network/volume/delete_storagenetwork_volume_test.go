package storagenetwork

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeleteStorageNetworkVolumeSuccess(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageDeleteVolume(RESOURCEID, RESOURCEID).
		Return(nil)

	// Run command
	err := DeleteStorageNetworkVolumeCmd.RunE(DeleteStorageNetworkVolumeCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteStorageNetworkVolumeClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageDeleteVolume(RESOURCEID, RESOURCEID).
		Return(testutil.TestError)

	// Run command
	err := DeleteStorageNetworkVolumeCmd.RunE(DeleteStorageNetworkVolumeCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
