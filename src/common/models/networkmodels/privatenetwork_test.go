package networkmodels

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

// tests
func TestPrivateNetworkFromSdk(test_framework *testing.T) {
	sdkPrivateNetwork := GeneratePrivateNetworkSdk()
	privateNetwork := PrivateNetworkFromSdk(sdkPrivateNetwork)

	assertEqualPrivateNetwork(test_framework, privateNetwork, sdkPrivateNetwork)
}

// assertions
func assertEqualPrivateNetwork(test_framework *testing.T, p1 PrivateNetwork, p2 networksdk.PrivateNetwork) {
	assert.Equal(test_framework, p1.Id, p2.Id)
	assert.Equal(test_framework, p1.Name, p2.Name)
	assert.Equal(test_framework, p1.Description, p2.Description)
	assert.Equal(test_framework, p1.VlanId, p2.VlanId)
	assert.Equal(test_framework, p1.Type, p2.Type)
	assert.Equal(test_framework, p1.Location, p2.Location)
	assert.Equal(test_framework, p1.LocationDefault, p2.LocationDefault)
	assert.Equal(test_framework, p1.Cidr, p2.Cidr)
	assert.Equal(test_framework, p1.CreatedOn, p2.CreatedOn)

	testutil.ForEachPair(p1.Servers, p2.Servers).
		Do(test_framework, assertEqualPrivateNetworkServer)
	testutil.ForEachPair(p2.Memberships, p1.Memberships).
		Do(test_framework, assertNetworkMembershipsEqual)
}
