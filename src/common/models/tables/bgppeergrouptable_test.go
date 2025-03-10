package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestBgpPeerGroupFromSdk(test_framework *testing.T) {
	bgpPeerGroup := generators.Generate[networkapi.BgpPeerGroup]()
	table := BgpPeerGroupFromSdk(bgpPeerGroup)

	assertBgpPeerGroupsEqual(test_framework, bgpPeerGroup, table)
}

func assertBgpPeerGroupsEqual(test_framework *testing.T, bgpPeerGroup networkapi.BgpPeerGroup, table BgpPeerGroupTable) {
	assert.Equal(test_framework, bgpPeerGroup.Id, table.Id)
	assert.Equal(test_framework, bgpPeerGroup.Status, table.Id)
	assert.Equal(test_framework, bgpPeerGroup.Location, table.Id)
	assert.Equal(test_framework, iterutils.MapRef(bgpPeerGroup.Ipv4Prefixes, models.BgpIpv4PrefixToTableString), table.Ipv4Prefixes)
	assert.Equal(test_framework, models.AsnDetailsToTableString(&bgpPeerGroup.TargetAsnDetails), table.TargetAsnDetails)
	assert.Equal(test_framework, models.AsnDetailsToTableString(bgpPeerGroup.ActiveAsnDetails), table.ActiveAsnDetails)
	assert.Equal(test_framework, bgpPeerGroup.Password, table.Password)
	assert.Equal(test_framework, bgpPeerGroup.AdvertisedRoutes, table.AdvertisedRoutes)
	assert.Equal(test_framework, bgpPeerGroup.RpkiRoaOriginAsn, table.RpkiRoaOriginAsn)
	assert.Equal(test_framework, bgpPeerGroup.EBgpMultiHop, table.EBgpMultiHop)
	assert.Equal(test_framework, bgpPeerGroup.PeeringLoopbacksV4, table.PeeringLoopbacksV4)
	assert.Equal(test_framework, bgpPeerGroup.KeepAliveTimerSeconds, table.KeepAliveTimerSeconds)
	assert.Equal(test_framework, bgpPeerGroup.HoldTimerSeconds, table.HoldTimerSeconds)
	assert.Equal(test_framework, bgpPeerGroup.CreatedOn, table.CreatedOn)
	assert.Equal(test_framework, bgpPeerGroup.LastUpdatedOn, table.LastUpdatedOn)
}
