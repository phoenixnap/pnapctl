package networkmodels

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

type PrivateNetworkCreate struct {
	Name            string
	Description     *string
	Location        string
	LocationDefault *bool
	Cidr            string
}

func (create *PrivateNetworkCreate) toSdk() networksdk.PrivateNetworkCreate {
	return networksdk.PrivateNetworkCreate{
		Name:            create.Name,
		Description:     create.Description,
		Location:        create.Location,
		LocationDefault: create.LocationDefault,
		Cidr:            create.Cidr,
	}
}

func PrivateNetworkCreateFromSdk(create *networksdk.PrivateNetworkCreate) *PrivateNetworkCreate {
	if create == nil {
		return nil
	}

	return &PrivateNetworkCreate{
		Name:            create.Name,
		Description:     create.Description,
		Location:        create.Location,
		LocationDefault: create.LocationDefault,
		Cidr:            create.Cidr,
	}
}
