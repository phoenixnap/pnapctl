package networkmodels

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

// tests
func TestPrivateNetworkCreateToSdk(test_framework *testing.T) {
	privateNetworkCreate := GeneratePrivateNetworkCreateCli()
	sdkPrivateNetworkCreate := privateNetworkCreate.ToSdk()

	assertEqualPrivateNetworkCreate(test_framework, privateNetworkCreate, *sdkPrivateNetworkCreate)
}

// assertion functions
func assertEqualPrivateNetworkCreate(test_framework *testing.T, p1 PrivateNetworkCreate, p2 networksdk.PrivateNetworkCreate) {
	assert.Equal(test_framework, p1.Name, p2.Name)
	assert.Equal(test_framework, p1.Description, p2.Description)
	assert.Equal(test_framework, p1.Location, p2.Location)
	assert.Equal(test_framework, p1.LocationDefault, p2.LocationDefault)
	assert.Equal(test_framework, p1.Cidr, p2.Cidr)
}
