package networkmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestPublicNetworkCreateToSdkSuccess(test_framework *testing.T) {
	cli := GeneratePublicNetworkCreateCli()
	sdk := cli.ToSdk()

	assertPublicNetworkCreateEqual(test_framework, *sdk, cli)
}

func assertPublicNetworkCreateEqual(test_framework *testing.T, sdk networkapi.PublicNetworkCreate, cli PublicNetworkCreate) {
	assert.Equal(test_framework, sdk.Name, cli.Name)
	assert.Equal(test_framework, sdk.Description, cli.Description)
	assert.Equal(test_framework, sdk.Location, cli.Location)

	testutil.ForEachPair(sdk.IpBlocks, cli.IpBlocks).
		Do(test_framework, assertPublicNetworksIpBlockEqual)
}
