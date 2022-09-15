package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
)

//tests

func TestPublicNetworkConfigurationToSdk(test_framework *testing.T) {
	cliModel := GeneratePublicNetworkConfigurationCli()
	sdkModel := cliModel.toSdk()

	assertEqualPublicNetworkConfiguration(test_framework, cliModel, *sdkModel)
}

func TestEmptyPublicNetworkConfigurationToSdk(test_framework *testing.T) {
	cliModel := &PublicNetworkConfiguration{}
	sdkModel := cliModel.toSdk()

	assertEqualPublicNetworkConfiguration(test_framework, *cliModel, *sdkModel)
}

func TestNilPublicNetworkConfigurationToSdk(test_framework *testing.T) {
	var cliModel *PublicNetworkConfiguration = nil

	assert.Nil(test_framework, cliModel)
}

func TestPublicNetworkConfigurationFromSdk(test_framework *testing.T) {
	sdkModel := GeneratePublicNetworkConfigurationSdk()
	cliModel := publicNetworkConfigurationFromSdk(sdkModel)

	assertEqualPublicNetworkConfiguration(test_framework, *cliModel, *sdkModel)
}

func TestEmptyPublicNetworkConfigurationFromSdk(test_framework *testing.T) {
	sdkModel := &bmcapisdk.PublicNetworkConfiguration{}
	cliModel := publicNetworkConfigurationFromSdk(sdkModel)

	assertEqualPublicNetworkConfiguration(test_framework, *cliModel, *sdkModel)
}

func TestNilPublicNetworkConfigurationFromSdk(test_framework *testing.T) {
	var sdkModel *bmcapisdk.PublicNetworkConfiguration = nil
	cliModel := publicNetworkConfigurationFromSdk(sdkModel)

	assert.Nil(test_framework, cliModel)
}

// assertion functions
func assertEqualPublicNetworkConfiguration(test_framework *testing.T, cliPublicNetworkConfiguration PublicNetworkConfiguration, sdkPublicNetworkConfiguration bmcapisdk.PublicNetworkConfiguration) {

	if testutil.AssertNilEquality(test_framework, "Public Network Configuration's Public Networks", cliPublicNetworkConfiguration.PublicNetworks, sdkPublicNetworkConfiguration.PublicNetworks) {
		assert.Equal(test_framework, len(cliPublicNetworkConfiguration.PublicNetworks), len(sdkPublicNetworkConfiguration.PublicNetworks))

		for i := range cliPublicNetworkConfiguration.PublicNetworks {
			assertEqualServerPublicNetwork(test_framework, (cliPublicNetworkConfiguration.PublicNetworks)[i], (sdkPublicNetworkConfiguration.PublicNetworks)[i])
		}
	}
}