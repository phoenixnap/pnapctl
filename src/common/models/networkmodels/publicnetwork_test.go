package networkmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestPublicNetworkFromSdkSuccess(test_framework *testing.T) {
	sdk := GeneratePublicNetworkSdk()
	cli := PublicNetworkFromSdk(sdk)

	assertPublicNetworksEqual(test_framework, sdk, cli)
}

func TestPublicNetworksIpBlockToSdkSuccess(test_framework *testing.T) {
	cli := GeneratePublicNetworkIpBlockCli()
	sdk := cli.ToSdk()

	assertPublicNetworksIpBlockEqual(test_framework, *sdk, cli)
}

func assertPublicNetworksEqual(test_framework *testing.T, sdk networkapi.PublicNetwork, cli PublicNetwork) {
	assert.Equal(test_framework, sdk.Id, cli.Id)
	assert.Equal(test_framework, sdk.VlanId, cli.VlanId)
	assert.Equal(test_framework, sdk.Name, cli.Name)
	assert.Equal(test_framework, sdk.Location, cli.Location)
	assert.Equal(test_framework, sdk.Description, cli.Description)
	assert.Equal(test_framework, sdk.CreatedOn, cli.CreatedOn)

	testutil.ForEachPair(sdk.Memberships, cli.Memberships).
		Do(test_framework, assertNetworkMembershipsEqual)
	testutil.ForEachPair(sdk.IpBlocks, cli.IpBlocks).
		Do(test_framework, assertPublicNetworksIpBlockEqual)
}

func assertNetworkMembershipsEqual(test_framework *testing.T, sdk networkapi.NetworkMembership, cli NetworkMembership) {
	assert.Equal(test_framework, sdk.ResourceId, cli.ResourceId)
	assert.Equal(test_framework, sdk.ResourceType, cli.ResourceType)

	testutil.ForEachPair(sdk.Ips, cli.Ips).
		Do(test_framework, testutil.AssertEqual[string])
}

func assertPublicNetworksIpBlockEqual(test_framework *testing.T, sdk networkapi.PublicNetworkIpBlock, cli PublicNetworkIpBlock) {
	assert.Equal(test_framework, sdk.Id, cli.Id)
}
