package networkmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type PublicNetworkCreate struct {
	Name        string                 `json:"name" yaml:"name"`
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Location    string                 `json:"location" yaml:"location"`
	IpBlocks    []PublicNetworkIpBlock `json:"ipBlocks,omitempty" yaml:"ipBlocks,omitempty"`
}

func (cli *PublicNetworkCreate) ToSdk() networkapi.PublicNetworkCreate {
	return networkapi.PublicNetworkCreate{
		Name:        cli.Name,
		Description: cli.Description,
		Location:    cli.Location,
		IpBlocks: iterutils.Map(cli.IpBlocks, func(ipBlock PublicNetworkIpBlock) networkapi.PublicNetworkIpBlock {
			return ipBlock.ToSdk()
		}),
	}
}
