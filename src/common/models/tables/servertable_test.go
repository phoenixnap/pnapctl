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

func assertShortServersEqual(test_framework *testing.T, server bmcapisdk.Server, table ShortServerTable) {
	assert.Equal(test_framework, server.Id, table.ID)
	assert.Equal(test_framework, server.Status, table.Status)
	assert.Equal(test_framework, server.Hostname, table.Name)
	assert.Equal(test_framework, DerefString(server.Description), table.Description)
	assert.Equal(test_framework, server.PrivateIpAddresses, table.PrivateIPAddresses)
	assert.Equal(test_framework, server.PublicIpAddresses, table.PublicIPAddresses)
}

func assertLongServersEqual(test_framework *testing.T, server bmcapisdk.Server, table LongServerTable) {
	assert.Equal(test_framework, server.Id, table.Id)
	assert.Equal(test_framework, server.Status, table.Status)
	assert.Equal(test_framework, server.Hostname, table.Hostname)
	assert.Equal(test_framework, DerefString(server.Description), table.Description)
	assert.Equal(test_framework, server.Os, table.Os)
	assert.Equal(test_framework, server.Type, table.Type)
	assert.Equal(test_framework, server.Location, table.Location)
	assert.Equal(test_framework, server.Cpu, table.Cpu)
	assert.Equal(test_framework, server.CpuCount, table.CpuCount)
	assert.Equal(test_framework, server.CoresPerCpu, table.CoresPerCpu)
	assert.Equal(test_framework, server.CpuFrequency, table.CpuFrequency)
	assert.Equal(test_framework, server.Ram, table.Ram)
	assert.Equal(test_framework, server.Storage, table.Storage)
	assert.Equal(test_framework, server.PrivateIpAddresses, table.PrivateIpAddresses)
	assert.Equal(test_framework, server.PublicIpAddresses, table.PublicIpAddresses)
	assert.Equal(test_framework, DerefString(server.ReservationId), table.ReservationId)
	assert.Equal(test_framework, server.PricingModel, table.PricingModel)
	assert.Equal(test_framework, DerefString(server.Password), table.Password)
	assert.Equal(test_framework, DerefString(server.NetworkType), table.NetworkType)
	assert.Equal(test_framework, DerefString(server.ClusterId), table.ClusterId)
	assert.Equal(test_framework, servermodels.TagsToTableStrings(server.Tags), table.Tags)
	assert.Equal(test_framework, DerefTimeAsString(server.ProvisionedOn), table.ProvisionedOn)
	assert.Equal(test_framework, servermodels.OsConfigurationToTableString(server.OsConfiguration), table.OsConfiguration)
	assert.Equal(test_framework, servermodels.NetworkConfigurationToTableString(&server.NetworkConfiguration), table.NetworkConfiguration)
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
