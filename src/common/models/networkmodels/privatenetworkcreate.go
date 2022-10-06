package networkmodels

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type PrivateNetworkCreate struct {
	Name            string  `json:"name" yaml:"name"`
	Description     *string `json:"description" yaml:"description"`
	Location        string  `json:"location" yaml:"location"`
	LocationDefault *bool   `json:"locationDefault" yaml:"locationDefault"`
	Cidr            string  `json:"cidr" yaml:"cidr"`
}

func (create *PrivateNetworkCreate) ToSdk() *networksdk.PrivateNetworkCreate {
	return &networksdk.PrivateNetworkCreate{
		Name:            create.Name,
		Description:     create.Description,
		Location:        create.Location,
		LocationDefault: create.LocationDefault,
		Cidr:            create.Cidr,
	}
}

func CreatePrivateNetworkCreateFromFile(filename string, commandname string) (*networksdk.PrivateNetworkCreate, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var privateNetworkCreate PrivateNetworkCreate

	err = files.Unmarshal(data, &privateNetworkCreate, commandname)

	if err != nil {
		return nil, err
	}

	return privateNetworkCreate.ToSdk(), nil
}
