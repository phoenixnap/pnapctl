package bgppeergroup

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDeleteBgpPeerGroupSuccess(test_framework *testing.T) {
	// What the server should return.
	deletedBgpPeerGroup := generators.Generate[networkapi.BgpPeerGroup]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupDeleteById(RESOURCEID).
		Return(deletedBgpPeerGroup)

	// Run command
	err := DeleteBgpPeerGroupCmd.RunE(DeleteBgpPeerGroupCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteBgpPeerGroupClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupDeleteById(RESOURCEID).
		Return(testutil.TestError)

	// Run command
	err := DeleteBgpPeerGroupCmd.RunE(DeleteBgpPeerGroupCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
