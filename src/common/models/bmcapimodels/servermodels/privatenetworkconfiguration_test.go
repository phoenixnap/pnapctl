package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tests

// TODO

// assertion functions
func assertEqualPrivateNetworkConfiguration(test_framework *testing.T, cliPrivateNetworkConfiguration PrivateNetworkConfiguration, sdkPrivateNetworkConfiguration bmcapisdk.PrivateNetworkConfiguration) {
	assert.Equal(test_framework, cliPrivateNetworkConfiguration.GatewayAddress, sdkPrivateNetworkConfiguration.GatewayAddress)
	assert.Equal(test_framework, cliPrivateNetworkConfiguration.ConfigurationType, sdkPrivateNetworkConfiguration.ConfigurationType)

	if cliPrivateNetworkConfiguration.PrivateNetworks == nil {
		assert.Nil(test_framework, sdkPrivateNetworkConfiguration.PrivateNetworks, "CLI Private Network Configuration's Private Networks are nil, but not SDK Private Network Configuration's Private Networks.")
	} else if sdkPrivateNetworkConfiguration.PrivateNetworks == nil {
		assert.Nil(test_framework, cliPrivateNetworkConfiguration.PrivateNetworks, "SDK Private Network Configuration's Private Networks are nil, but not Private Network Configuration's Private Networks.")
	} else {
		assert.Equal(test_framework, len(*cliPrivateNetworkConfiguration.PrivateNetworks), len(*sdkPrivateNetworkConfiguration.PrivateNetworks))

		for i := range *cliPrivateNetworkConfiguration.PrivateNetworks {
			assertEqualServerPrivateNetwork(test_framework, (*cliPrivateNetworkConfiguration.PrivateNetworks)[i], (*sdkPrivateNetworkConfiguration.PrivateNetworks)[i])
		}
	}
}
