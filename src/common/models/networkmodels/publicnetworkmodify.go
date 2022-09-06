package networkmodels

import "github.com/phoenixnap/go-sdk-bmc/networkapi"

type PublicNetworkModify struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (cli *PublicNetworkModify) ToSdk() networkapi.PublicNetworkModify {
	return networkapi.PublicNetworkModify{
		Name:        cli.Name,
		Description: cli.Description,
	}
}
