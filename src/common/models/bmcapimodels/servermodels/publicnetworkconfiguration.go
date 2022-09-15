package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

type PublicNetworkConfiguration struct {
	PublicNetworks []ServerPublicNetwork `yaml:"publicNetworks" json:"publicNetworks"`
}

func (publicNetConf *PublicNetworkConfiguration) toSdk() *bmcapisdk.PublicNetworkConfiguration {
	if publicNetConf == nil {
		return nil
	}

	return &bmcapisdk.PublicNetworkConfiguration{
		PublicNetworks: mapServerPublicNetworkListToSdk(publicNetConf.PublicNetworks),
	}
}

func publicNetworkConfigurationFromSdk(publicNetworkConfiguration *bmcapisdk.PublicNetworkConfiguration) *PublicNetworkConfiguration {
	if publicNetworkConfiguration == nil {
		return nil
	}

	return &PublicNetworkConfiguration{
		PublicNetworks: serverPublicNetworkListFromSdk(publicNetworkConfiguration.PublicNetworks),
	}
}
