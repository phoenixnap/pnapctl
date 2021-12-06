package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type ServerPrivateNetwork struct {
	Id                string    `yaml:"id,omitempty" json:"id,omitempty"`
	Ips               *[]string `yaml:"ips,omitempty" json:"ips,omitempty"`
	Dhcp              *bool     `yaml:"dhcp,omitempty" json:"dhcp,omitempty"`
	StatusDescription *string   `yaml:"statusDescription,omitempty" json:"statusDescription,omitempty"`
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

func privateNetworksFromSdk(privateNetworks *[]bmcapisdk.ServerPrivateNetwork) *[]ServerPrivateNetwork {
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
