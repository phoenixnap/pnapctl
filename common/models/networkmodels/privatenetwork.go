package networkmodels

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

type PrivateNetwork struct {
	Id              string
	Name            string
	Description     *string
	VlanId          int32
	Type            string
	Location        string
	LocationDefault bool
	Cidr            string
	Servers         []PrivateNetworkServer
}

func (network *PrivateNetwork) toSdk() networksdk.PrivateNetwork {
	var servers []networksdk.PrivateNetworkServer

	for _, server := range network.Servers {
		servers = append(servers, server.toSdk())
	}

	return networksdk.PrivateNetwork{
		Id:              network.Id,
		Name:            network.Name,
		Description:     network.Description,
		VlanId:          network.VlanId,
		Type:            network.Type,
		Location:        network.Location,
		LocationDefault: network.LocationDefault,
		Cidr:            network.Cidr,
		Servers:         servers,
	}
}

func PrivateNetworkFromSdk(network networksdk.PrivateNetwork) PrivateNetwork {
	var servers []PrivateNetworkServer

	for _, server := range network.Servers {
		servers = append(servers, *PrivateNetworkServerFromSdk(&server))
	}

	return PrivateNetwork{
		Id:              network.Id,
		Name:            network.Name,
		Description:     network.Description,
		VlanId:          network.VlanId,
		Type:            network.Type,
		Location:        network.Location,
		LocationDefault: network.LocationDefault,
		Cidr:            network.Cidr,
		Servers:         servers,
	}
}
