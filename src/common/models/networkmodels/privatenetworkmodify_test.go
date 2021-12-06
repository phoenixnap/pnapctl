package networkmodels

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

// tests
func TestPrivateNetworkModifyToSdk(test_framework *testing.T) {
	privateNetworkModify := GeneratePrivateNetworkModifyCli()
	sdkPrivateNetworkModify := privateNetworkModify.ToSdk()

	assertEqualPrivateNetworkModify(test_framework, privateNetworkModify, *sdkPrivateNetworkModify)
}

// assertion functions
func assertEqualPrivateNetworkModify(test_framework *testing.T, p1 PrivateNetworkModify, p2 networksdk.PrivateNetworkModify) {
	assert.Equal(test_framework, p1.Name, p2.Name)
	assert.Equal(test_framework, p1.Description, p2.Description)
	assert.Equal(test_framework, p1.LocationDefault, p2.LocationDefault)
}
