package servermodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerToShortServer(test_framework *testing.T) {
	serverSdk := GenerateServerSdk()
	shortServer := ToShortServer(serverSdk)

	assert.Equal(test_framework, shortServer.ID, serverSdk.Id)
	assert.Equal(test_framework, shortServer.Status, serverSdk.Status)
	assert.Equal(test_framework, shortServer.Name, serverSdk.Hostname)
	assert.Equal(test_framework, shortServer.Description, serverSdk.Description)
	assert.Equal(test_framework, shortServer.PrivateIPAddresses, serverSdk.PrivateIpAddresses)
	assert.Equal(test_framework, shortServer.PublicIPAddresses, serverSdk.PublicIpAddresses)
}

func TestServerToFullServer(test_framework *testing.T) {
	serverSdk := GenerateServerSdk()
	fullServer := ToFullServer(serverSdk)

	assert.Equal(test_framework, fullServer.Id, serverSdk.Id)
	assert.Equal(test_framework, fullServer.Status, serverSdk.Status)
	assert.Equal(test_framework, fullServer.Hostname, serverSdk.Hostname)
	assert.Equal(test_framework, fullServer.Description, serverSdk.Description)
	assert.Equal(test_framework, fullServer.Os, serverSdk.Os)
	assert.Equal(test_framework, fullServer.Type, serverSdk.Type)
	assert.Equal(test_framework, fullServer.Location, serverSdk.Location)
	assert.Equal(test_framework, fullServer.Cpu, serverSdk.Cpu)
	assert.Equal(test_framework, fullServer.CpuCount, serverSdk.CpuCount)
	assert.Equal(test_framework, fullServer.CoresPerCpu, serverSdk.CoresPerCpu)
	assert.Equal(test_framework, fullServer.CpuFrequency, serverSdk.CpuFrequency)
	assert.Equal(test_framework, fullServer.Ram, serverSdk.Ram)
	assert.Equal(test_framework, fullServer.Storage, serverSdk.Storage)
	assert.Equal(test_framework, fullServer.PrivateIpAddresses, serverSdk.PrivateIpAddresses)
	assert.Equal(test_framework, fullServer.PublicIpAddresses, serverSdk.PublicIpAddresses)
	assert.Equal(test_framework, fullServer.ReservationId, serverSdk.ReservationId)
	assert.Equal(test_framework, fullServer.PricingModel, serverSdk.PricingModel)
	assert.Equal(test_framework, fullServer.Password, serverSdk.Password)
	assert.Equal(test_framework, fullServer.NetworkType, serverSdk.NetworkType)
	assert.Equal(test_framework, fullServer.ClusterId, serverSdk.ClusterId)

	tagAssignments := serverSdk.Tags
	for i, tagAssignment := range fullServer.Tags {
		assertEqualTagAssignment(test_framework, tagAssignment, tagAssignments[i])
	}

	assert.Equal(test_framework, fullServer.ProvisionedOn, serverSdk.ProvisionedOn)
	assertEqualOsConfiguration(test_framework, *fullServer.OsConfiguration, *serverSdk.OsConfiguration)
	assertEqualNetworkConfiguration(test_framework, *fullServer.NetworkConfiguration, serverSdk.NetworkConfiguration)
}
