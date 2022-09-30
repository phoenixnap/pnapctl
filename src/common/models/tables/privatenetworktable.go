package tables

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
)

type PrivateNetworkTable struct {
	Id              string   `header:"ID"`
	Name            string   `header:"Name"`
	Description     string   `header:"Description"`
	VlanId          int32    `header:"Vlan ID"`
	Type            string   `header:"Type"`
	Location        string   `header:"Location"`
	LocationDefault bool     `header:"Location Default"`
	Cidr            string   `header:"Cidr"`
	Servers         []string `header:"Servers"`
}

func PrivateNetworkFromSdk(network networksdk.PrivateNetwork) PrivateNetworkTable {
	var servers []string

	for _, server := range network.Servers {
		servers = append(servers, networkmodels.PrivateNetworkServerToTableString(&server))
	}

	return PrivateNetworkTable{
		Id:              network.Id,
		Name:            network.Name,
		Description:     DerefString(network.Description),
		VlanId:          network.VlanId,
		Type:            network.Type,
		Location:        network.Location,
		LocationDefault: network.LocationDefault,
		Cidr:            network.Cidr,
		Servers:         servers,
	}
}
