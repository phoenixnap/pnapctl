package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"
)

func TestToShortServerTable(test_framework *testing.T) {
	server := servermodels.GenerateServerSdk()
	table := ToShortServerTable(server)

	assertShortServersEqual(test_framework, server, table)
}

func TestToLongServerTable(test_framework *testing.T) {
	server := servermodels.GenerateServerSdk()
	table := ToLongServerTable(server)

	assertLongServersEqual(test_framework, server, table)
}

func TestToServerPrivateNetworkTable(test_framework *testing.T) {
	network := servermodels.GenerateServerPrivateNetworkSdk()
	table := ToServerPrivateNetworkTable(network)

	assertServerPrivateNetworksEqual(test_framework, network, table)
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
	assert.Equal(test_framework, servermodels.TagsToTableStrings(server.Tags), longServerTable.Tags)
	assert.Equal(test_framework, DerefTimeAsString(server.ProvisionedOn), longServerTable.ProvisionedOn)
	assert.Equal(test_framework, servermodels.OsConfigurationToTableString(server.OsConfiguration), longServerTable.OsConfiguration)
	assert.Equal(test_framework, servermodels.NetworkConfigurationToTableString(&server.NetworkConfiguration), longServerTable.NetworkConfiguration)
}

func assertServerPrivateNetworksEqual(test_framework *testing.T, privateNetwork bmcapisdk.ServerPrivateNetwork, table ServerPrivateNetworkTable) {
	Dhcp := false
	if privateNetwork.Dhcp != nil {
		Dhcp = *privateNetwork.Dhcp
	}

	assert.Equal(test_framework, privateNetwork.Id, table.Id)
	assert.Equal(test_framework, DerefStringList(privateNetwork.Ips), table.Ips)
	assert.Equal(test_framework, Dhcp, table.Dhcp)
	assert.Equal(test_framework, DerefString(privateNetwork.StatusDescription), table.StatusDescription)
}
