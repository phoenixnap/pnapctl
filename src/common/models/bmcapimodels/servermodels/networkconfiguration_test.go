package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
)

// tests
func TestNetworkConfigurationToSdk(test_framework *testing.T) {
	cliModel := GenerateNetworkConfigurationCli()
	sdkModel := cliModel.toSdk()

	assertEqualNetworkConfiguration(test_framework, cliModel, *sdkModel)
}

func TestNetworkConfigurationNilPropertiesToSdk(test_framework *testing.T) {
	cliModel := NetworkConfiguration{
		PrivateNetworkConfiguration: nil,
		IpBlocksConfiguration:       nil,
	}

	sdkModel := cliModel.toSdk()

	assertEqualNetworkConfiguration(test_framework, cliModel, *sdkModel)
}

func TestNilNetworkConfigurationToSdk(test_framework *testing.T) {
	var cliModel *NetworkConfiguration = nil
	sdkModel := cliModel.toSdk()

	assert.Nil(test_framework, sdkModel)
}

func TestNetworkConfigurationFromSdk(test_framework *testing.T) {
	sdkModel := GenerateNetworkConfigurationSdk()
	cliModel := NetworkConfigurationFromSdk(&sdkModel)

	assertEqualNetworkConfiguration(test_framework, *cliModel, sdkModel)
}

func TestNetworkConfigurationNilPropertiesFromSdk(test_framework *testing.T) {
	sdkModel := bmcapisdk.NetworkConfiguration{
		PrivateNetworkConfiguration: nil,
		IpBlocksConfiguration:       nil,
	}
	cliModel := NetworkConfigurationFromSdk(&sdkModel)

	assertEqualNetworkConfiguration(test_framework, *cliModel, sdkModel)
}

func TestNilNetworkConfigurationFromSdk(test_framework *testing.T) {
	var sdkModel *bmcapisdk.NetworkConfiguration = nil

	assert.Nil(test_framework, NetworkConfigurationFromSdk(sdkModel))
}

func TestNilNetworkConfigurationToTableString(test_framework *testing.T) {
	var sdkModel *bmcapisdk.NetworkConfiguration = nil
	tableString := NetworkConfigurationToTableString(sdkModel)

	assert.Equal(test_framework, "", tableString)
}

func TestPrivateNetworkConfigurationToTableString(test_framework *testing.T) {
	var sdkModel *bmcapisdk.NetworkConfiguration = &bmcapisdk.NetworkConfiguration{
		PrivateNetworkConfiguration: &bmcapisdk.PrivateNetworkConfiguration{},
	}
	tableString := NetworkConfigurationToTableString(sdkModel)

	assert.Equal(test_framework, "Private", tableString)
}

func TestPublicNetworkConfigurationToTableString(test_framework *testing.T) {
	var sdkModel *bmcapisdk.NetworkConfiguration = &bmcapisdk.NetworkConfiguration{
		PrivateNetworkConfiguration: nil,
	}
	tableString := NetworkConfigurationToTableString(sdkModel)

	assert.Equal(test_framework, "Public", tableString)
}

// assertion functions
func assertEqualNetworkConfiguration(test_framework *testing.T, cliNetworkConfiguration NetworkConfiguration, sdkNetworkConfiguration bmcapisdk.NetworkConfiguration) {
	if testutil.AssertNilEquality(test_framework, "Network Configuration's Private Networks", cliNetworkConfiguration.PrivateNetworkConfiguration, sdkNetworkConfiguration.PrivateNetworkConfiguration) {
		assertEqualPrivateNetworkConfiguration(test_framework, *cliNetworkConfiguration.PrivateNetworkConfiguration, *sdkNetworkConfiguration.PrivateNetworkConfiguration)
		assertEqualIpBlocksConfiguration(test_framework, cliNetworkConfiguration.IpBlocksConfiguration, sdkNetworkConfiguration.IpBlocksConfiguration)
	}
}
