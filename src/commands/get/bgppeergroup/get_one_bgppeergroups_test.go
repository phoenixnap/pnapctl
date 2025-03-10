package bgppeergroup

import (
	"testing"

	networkapisdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetBgpPeerGroupSuccess(test_framework *testing.T) {
	bgpPeerGroup := generators.Generate[networkapisdk.BgpPeerGroup]()
	var bgpPeerGroupTable = tables.BgpPeerGroupFromSdk(bgpPeerGroup)

	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupGetById(RESOURCEID).
		Return(&bgpPeerGroup, nil)

	ExpectToPrintSuccess(test_framework, bgpPeerGroupTable)

	err := GetBgpPeerGroupsCmd.RunE(GetBgpPeerGroupsCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetBgpPeerGroupClientFailure(test_framework *testing.T) {
	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupGetById(RESOURCEID).
		Return(nil, testutil.TestError)

	err := GetBgpPeerGroupsCmd.RunE(GetBgpPeerGroupsCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetBgpPeerGroupPrinterFailure(test_framework *testing.T) {
	bgpPeerGroup := generators.Generate[networkapisdk.BgpPeerGroup]()
	bgpPeerGroupTable := tables.BgpPeerGroupFromSdk(bgpPeerGroup)

	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupGetById(RESOURCEID).
		Return(&bgpPeerGroup, nil)

	expectedErr := ExpectToPrintFailure(test_framework, bgpPeerGroupTable)

	err := GetBgpPeerGroupsCmd.RunE(GetBgpPeerGroupsCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
