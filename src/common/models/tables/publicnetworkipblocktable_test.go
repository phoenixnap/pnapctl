package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnapctl/common/models/generators"
)

func TestPublicNetworkIpBlockTableFromSdkSuccess(test_framework *testing.T) {
	sdk := generators.GeneratePublicNetworkIpBlockSdk()
	tbl := PublicNetworkIpBlockTableFromSdk(sdk)

	assertPublicNetworkIpBlocksEqual(test_framework, sdk, tbl)
}

func assertPublicNetworkIpBlocksEqual(test_framework *testing.T, sdk networkapi.PublicNetworkIpBlock, tbl PublicNetworkIpBlockTable) {
	assert.Equal(test_framework, sdk.Id, tbl.Id)
}
