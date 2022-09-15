package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type ServerPublicNetwork struct {
	Id                string   `yaml:"id,omitempty" json:"id,omitempty"`
	Ips               []string `yaml:"ips,omitempty" json:"ips,omitempty"`
	StatusDescription *string  `yaml:"statusDescription,omitempty" json:"statusDescription,omitempty"`
}

func mapServerPublicNetworkListToSdk(serverPublicNetworks []ServerPublicNetwork) []bmcapisdk.ServerPublicNetwork {
	if serverPublicNetworks == nil {
		return nil
	}

	var bmcServerPublicNetworks []bmcapisdk.ServerPublicNetwork

	for _, serverPublicNetwork := range serverPublicNetworks {
		bmcServerPublicNetworks = append(bmcServerPublicNetworks, serverPublicNetwork.toSdk())
	}

	return bmcServerPublicNetworks
}

func (serverPublicNetwork ServerPublicNetwork) toSdk() bmcapisdk.ServerPublicNetwork {
	var serverPublicNetworkSdk = bmcapisdk.ServerPublicNetwork{
		Id:                serverPublicNetwork.Id,
		Ips:               serverPublicNetwork.Ips,
		StatusDescription: serverPublicNetwork.StatusDescription,
	}

	return serverPublicNetworkSdk
}

func serverPublicNetworkListFromSdk(publicNetworks []bmcapisdk.ServerPublicNetwork) []ServerPublicNetwork {
	if publicNetworks == nil {
		return nil
	}

	var bmcServerPublicNetworks []ServerPublicNetwork

	for _, bmcServerPublicNetwork := range publicNetworks {
		bmcServerPublicNetworks = append(bmcServerPublicNetworks, ServerPublicNetwork{
			Id:                bmcServerPublicNetwork.Id,
			Ips:               bmcServerPublicNetwork.Ips,
			StatusDescription: bmcServerPublicNetwork.StatusDescription,
		})
	}

	return bmcServerPublicNetworks
}

func CreateServerPublicNetworkFromFile(filename string, commandname string) (*bmcapisdk.ServerPublicNetwork, error){
	files.ExpandPath(&filename)

	data,err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	var serverPublicNetwork ServerPublicNetwork

	err = files.Unmarshal(data, &serverPublicNetwork, commandname)

	if err != nil {
		return nil, err
	}

	serverPublicNetworkSdk := serverPublicNetwork.toSdk()

	return &serverPublicNetworkSdk, nil
}

