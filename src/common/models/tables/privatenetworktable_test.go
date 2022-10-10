package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
)

func TestPrivateNetworkFromSdk(test_framework *testing.T) {
	privateNetwork := generators.GeneratePrivateNetworkSdk()
	table := PrivateNetworkFromSdk(privateNetwork)

	assertPrivateNetworksEqual(test_framework, privateNetwork, table)
}

func assertPrivateNetworksEqual(test_framework *testing.T, network networksdk.PrivateNetwork, table PrivateNetworkTable) {
	var servers []string

	for _, server := range network.Servers {
		servers = append(servers, models.PrivateNetworkServerToTableString(&server))
	}

	assert.Equal(test_framework, network.Id, table.Id)
	assert.Equal(test_framework, network.Name, table.Name)
	assert.Equal(test_framework, DerefString(network.Description), table.Description)
	assert.Equal(test_framework, network.VlanId, table.VlanId)
	assert.Equal(test_framework, network.Type, table.Type)
	assert.Equal(test_framework, network.Location, table.Location)
	assert.Equal(test_framework, network.LocationDefault, table.LocationDefault)
	assert.Equal(test_framework, network.Cidr, table.Cidr)
	assert.Equal(test_framework, servers, table.Servers)
}
