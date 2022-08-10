package servermodels

import bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"

type IpBlocksConfiguration struct {
	ConfigurationType *string         `yaml:"configurationType,omitempty" json:"configurationType,omitempty"`
	IpBlocks          []ServerIpBlock `yaml:"ipBlocks,omitempty" json:"ipBlocks,omitempty"`
}

func (ipBlocksConfiguration *IpBlocksConfiguration) toSdk() *bmcapisdk.IpBlocksConfiguration {
	if ipBlocksConfiguration == nil {
		return nil
	}

	var ipBlocksConfigurationSdk = bmcapisdk.IpBlocksConfiguration{
		ConfigurationType: ipBlocksConfiguration.ConfigurationType,
		IpBlocks:          mapServerIpBlocksToSdk(ipBlocksConfiguration.IpBlocks),
	}

	return &ipBlocksConfigurationSdk
}

func ipBlocksConfigurationFromSdk(ipBlocksConfiguration *bmcapisdk.IpBlocksConfiguration) *IpBlocksConfiguration {
	if ipBlocksConfiguration == nil {
		return nil
	}

	return &IpBlocksConfiguration{
		ConfigurationType: ipBlocksConfiguration.ConfigurationType,
		IpBlocks:          mapServerIpBlocksToCLI(ipBlocksConfiguration.IpBlocks),
	}
}
