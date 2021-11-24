package tables

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnap-cli/common/models/networkmodels"
)

type PrivateNetwork struct {
	Id              string   `header:"ID"`
	Name            string   `header:"Name"`
	Description     *string  `header:"Description"`
	VlanId          int32    `header:"Vlan ID"`
	Type            string   `header:"Type"`
	Location        string   `header:"Location"`
	LocationDefault bool     `header:"Location Default"`
	Cidr            string   `header:"Cidr"`
	Servers         []string `header:"Servers"`
}

func PrivateNetworkFromSdk(network networksdk.PrivateNetwork) PrivateNetwork {
	var servers []string

	for _, server := range network.Servers {
		servers = append(servers, networkmodels.PrivateNetworkServerToTableString(&server))
	}

	return PrivateNetwork{
		Id:              network.Id,
		Name:            DerefString(network.Description),
		Description:     network.Description,
		VlanId:          network.VlanId,
		Type:            network.Type,
		Location:        network.Location,
		LocationDefault: network.LocationDefault,
		Cidr:            network.Cidr,
		Servers:         servers,
	}
}
