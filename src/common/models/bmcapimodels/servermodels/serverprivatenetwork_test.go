package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
)

// tests
func TestMapServerPrivateNetworksToSdk(test_framework *testing.T) {
	cliModels := GenerateServerPrivateNetworkListCli(2)
	sdkModels := mapServerPrivateNetworksToSdk(&cliModels)

	assert.Equal(test_framework, len(cliModels), len(*sdkModels))

	for i := range cliModels {
		assertEqualServerPrivateNetwork(test_framework, cliModels[i], (*sdkModels)[i])
	}
}

func TestEmptyListMapServerPrivateNetworksToSdk(test_framework *testing.T) {
	cliModels := GenerateServerPrivateNetworkListCli(0)
	sdkModels := mapServerPrivateNetworksToSdk(&cliModels)

	assert.Equal(test_framework, len(cliModels), len(*sdkModels))

	for i := range cliModels {
		assertEqualServerPrivateNetwork(test_framework, cliModels[i], (*sdkModels)[i])
	}
}

func TestNilMapServerPrivateNetworksToSdk(test_framework *testing.T) {
	var cliModels *[]ServerPrivateNetwork = nil
	sdkModels := mapServerPrivateNetworksToSdk(cliModels)

	assert.Nil(test_framework, sdkModels)
}

// TODO: Add toSdk tests and the rest

// assertion functions
func assertEqualServerPrivateNetwork(test_framework *testing.T, cliServerPrivateNetwork ServerPrivateNetwork, sdkServerPrivateNetwork bmcapisdk.ServerPrivateNetwork) {
	assert.Equal(test_framework, cliServerPrivateNetwork.Id, sdkServerPrivateNetwork.Id)

	if testutil.AssertNilEquality(test_framework, "Server Private Network's IPs", cliServerPrivateNetwork.Ips, sdkServerPrivateNetwork.Ips) {
		assert.Equal(test_framework, len(*cliServerPrivateNetwork.Ips), len(*sdkServerPrivateNetwork.Ips))

		for i := range *cliServerPrivateNetwork.Ips {
			assert.Equal(test_framework, (*cliServerPrivateNetwork.Ips)[i], (*sdkServerPrivateNetwork.Ips)[i])
		}
	}
	assert.Equal(test_framework, len(*cliServerPrivateNetwork.Ips), len(*sdkServerPrivateNetwork.Ips))

	for i := range *cliServerPrivateNetwork.Ips {
		assert.Equal(test_framework, (*cliServerPrivateNetwork.Ips)[i], (*sdkServerPrivateNetwork.Ips)[i])
	}

	assert.Equal(test_framework, cliServerPrivateNetwork.Dhcp, sdkServerPrivateNetwork.Dhcp)
	assert.Equal(test_framework, cliServerPrivateNetwork.StatusDescription, sdkServerPrivateNetwork.StatusDescription)
}
