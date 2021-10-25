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

func networkConfigurationDtoToSdk(networkConf *NetworkConfiguration) *bmcapi.NetworkConfiguration {
	if networkConf == nil {
		return nil
	}

	return &bmcapi.NetworkConfiguration{
		PrivateNetworkConfiguration: privateNetworkConfigurationDtoToSdk(networkConf.PrivateNetworkConfiguration),
	}
}

func privateNetworkConfigurationDtoToSdk(privateNetConf *PrivateNetworkConfiguration) *bmcapi.PrivateNetworkConfiguration {
	if privateNetConf == nil {
		return nil
	}

	return &bmcapi.PrivateNetworkConfiguration{
		GatewayAddress:    privateNetConf.GatewayAddress,
		ConfigurationType: privateNetConf.ConfigurationType,
		PrivateNetworks:   privateNetworksDtoToSdk(privateNetConf.PrivateNetworks),
	}
}

func privateNetworksDtoToSdk(privateNetworks *[]ServerPrivateNetwork) *[]bmcapi.ServerPrivateNetwork {
	if privateNetworks == nil {
		return nil
	}

	var bmcPrivNet []bmcapi.ServerPrivateNetwork

	for _, x := range *privateNetworks {
		bmcPrivNet = append(bmcPrivNet, bmcapi.ServerPrivateNetwork{
			Id:                x.Id,
			Ips:               x.Ips,
			Dhcp:              x.Dhcp,
			StatusDescription: x.StatusDescription,
		})
	}

	return &bmcPrivNet
}
