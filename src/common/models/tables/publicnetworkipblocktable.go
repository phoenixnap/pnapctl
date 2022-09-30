package tables

import "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"

type PublicNetworkIpBlockTable struct {
	Id string `header:"ID"`
}

func PublicNetworkIpBlockTableFromSdk(sdk networkapi.PublicNetworkIpBlock) PublicNetworkIpBlockTable {
	return PublicNetworkIpBlockTable{
		Id: sdk.Id,
	}
}
