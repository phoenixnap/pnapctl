package networkmodels

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
)

// tests
func TestPrivateNetworkServerFromSdk(test_framework *testing.T) {
	sdkPrivateNetwork := GeneratePrivateNetworkSdk()
	privateNetwork := PrivateNetworkFromSdk(sdkPrivateNetwork)

	assertEqualPrivateNetwork(test_framework, privateNetwork, sdkPrivateNetwork)
}

// assertion functions
func assertEqualPrivateNetworkServer(test_framework *testing.T, p1 PrivateNetworkServer, p2 networksdk.PrivateNetworkServer) {
	assert.Equal(test_framework, p1.Id, p2.Id)
	assert.Equal(test_framework, len(p1.Ips), len(p2.Ips))

	for i := range p1.Ips {
		assert.Equal(test_framework, p1.Ips[i], p2.Ips[i])
	}
}
