package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpBlocksConfigurationToSdk(test_framework *testing.T) {
	cliModel := GenerateIpBlocksConfigurationCli()
	sdkModel := cliModel.toSdk()

	assertEqualIpBlocksConfiguration(test_framework, cliModel, sdkModel)
}

func TestNilIpBlocksConfigurationToSdk(test_framework *testing.T) {
	var ipBlocksConfiguration *IpBlocksConfiguration = nil
	assert.Nil(test_framework, ipBlocksConfiguration.toSdk())
}

func TestNilIpBlocksConfigurationFromSdk(test_framework *testing.T) {
	assert.Nil(test_framework, ipBlocksConfigurationFromSdk(nil))
}

func TestIpBlocksConfigurationFromSdk(test_framework *testing.T) {
	sdkModel := GenerateIpBlockConfigurationSdk()
	cliModel := ipBlocksConfigurationFromSdk(sdkModel)

	assertEqualIpBlocksConfiguration(test_framework, cliModel, sdkModel)
}

func assertEqualIpBlocksConfiguration(test_framework *testing.T, cliModel *IpBlocksConfiguration, sdkModel *bmcapisdk.IpBlocksConfiguration) {
	assert.Equal(test_framework, *cliModel.ConfigurationType, *sdkModel.ConfigurationType)
	assertServerIpBlockListEquality(test_framework, cliModel.IpBlocks, sdkModel.IpBlocks)
}
