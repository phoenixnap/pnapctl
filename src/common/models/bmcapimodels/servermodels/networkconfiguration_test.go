package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tests
func TestNetworkConfigurationToSdk(test_framework *testing.T) {
	cliNetworkConfiguration := GenerateNetworkConfigurationCli()
	sdkModel := cliNetworkConfiguration.toSdk()

	assertEqualNetworkconfiguration(test_framework, cliNetworkConfiguration, *sdkModel)
}

func TestNilNetworkConfigurationToSdk(test_framework *testing.T) {
	var test *NetworkConfiguration = nil
	sdkModel := test.toSdk()

	assert.Nil(test_framework, sdkModel)
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
