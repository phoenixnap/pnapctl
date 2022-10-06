package tables

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type PublicNetworkTable struct {
	Id          string   `header:"ID"`
	VlanId      int32    `header:"Vlan ID"`
	Memberships []string `header:"Memberships"`
	Name        string   `header:"Name"`
	Location    string   `header:"Location"`
	Description string   `header:"Description"`
	CreatedOn   string   `header:"Created On"`
	IpBlocks    []string `header:"Ip Blocks"`
}

func PublicNetworkTableFromSdk(sdk networkapi.PublicNetwork) PublicNetworkTable {
	return PublicNetworkTable{
		Id:          sdk.Id,
		VlanId:      sdk.VlanId,
		Memberships: iterutils.Map(sdk.Memberships, models.NetworkMembershipToTableString),
		Name:        sdk.Name,
		Location:    sdk.Location,
		Description: DerefString(sdk.Description),
		CreatedOn:   sdk.CreatedOn.String(),
		IpBlocks:    iterutils.Map(sdk.IpBlocks, models.PublicNetworkIpBlockToTableString),
	}
}
