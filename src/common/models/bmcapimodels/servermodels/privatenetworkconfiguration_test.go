package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
)

// tests

func TestPrivateNetworkConfigurationToSdk(test_framework *testing.T) {
	cliModel := GeneratePrivateNetworkConfigurationCli()
	sdkModel := cliModel.toSdk()

	assertEqualPrivateNetworkConfiguration(test_framework, cliModel, *sdkModel)
}

func TestEmptyPrivateNetworkConfigurationToSdk(test_framework *testing.T) {
	cliModel := &PrivateNetworkConfiguration{}
	sdkModel := cliModel.toSdk()

	assertEqualPrivateNetworkConfiguration(test_framework, *cliModel, *sdkModel)
}

func TestNilPrivateNetworkConfigurationToSdk(test_framework *testing.T) {
	var cliModel *PrivateNetworkConfiguration = nil

	assert.Nil(test_framework, cliModel)
}

func TestPrivateNetworkConfigurationFromSdk(test_framework *testing.T) {
	sdkModel := GeneratePrivateNetworkConfigurationSdk()
	cliModel := privateNetworkConfigurationFromSdk(sdkModel)

	assertEqualPrivateNetworkConfiguration(test_framework, *cliModel, *sdkModel)
}

func TestEmptyPrivateNetworkConfigurationFromSdk(test_framework *testing.T) {
	sdkModel := &bmcapisdk.PrivateNetworkConfiguration{}
	cliModel := privateNetworkConfigurationFromSdk(sdkModel)

	assertEqualPrivateNetworkConfiguration(test_framework, *cliModel, *sdkModel)
}

func TestNilPrivateNetworkConfigurationFromSdk(test_framework *testing.T) {
	var sdkModel *bmcapisdk.PrivateNetworkConfiguration = nil
	cliModel := privateNetworkConfigurationFromSdk(sdkModel)

	assert.Nil(test_framework, cliModel)
}

// assertion functions
func assertEqualPrivateNetworkConfiguration(test_framework *testing.T, cliPrivateNetworkConfiguration PrivateNetworkConfiguration, sdkPrivateNetworkConfiguration bmcapisdk.PrivateNetworkConfiguration) {
	assert.Equal(test_framework, cliPrivateNetworkConfiguration.GatewayAddress, sdkPrivateNetworkConfiguration.GatewayAddress)
	assert.Equal(test_framework, cliPrivateNetworkConfiguration.ConfigurationType, sdkPrivateNetworkConfiguration.ConfigurationType)

	if testutil.AssertNilEquality(test_framework, "Private Network Configuration's Private Networks", cliPrivateNetworkConfiguration.PrivateNetworks, sdkPrivateNetworkConfiguration.PrivateNetworks) {
		assert.Equal(test_framework, len(cliPrivateNetworkConfiguration.PrivateNetworks), len(sdkPrivateNetworkConfiguration.PrivateNetworks))

		for i := range cliPrivateNetworkConfiguration.PrivateNetworks {
			assertEqualServerPrivateNetwork(test_framework, (cliPrivateNetworkConfiguration.PrivateNetworks)[i], (sdkPrivateNetworkConfiguration.PrivateNetworks)[i])
		}
	}
}
