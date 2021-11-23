package networkmodels

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

type PrivateNetworkModify struct {
	Name            string
	Description     *string
	LocationDefault bool
}

func (modify *PrivateNetworkModify) toSdk() networksdk.PrivateNetworkModify {
	return networksdk.PrivateNetworkModify{
		Name:            modify.Name,
		Description:     modify.Description,
		LocationDefault: modify.LocationDefault,
	}
}

func PrivateNetworkModifyFromSdk(modify *networksdk.PrivateNetworkModify) *PrivateNetworkModify {
	if modify == nil {
		return nil
	}

	return &PrivateNetworkModify{
		Name:            modify.Name,
		Description:     modify.Description,
		LocationDefault: modify.LocationDefault,
	}
}
