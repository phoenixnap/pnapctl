package tables

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestPublicNetworkTableFromSdkSuccess(test_framework *testing.T) {
	sdk := networkmodels.GeneratePublicNetworkSdk()
	tbl := PublicNetworkTableFromSdk(sdk)

	assertPublicNetworksEqual(test_framework, sdk, tbl)
}

func assertPublicNetworksEqual(test_framework *testing.T, sdk networkapi.PublicNetwork, tbl PublicNetworkTable) {
	assert.Equal(test_framework, sdk.Id, tbl.Id)
	assert.Equal(test_framework, sdk.VlanId, tbl.VlanId)
	assert.Equal(test_framework, sdk.Name, tbl.Name)
	assert.Equal(test_framework, sdk.Location, tbl.Location)
	assert.Equal(test_framework, *sdk.Description, tbl.Description)
	assert.Equal(test_framework, sdk.CreatedOn.String(), tbl.CreatedOn)

	sdkMemberships := iterutils.Map(sdk.Memberships, networkmodels.NetworkMembershipToTableStrings)
	sdkIpBlocks := iterutils.Map(sdk.IpBlocks, networkmodels.PublicNetworkIpBlockToTableStrings)

	assert.Equal(test_framework, sdkMemberships, tbl.Memberships)
	assert.Equal(test_framework, sdkIpBlocks, tbl.IpBlocks)
}
