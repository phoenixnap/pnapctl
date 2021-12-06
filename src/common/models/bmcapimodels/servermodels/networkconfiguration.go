package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

type NetworkConfiguration struct {
	PrivateNetworkConfiguration *PrivateNetworkConfiguration `yaml:"privateNetworkConfiguration" json:"privateNetworkConfiguration"`
}

func (networkconfiguration *NetworkConfiguration) toSdk() *bmcapisdk.NetworkConfiguration {
	if networkconfiguration == nil {
		return nil
	}

	return &bmcapisdk.NetworkConfiguration{
		PrivateNetworkConfiguration: networkconfiguration.PrivateNetworkConfiguration.toSdk(),
	}
}

func NetworkConfigurationFromSdk(networkConf *bmcapisdk.NetworkConfiguration) *NetworkConfiguration {
	if networkConf == nil {
		return nil
	}

	return &NetworkConfiguration{
		PrivateNetworkConfiguration: privateNetworkConfigurationFromSdk(networkConf.PrivateNetworkConfiguration),
	}
}

func (n NetworkConfiguration) ToTableString() string {
	if n.PrivateNetworkConfiguration == nil {
		return "Public"
	} else {
		return "Private"
	}
}

func NetworkConfigurationToTableString(networkConfiguration *bmcapisdk.NetworkConfiguration) string {
	if networkConfiguration == nil {
		return ""
	} else {
		cliObj := NetworkConfigurationFromSdk(networkConfiguration)
		return cliObj.ToTableString()
	}
}
