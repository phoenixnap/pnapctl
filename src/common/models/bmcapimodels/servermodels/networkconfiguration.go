package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

type NetworkConfiguration struct {
	GatewayAddress              *string                      `yaml:"gatewayAddress,omitempty" json:"gatewayAddress,omitempty"`
	PrivateNetworkConfiguration *PrivateNetworkConfiguration `yaml:"privateNetworkConfiguration,omitempty" json:"privateNetworkConfiguration,omitempty"`
	IpBlocksConfiguration       *IpBlocksConfiguration       `yaml:"ipBlocksConfiguration,omitempty" json:"ipBlocksConfiguration,omitempty"`
	PublicNetworkConfiguration  *PublicNetworkConfiguration  `yaml:"publicNetworkConfiguration,omitempty" json:"publicNetworkConfiguration,omitempty"`
}

func (networkconfiguration *NetworkConfiguration) toSdk() *bmcapisdk.NetworkConfiguration {
	if networkconfiguration == nil {
		return nil
	}

	return &bmcapisdk.NetworkConfiguration{
		GatewayAddress:              networkconfiguration.GatewayAddress,
		PrivateNetworkConfiguration: networkconfiguration.PrivateNetworkConfiguration.toSdk(),
		IpBlocksConfiguration:       networkconfiguration.IpBlocksConfiguration.toSdk(),
		PublicNetworkConfiguration:  networkconfiguration.PublicNetworkConfiguration.toSdk(),
	}
}

func NetworkConfigurationFromSdk(networkConf *bmcapisdk.NetworkConfiguration) *NetworkConfiguration {
	if networkConf == nil {
		return nil
	}

	return &NetworkConfiguration{
		GatewayAddress:              networkConf.GatewayAddress,
		PrivateNetworkConfiguration: privateNetworkConfigurationFromSdk(networkConf.PrivateNetworkConfiguration),
		IpBlocksConfiguration:       ipBlocksConfigurationFromSdk(networkConf.IpBlocksConfiguration),
		PublicNetworkConfiguration:  publicNetworkConfigurationFromSdk(networkConf.PublicNetworkConfiguration),
	}
}

func NetworkConfigurationToTableString(networkConfiguration *bmcapisdk.NetworkConfiguration) string {
	if networkConfiguration == nil {
		return ""
	} else if networkConfiguration.PrivateNetworkConfiguration == nil {
		return "Public"
	} else {
		return "Private"
	}
}
