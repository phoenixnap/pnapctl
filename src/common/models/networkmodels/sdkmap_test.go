package networkmodels

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

// tests
func TestPrivateNetworkFromSdk(test_framework *testing.T) {
	sdkPrivateNetwork := GeneratePrivateNetworkSdk()
	privateNetwork := PrivateNetworkFromSdk(sdkPrivateNetwork)

	assertEqualPrivateNetwork(test_framework, privateNetwork, sdkPrivateNetwork)
}

func TestPrivateNetworkCreateToSdk(test_framework *testing.T) {
	privateNetworkCreate := GeneratePrivateNetworkCreateCli()
	sdkPrivateNetworkCreate := privateNetworkCreate.ToSdk()

	assertEqualPrivateNetworkCreate(test_framework, privateNetworkCreate, *sdkPrivateNetworkCreate)
}

func TestPrivateNetworkModifyToSdk(test_framework *testing.T) {
	privateNetworkModify := GeneratePrivateNetworkModifyCli()
	sdkPrivateNetworkModify := privateNetworkModify.ToSdk()

	assertEqualPrivateNetworkModify(test_framework, privateNetworkModify, *sdkPrivateNetworkModify)
}

func TestPrivateNetworkServerFromSdk(test_framework *testing.T) {
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

	assert.Equal(test_framework, len(p1.Servers), len(p2.Servers))

	for i := range p1.Servers {
		assertEqualPrivateNetworkServer(test_framework, p1.Servers[i], p2.Servers[i])
	}
}

func assertEqualPrivateNetworkCreate(test_framework *testing.T, p1 PrivateNetworkCreate, p2 networksdk.PrivateNetworkCreate) {
	assert.Equal(test_framework, p1.Name, p2.Name)
	assert.Equal(test_framework, p1.Description, p2.Description)
	assert.Equal(test_framework, p1.Location, p2.Location)
	assert.Equal(test_framework, p1.LocationDefault, p2.LocationDefault)
	assert.Equal(test_framework, p1.Cidr, p2.Cidr)
}

func assertEqualPrivateNetworkModify(test_framework *testing.T, p1 PrivateNetworkModify, p2 networksdk.PrivateNetworkModify) {
	assert.Equal(test_framework, p1.Name, p2.Name)
	assert.Equal(test_framework, p1.Description, p2.Description)
	assert.Equal(test_framework, p1.LocationDefault, p2.LocationDefault)
}

func assertEqualPrivateNetworkServer(test_framework *testing.T, p1 PrivateNetworkServer, p2 networksdk.PrivateNetworkServer) {
	assert.Equal(test_framework, p1.Id, p2.Id)
	assert.Equal(test_framework, len(p1.Ips), len(p2.Ips))

	for i := range p1.Ips {
		assert.Equal(test_framework, p1.Ips[i], p2.Ips[i])
	}
}
