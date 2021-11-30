package bmcapimodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
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
	Id                string    `yaml:"id,omitempty" json:"id,omitempty"`
	Ips               *[]string `yaml:"ips,omitempty" json:"ips,omitempty"`
	Dhcp              *bool     `yaml:"dhcp,omitempty" json:"dhcp,omitempty"`
	StatusDescription *string   `yaml:"statusDescription,omitempty" json:"statusDescription,omitempty"`
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

func CreateServerPrivateNetworkFromFile(filename string, commandname string) (*bmcapisdk.ServerPrivateNetwork, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var serverPrivateNetwork ServerPrivateNetwork

	err = files.Unmarshal(data, &serverPrivateNetwork, commandname)

	if err != nil {
		return nil, err
	}

	serverPrivateNetworkSdk := serverPrivateNetwork.toSdk()

	return &serverPrivateNetworkSdk, nil
}
