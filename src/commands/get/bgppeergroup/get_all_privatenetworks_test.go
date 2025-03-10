package bgppeergroup

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllBgpPeerGroupsShortSuccess(test_framework *testing.T) {
	bgpPeerGroups := testutil.GenN(5, generators.Generate[networkapi.BgpPeerGroup])
	bgpPeerGroupList := iterutils.MapInterface(bgpPeerGroups, tables.BgpPeerGroupFromSdk)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupsGet("").
		Return(bgpPeerGroups, nil)

	ExpectToPrintSuccess(test_framework, bgpPeerGroupList)

	err := GetBgpPeerGroupsCmd.RunE(GetBgpPeerGroupsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllBgpPeerGroupsClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupsGet("").
		Return(nil, testutil.TestError)

	err := GetBgpPeerGroupsCmd.RunE(GetBgpPeerGroupsCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetAllBgpPeerGroupsPrinterFailure(test_framework *testing.T) {
	bgpPeerGroups := testutil.GenN(5, generators.Generate[networkapi.BgpPeerGroup])
	bgpPeerGroupList := iterutils.MapInterface(bgpPeerGroups, tables.BgpPeerGroupFromSdk)

	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupsGet("").
		Return(bgpPeerGroups, nil)

	expectedErr := ExpectToPrintFailure(test_framework, bgpPeerGroupList)

	err := GetBgpPeerGroupsCmd.RunE(GetBgpPeerGroupsCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
