package networkmodels

import (
	"time"

	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type PublicNetwork struct {
	Id          string                 `json:"id" yaml:"id"`
	VlanId      int32                  `json:"vlanId" yaml:"vlanId"`
	Memberships []NetworkMembership    `json:"memberships" yaml:"memberships"`
	Name        string                 `json:"name" yaml:"name"`
	Location    string                 `json:"location" yaml:"location"`
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	CreatedOn   time.Time              `json:"createdOn" yaml:"createdOn"`
	IpBlocks    []PublicNetworkIpBlock `json:"ipBlocks" yaml:"ipBlocks"`
}

type NetworkMembership struct {
	ResourceId   string   `json:"resourceId" yaml:"resourceId"`
	ResourceType string   `json:"resourceType" yaml:"resourceType"`
	Ips          []string `json:"ips" yaml:"ips"`
}

type PublicNetworkIpBlock struct {
	Id string `json:"id" yaml:"id"`
}

func PublicNetworkFromSdk(sdk networkapi.PublicNetwork) PublicNetwork {
	return PublicNetwork{
		Id:          sdk.Id,
		VlanId:      sdk.VlanId,
		Memberships: iterutils.Map(sdk.Memberships, NetworkMembershipFromSdk),
		Name:        sdk.Name,
		Location:    sdk.Location,
		Description: sdk.Description,
		CreatedOn:   sdk.CreatedOn,
		IpBlocks:    iterutils.Map(sdk.IpBlocks, PublicNetworkIpBlockFromSdk),
	}
}

func NetworkMembershipFromSdk(sdk networkapi.NetworkMembership) NetworkMembership {
	return NetworkMembership{
		ResourceId:   sdk.ResourceId,
		ResourceType: sdk.ResourceType,
		Ips:          sdk.Ips,
	}
}

func PublicNetworkIpBlockFromSdk(sdk networkapi.PublicNetworkIpBlock) PublicNetworkIpBlock {
	return PublicNetworkIpBlock{
		Id: sdk.Id,
	}
}

func (cli *PublicNetworkIpBlock) ToSdk() networkapi.PublicNetworkIpBlock {
	return networkapi.PublicNetworkIpBlock{
		Id: cli.Id,
	}
}
