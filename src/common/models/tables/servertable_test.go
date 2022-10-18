package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestToShortServerTable(test_framework *testing.T) {
	server := generators.GenerateServerSdk()
	table := ToShortServerTable(server)

	assertShortServersEqual(test_framework, server, table)
}

func TestToLongServerTable(test_framework *testing.T) {
	server := generators.GenerateServerSdk()
	table := ToLongServerTable(server)

	assertLongServersEqual(test_framework, server, table)
}

func TestToServerPrivateNetworkTable(test_framework *testing.T) {
	network := generators.GenerateServerPrivateNetworkSdk()
	table := ToServerPrivateNetworkTable(*network)

	assertServerPrivateNetworksEqual(test_framework, *network, table)
}

func TestToServerIpBlockTable(test_framework *testing.T) {
	sdkModel := generators.GenerateServerIpBlockSdk()
	table := ToServerIpBlockTable(sdkModel)

	assertServerIpBlockEqual(test_framework, sdkModel, table)
}

func assertShortServersEqual(test_framework *testing.T, server bmcapisdk.Server, shortServerTable ShortServerTable) {
	assert.Equal(test_framework, server.Id, shortServerTable.ID)
	assert.Equal(test_framework, server.Status, shortServerTable.Status)
	assert.Equal(test_framework, server.Hostname, shortServerTable.Name)
	assert.Equal(test_framework, DerefString(server.Description), shortServerTable.Description)
	assert.Equal(test_framework, server.PrivateIpAddresses, shortServerTable.PrivateIPAddresses)
	assert.Equal(test_framework, server.PublicIpAddresses, shortServerTable.PublicIPAddresses)
}

func assertLongServersEqual(test_framework *testing.T, server bmcapisdk.Server, longServerTable LongServerTable) {
	assert.Equal(test_framework, server.Id, longServerTable.Id)
	assert.Equal(test_framework, server.Status, longServerTable.Status)
	assert.Equal(test_framework, server.Hostname, longServerTable.Hostname)
	assert.Equal(test_framework, DerefString(server.Description), longServerTable.Description)
	assert.Equal(test_framework, server.Os, longServerTable.Os)
	assert.Equal(test_framework, server.Type, longServerTable.Type)
	assert.Equal(test_framework, server.Location, longServerTable.Location)
	assert.Equal(test_framework, server.Cpu, longServerTable.Cpu)
	assert.Equal(test_framework, server.CpuCount, longServerTable.CpuCount)
	assert.Equal(test_framework, server.CoresPerCpu, longServerTable.CoresPerCpu)
	assert.Equal(test_framework, server.CpuFrequency, longServerTable.CpuFrequency)
	assert.Equal(test_framework, server.Ram, longServerTable.Ram)
	assert.Equal(test_framework, server.Storage, longServerTable.Storage)
	assert.Equal(test_framework, server.PrivateIpAddresses, longServerTable.PrivateIpAddresses)
	assert.Equal(test_framework, server.PublicIpAddresses, longServerTable.PublicIpAddresses)
	assert.Equal(test_framework, DerefString(server.ReservationId), longServerTable.ReservationId)
	assert.Equal(test_framework, server.PricingModel, longServerTable.PricingModel)
	assert.Equal(test_framework, DerefString(server.Password), longServerTable.Password)
	assert.Equal(test_framework, DerefString(server.NetworkType), longServerTable.NetworkType)
	assert.Equal(test_framework, DerefString(server.ClusterId), longServerTable.ClusterId)
	assert.Equal(test_framework, iterutils.MapRef(server.Tags, models.TagsToTableString), longServerTable.Tags)
	assert.Equal(test_framework, DerefTimeAsString(server.ProvisionedOn), longServerTable.ProvisionedOn)
	assert.Equal(test_framework, models.OsConfigurationToTableString(server.OsConfiguration), longServerTable.OsConfiguration)
	assert.Equal(test_framework, models.NetworkConfigurationToTableString(&server.NetworkConfiguration), longServerTable.NetworkConfiguration)
}

func assertServerPrivateNetworksEqual(test_framework *testing.T, privateNetwork bmcapisdk.ServerPrivateNetwork, table ServerPrivateNetworkTable) {
	Dhcp := false
	if privateNetwork.Dhcp != nil {
		Dhcp = *privateNetwork.Dhcp
	}

	assert.Equal(test_framework, privateNetwork.Id, table.Id)
	assert.Equal(test_framework, privateNetwork.Ips, table.Ips)
	assert.Equal(test_framework, Dhcp, table.Dhcp)
	assert.Equal(test_framework, DerefString(privateNetwork.StatusDescription), table.StatusDescription)
}

func assertServerIpBlockEqual(test_framework *testing.T, serverIpBlock bmcapisdk.ServerIpBlock, table ServerIpBlockTable) {
	assert.Equal(test_framework, serverIpBlock.Id, table.Id)
	assert.Equal(test_framework, *serverIpBlock.VlanId, table.VlanId)
}
