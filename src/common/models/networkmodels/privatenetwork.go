package networkmodels

import (
	"time"

	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type PrivateNetwork struct {
	Id              string                 `json:"id" yaml:"id"`
	Name            string                 `json:"name" yaml:"name"`
	Description     *string                `json:"description" yaml:"description"`
	VlanId          int32                  `json:"vlanId" yaml:"vlanId"`
	Type            string                 `json:"type" yaml:"type"`
	Location        string                 `json:"location" yaml:"location"`
	LocationDefault bool                   `json:"locationDefault" yaml:"locationDefault"`
	Cidr            string                 `json:"cidr" yaml:"cidr"`
	Servers         []PrivateNetworkServer `json:"server" yaml:"server"`
	Memberships     []NetworkMembership    `json:"memberships" yaml:"memberships"`
	CreatedOn       time.Time              `json:"createdOn" yaml:"createdOn"`
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
		Memberships:     iterutils.Map(network.Memberships, NetworkMembershipFromSdk),
		CreatedOn:       network.CreatedOn,
	}
}
