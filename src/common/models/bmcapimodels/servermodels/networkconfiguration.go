package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

type NetworkConfiguration struct {
	PrivateNetworkConfiguration *PrivateNetworkConfiguration `yaml:"privateNetworkConfiguration,omitempty" json:"privateNetworkConfiguration,omitempty"`
	IpBlocksConfiguration       *IpBlocksConfiguration       `yaml:"ipBlocksConfiguration,omitempty" json:"ipBlocksConfiguration,omitempty"`
}

func (networkconfiguration *NetworkConfiguration) toSdk() *bmcapisdk.NetworkConfiguration {
	if networkconfiguration == nil {
		return nil
	}

	return &bmcapisdk.NetworkConfiguration{
		PrivateNetworkConfiguration: networkconfiguration.PrivateNetworkConfiguration.toSdk(),
		IpBlocksConfiguration:       networkconfiguration.IpBlocksConfiguration.toSdk(),
	}
}

func NetworkConfigurationFromSdk(networkConf *bmcapisdk.NetworkConfiguration) *NetworkConfiguration {
	if networkConf == nil {
		return nil
	}

	return &NetworkConfiguration{
		PrivateNetworkConfiguration: privateNetworkConfigurationFromSdk(networkConf.PrivateNetworkConfiguration),
		IpBlocksConfiguration:       ipBlocksConfigurationFromSdk(networkConf.IpBlocksConfiguration),
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
