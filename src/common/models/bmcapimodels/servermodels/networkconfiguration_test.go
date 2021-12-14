package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"

	"github.com/stretchr/testify/assert"
)

// tests
func TestNetworkConfigurationToSdk(test_framework *testing.T) {
	cliModel := GenerateNetworkConfigurationCli()
	sdkModel := cliModel.toSdk()

	assertEqualNetworkconfiguration(test_framework, cliModel, *sdkModel)
}

func TestNilNetworkConfigurationPrivateNetworkConfigurationNilToSdk(test_framework *testing.T) {
	cliModel := NetworkConfiguration{
		PrivateNetworkConfiguration: nil,
	}

	sdkModel := cliModel.toSdk()

	assertEqualNetworkconfiguration(test_framework, cliModel, *sdkModel)
}

func TestNilNetworkConfigurationToSdk(test_framework *testing.T) {
	var cliModel *NetworkConfiguration = nil
	sdkModel := cliModel.toSdk()

	assert.Nil(test_framework, sdkModel)
}

func TestNetworkConfigurationFromSdk(test_framework *testing.T) {
	sdkModel := GenerateNetworkConfigurationSdk()
	cliModel := NetworkConfigurationFromSdk(&sdkModel)

	assertEqualNetworkconfiguration(test_framework, *cliModel, sdkModel)
}

func TestNetworkConfigurationPrivateNetworkConfigurationNilFromSdk(test_framework *testing.T) {
	sdkModel := bmcapisdk.NetworkConfiguration{
		PrivateNetworkConfiguration: nil,
	}
	cliModel := NetworkConfigurationFromSdk(&sdkModel)

	assertEqualNetworkconfiguration(test_framework, *cliModel, sdkModel)
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
func assertEqualNetworkconfiguration(test_framework *testing.T, cliNetworkConfiguration NetworkConfiguration, sdkNetworkConfiguration bmcapisdk.NetworkConfiguration) {
	if cliNetworkConfiguration.PrivateNetworkConfiguration == nil {
		assert.Nil(test_framework, sdkNetworkConfiguration.PrivateNetworkConfiguration, "CLI Network Configuration's Private Networks are nil, but not SDK Network Configuration's Private Networks.")
	} else if sdkNetworkConfiguration.PrivateNetworkConfiguration == nil {
		assert.Nil(test_framework, cliNetworkConfiguration.PrivateNetworkConfiguration, "SDK Network Configuration's Private Networks are nil, but not CLI Network Configuration's Private Networks.")
	} else {
		assertEqualPrivateNetworkConfiguration(test_framework, *cliNetworkConfiguration.PrivateNetworkConfiguration, *sdkNetworkConfiguration.PrivateNetworkConfiguration)
	}
}
