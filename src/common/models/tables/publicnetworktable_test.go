package tables

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestPublicNetworkTableFromSdkSuccess(test_framework *testing.T) {
	sdk := generators.Generate[networkapi.PublicNetwork]()
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

	sdkMemberships := iterutils.MapRef(sdk.Memberships, models.NetworkMembershipToTableString)
	sdkIpBlocks := iterutils.MapRef(sdk.IpBlocks, models.PublicNetworkIpBlockToTableString)

	assert.Equal(test_framework, sdkMemberships, tbl.Memberships)
	assert.Equal(test_framework, sdkIpBlocks, tbl.IpBlocks)
}
