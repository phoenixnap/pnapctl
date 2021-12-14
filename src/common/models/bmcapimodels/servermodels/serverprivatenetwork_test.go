package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
)

// tests
func TestMapServerPrivateNetworkListToSdk(test_framework *testing.T) {
	cliModels := GenerateServerPrivateNetworkListCli(2)
	sdkModels := mapServerPrivateNetworkListToSdk(&cliModels)

	assertEqualServerPrivateNetworkLists(test_framework, &cliModels, sdkModels)
}

func TestEmptyListMapServerPrivateNetworkListToSdk(test_framework *testing.T) {
	cliModels := GenerateServerPrivateNetworkListCli(0)
	sdkModels := mapServerPrivateNetworkListToSdk(&cliModels)

	assert.Equal(test_framework, len(cliModels), len(*sdkModels))

	for i := range cliModels {
		assertEqualServerPrivateNetwork(test_framework, cliModels[i], (*sdkModels)[i])
	}
}

func TestNilMapServerPrivateNetworkListToSdk(test_framework *testing.T) {
	var cliModels *[]ServerPrivateNetwork = nil
	sdkModels := mapServerPrivateNetworkListToSdk(cliModels)

	assert.Nil(test_framework, sdkModels)
}

func TestServerPrivateNetworkToSdk(test_framework *testing.T) {
	cliModel := GenerateServerPrivateNetworkCli()
	sdkModel := cliModel.toSdk()

	assertEqualServerPrivateNetwork(test_framework, cliModel, sdkModel)
}

func TestEmptyServerPrivateNetworkToSdk(test_framework *testing.T) {
	var cliModel *ServerPrivateNetwork = &ServerPrivateNetwork{}
	sdkModel := cliModel.toSdk()

	assertEqualServerPrivateNetwork(test_framework, *cliModel, sdkModel)
}

func TestPrivateNetworkListFromSdk(test_framework *testing.T) {
	sdkModel := GenerateServerPrivateNetworkListSdk(2)
	cliModel := serverPrivateNetworkListFromSdk(&sdkModel)

	assertEqualServerPrivateNetworkLists(test_framework, cliModel, &sdkModel)
}

func TestEmptyPrivateNetworkListFromSdk(test_framework *testing.T) {
	sdkModel := GenerateServerPrivateNetworkListSdk(0)
	cliModel := serverPrivateNetworkListFromSdk(&sdkModel)

	assertEqualServerPrivateNetworkLists(test_framework, cliModel, &sdkModel)
}

func TestNilPrivateNetworkListFromSdk(test_framework *testing.T) {
	var sdkModel *[]bmcapisdk.ServerPrivateNetwork = nil
	cliModel := serverPrivateNetworkListFromSdk(sdkModel)

	assert.Nil(test_framework, cliModel)
}

// assertion functions
func assertEqualServerPrivateNetworkLists(test_framework *testing.T, cliServerPrivateNetworkList *[]ServerPrivateNetwork, sdkServerPrivateNetworkList *[]bmcapisdk.ServerPrivateNetwork) {

	if testutil.AssertNilEquality(test_framework, "Private Networks List", cliServerPrivateNetworkList, sdkServerPrivateNetworkList) {
		assert.Equal(test_framework, len(*cliServerPrivateNetworkList), len(*sdkServerPrivateNetworkList))

		for i := range *cliServerPrivateNetworkList {
			assertEqualServerPrivateNetwork(test_framework, (*cliServerPrivateNetworkList)[i], (*sdkServerPrivateNetworkList)[i])
		}
	}
}

func assertEqualServerPrivateNetwork(test_framework *testing.T, cliServerPrivateNetwork ServerPrivateNetwork, sdkServerPrivateNetwork bmcapisdk.ServerPrivateNetwork) {
	assert.Equal(test_framework, cliServerPrivateNetwork.Id, sdkServerPrivateNetwork.Id)

	if testutil.AssertNilEquality(test_framework, "Server Private Network's IPs", cliServerPrivateNetwork.Ips, sdkServerPrivateNetwork.Ips) {
		assert.Equal(test_framework, len(*cliServerPrivateNetwork.Ips), len(*sdkServerPrivateNetwork.Ips))

		for i := range *cliServerPrivateNetwork.Ips {
			assert.Equal(test_framework, (*cliServerPrivateNetwork.Ips)[i], (*sdkServerPrivateNetwork.Ips)[i])
		}
	}

	assert.Equal(test_framework, cliServerPrivateNetwork.Dhcp, sdkServerPrivateNetwork.Dhcp)
	assert.Equal(test_framework, cliServerPrivateNetwork.StatusDescription, sdkServerPrivateNetwork.StatusDescription)
}
