package bmcapimodels

import (
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
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
func (networkconfiguration *NetworkConfiguration) toSdk() *bmcapisdk.NetworkConfiguration {
	if networkconfiguration == nil {
		return nil
	}

	return &bmcapisdk.NetworkConfiguration{
		PrivateNetworkConfiguration: networkconfiguration.PrivateNetworkConfiguration.toSdk(),
	}
}

func (privateNetConf *PrivateNetworkConfiguration) toSdk() *bmcapisdk.PrivateNetworkConfiguration {
	if privateNetConf == nil {
		return nil
	}

	return &bmcapisdk.PrivateNetworkConfiguration{
		GatewayAddress:    privateNetConf.GatewayAddress,
		ConfigurationType: privateNetConf.ConfigurationType,
		PrivateNetworks:   mapServerPrivateNetworksToSdk(privateNetConf.PrivateNetworks),
	}
}

func mapServerPrivateNetworksToSdk(serverPrivateNetworks *[]ServerPrivateNetwork) *[]bmcapisdk.ServerPrivateNetwork {
	if serverPrivateNetworks == nil {
		return nil
	}

	var bmcServerPrivateNetworks []bmcapisdk.ServerPrivateNetwork

	for _, serverPrivateNetwork := range *serverPrivateNetworks {
		bmcServerPrivateNetworks = append(bmcServerPrivateNetworks, serverPrivateNetwork.toSdk())
	}

	return &bmcServerPrivateNetworks
}

func (serverPrivateNetwork ServerPrivateNetwork) toSdk() bmcapisdk.ServerPrivateNetwork {
	var serverPrivateNetworkSdk = bmcapisdk.ServerPrivateNetwork{
		Id:                serverPrivateNetwork.Id,
		Ips:               serverPrivateNetwork.Ips,
		Dhcp:              serverPrivateNetwork.Dhcp,
		StatusDescription: serverPrivateNetwork.StatusDescription,
	}

	return serverPrivateNetworkSdk
}

/* SDK to DTO mapping functions */
func NetworkConfigurationSdkToDto(networkConf *bmcapisdk.NetworkConfiguration) *NetworkConfiguration {
	if networkConf == nil {
		return nil
	}

	return &NetworkConfiguration{
		PrivateNetworkConfiguration: privateNetworkConfigurationSdkToDto(networkConf.PrivateNetworkConfiguration),
	}
}

func privateNetworkConfigurationSdkToDto(privateNetworkConfnfiguration *bmcapisdk.PrivateNetworkConfiguration) *PrivateNetworkConfiguration {
	if privateNetworkConfnfiguration == nil {
		return nil
	}

	return &PrivateNetworkConfiguration{
		GatewayAddress:    privateNetworkConfnfiguration.GatewayAddress,
		ConfigurationType: privateNetworkConfnfiguration.ConfigurationType,
		PrivateNetworks:   privateNetworksSdkToDto(privateNetworkConfnfiguration.PrivateNetworks),
	}
}

func privateNetworksSdkToDto(privateNetworks *[]bmcapisdk.ServerPrivateNetwork) *[]ServerPrivateNetwork {
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
		sdkObj := NetworkConfigurationSdkToDto(networkConfiguration)
		return sdkObj.ToTableString()
	}
}
