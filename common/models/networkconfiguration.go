package models

import (
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
)

type NetworkConfiguration struct {
	PrivateNetworkConfiguration *PrivateNetworkConfiguration `yaml:"privateNetworkConfiguration" json:"privateNetworkConfiguration"`
}

type PrivateNetworkConfiguration struct {
	GatewayAddress    *string                 `yaml:"gatewayAddress" json:"gatewayAddress"`
	ConfigurationType *string                 `yaml:"configurationType" json:"configurationType"`
	PrivateNetworks   *[]ServerPrivateNetwork `yaml:"privateNetworks" json:"privateNetworks"`
}

type ServerPrivateNetwork struct {
	Id                string    `yaml:"id" json:"id"`
	Ips               *[]string `yaml:"ips" json:"ips"`
	Dhcp              *bool     `yaml:"dhcp" json:"dhcp"`
	StatusDescription *string   `yaml:"statusDescription" json:"statusDescription"`
}

/* DTO to SDK mapping functions*/
func (networkconfiguration *NetworkConfiguration) toSdk() *bmcapi.NetworkConfiguration {
	if networkconfiguration == nil {
		return nil
	}

	return &bmcapi.NetworkConfiguration{
		PrivateNetworkConfiguration: networkconfiguration.PrivateNetworkConfiguration.toSdk(),
	}
}

func (privateNetConf *PrivateNetworkConfiguration) toSdk() *bmcapi.PrivateNetworkConfiguration {
	if privateNetConf == nil {
		return nil
	}

	return &bmcapi.PrivateNetworkConfiguration{
		GatewayAddress:    privateNetConf.GatewayAddress,
		ConfigurationType: privateNetConf.ConfigurationType,
		PrivateNetworks:   mapServerPrivateNetworksToSdk(privateNetConf.PrivateNetworks),
	}
}

func mapServerPrivateNetworksToSdk(serverPrivateNetworks *[]ServerPrivateNetwork) *[]bmcapi.ServerPrivateNetwork {
	if serverPrivateNetworks == nil {
		return nil
	}

	var bmcServerPrivateNetworks []bmcapi.ServerPrivateNetwork

	for _, serverPrivateNetwork := range *serverPrivateNetworks {
		bmcServerPrivateNetworks = append(bmcServerPrivateNetworks, serverPrivateNetwork.toSdk())
	}

	return &bmcServerPrivateNetworks
}

func (serverPrivateNetwork ServerPrivateNetwork) toSdk() bmcapi.ServerPrivateNetwork {
	var serverPrivateNetworkSdk = bmcapi.ServerPrivateNetwork{
		Id:                serverPrivateNetwork.Id,
		Ips:               serverPrivateNetwork.Ips,
		Dhcp:              serverPrivateNetwork.Dhcp,
		StatusDescription: serverPrivateNetwork.StatusDescription,
	}

	return serverPrivateNetworkSdk
}

/* SDK to DTO mapping functions */
func networkConfigurationSdkToDto(networkConf *bmcapi.NetworkConfiguration) *NetworkConfiguration {
	if networkConf == nil {
		return nil
	}

	return &NetworkConfiguration{
		PrivateNetworkConfiguration: privateNetworkConfigurationSdkToDto(networkConf.PrivateNetworkConfiguration),
	}
}

func privateNetworkConfigurationSdkToDto(privateNetworkConfnfiguration *bmcapi.PrivateNetworkConfiguration) *PrivateNetworkConfiguration {
	if privateNetworkConfnfiguration == nil {
		return nil
	}

	return &PrivateNetworkConfiguration{
		GatewayAddress:    privateNetworkConfnfiguration.GatewayAddress,
		ConfigurationType: privateNetworkConfnfiguration.ConfigurationType,
		PrivateNetworks:   privateNetworksSdkToDto(privateNetworkConfnfiguration.PrivateNetworks),
	}
}

func privateNetworksSdkToDto(privateNetworks *[]bmcapi.ServerPrivateNetwork) *[]ServerPrivateNetwork {
	if privateNetworks == nil {
		return nil
	}

	var bmcServerPrivateNetworks []ServerPrivateNetwork

	for _, bmcServerPrivateNetwork := range *privateNetworks {
		bmcServerPrivateNetworks = append(bmcServerPrivateNetworks, ServerPrivateNetwork{
			Id:                bmcServerPrivateNetwork.Id,
			Ips:               bmcServerPrivateNetwork.Ips,
			Dhcp:              bmcServerPrivateNetwork.Dhcp,
			StatusDescription: bmcServerPrivateNetwork.StatusDescription,
		})
	}

	return &bmcServerPrivateNetworks
}
