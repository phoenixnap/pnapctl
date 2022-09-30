package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
)

// tests
func TestMapServerPublicNetworkListToSdk(test_framework *testing.T) {
	cliModels := GenerateServerPublicNetworkListCli(2)
	sdkModels := mapServerPublicNetworkListToSdk(cliModels)

	assertEqualServerPublicNetworkLists(test_framework, cliModels, sdkModels)
}

func TestEmptyListMapServerPublicNetworkListToSdk(test_framework *testing.T) {
	cliModels := GenerateServerPublicNetworkListCli(0)
	sdkModels := mapServerPublicNetworkListToSdk(cliModels)

	assert.Equal(test_framework, len(cliModels), len(sdkModels))

	for i := range cliModels {
		assertEqualServerPublicNetwork(test_framework, cliModels[i], (sdkModels)[i])
	}
}

func TestNilMapServerPublicNetworkListToSdk(test_framework *testing.T) {
	var cliModels []ServerPublicNetwork = nil
	sdkModels := mapServerPublicNetworkListToSdk(cliModels)

	assert.Nil(test_framework, sdkModels)
}

func TestServerPublicNetworkToSdk(test_framework *testing.T) {
	cliModel := GenerateServerPublicNetworkCli()
	sdkModel := cliModel.toSdk()

	assertEqualServerPublicNetwork(test_framework, cliModel, sdkModel)
}

func TestEmptyServerPublicNetworkToSdk(test_framework *testing.T) {
	var cliModel *ServerPublicNetwork = &ServerPublicNetwork{}
	sdkModel := cliModel.toSdk()

	assertEqualServerPublicNetwork(test_framework, *cliModel, sdkModel)
}

func TestPublicNetworkListFromSdk(test_framework *testing.T) {
	sdkModel := GenerateServerPublicNetworkListSdk(2)
	cliModel := serverPublicNetworkListFromSdk(sdkModel)

	assertEqualServerPublicNetworkLists(test_framework, cliModel, sdkModel)
}

func TestEmptyPublicNetworkListFromSdk(test_framework *testing.T) {
	sdkModel := GenerateServerPublicNetworkListSdk(0)
	cliModel := serverPublicNetworkListFromSdk(sdkModel)

	assertEqualServerPublicNetworkLists(test_framework, cliModel, sdkModel)
}

func TestNilPublicNetworkListFromSdk(test_framework *testing.T) {
	var sdkModel []bmcapisdk.ServerPublicNetwork = nil
	cliModel := serverPublicNetworkListFromSdk(sdkModel)

	assert.Nil(test_framework, cliModel)
}

// assertion functions
func assertEqualServerPublicNetworkLists(test_framework *testing.T, cliServerPublicNetworkList []ServerPublicNetwork, sdkServerPublicNetworkList []bmcapisdk.ServerPublicNetwork) {

	if testutil.AssertNilEquality(test_framework, "Public Networks List", cliServerPublicNetworkList, sdkServerPublicNetworkList) {
		assert.Equal(test_framework, len(cliServerPublicNetworkList), len(sdkServerPublicNetworkList))

		for i := range cliServerPublicNetworkList {
			assertEqualServerPublicNetwork(test_framework, (cliServerPublicNetworkList)[i], (sdkServerPublicNetworkList)[i])
		}
	}
}

func assertEqualServerPublicNetwork(test_framework *testing.T, cliServerPublicNetwork ServerPublicNetwork, sdkServerPublicNetwork bmcapisdk.ServerPublicNetwork) {
	assert.Equal(test_framework, cliServerPublicNetwork.Id, sdkServerPublicNetwork.Id)

	if testutil.AssertNilEquality(test_framework, "Server Public Network's IPs", cliServerPublicNetwork.Ips, sdkServerPublicNetwork.Ips) {
		assert.Equal(test_framework, len(cliServerPublicNetwork.Ips), len(sdkServerPublicNetwork.Ips))

		for i := range cliServerPublicNetwork.Ips {
			assert.Equal(test_framework, (cliServerPublicNetwork.Ips)[i], (sdkServerPublicNetwork.Ips)[i])
		}
	}

	assert.Equal(test_framework, cliServerPublicNetwork.StatusDescription, sdkServerPublicNetwork.StatusDescription)
}
