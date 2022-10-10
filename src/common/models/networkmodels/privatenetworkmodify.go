package networkmodels

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type PrivateNetworkModify struct {
	Name            string  `json:"name" yaml:"name"`
	Description     *string `json:"description" yaml:"description"`
	LocationDefault bool    `json:"locationDefault" yaml:"locationDefault"`
}

func (modify *PrivateNetworkModify) ToSdk() *networksdk.PrivateNetworkModify {
	return &networksdk.PrivateNetworkModify{
		Name:            modify.Name,
		Description:     modify.Description,
		LocationDefault: modify.LocationDefault,
	}
}

func CreatePrivateNetworkUpdateFromFile(filename string, commandname string) (*networksdk.PrivateNetworkModify, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var privateNetworkModify PrivateNetworkModify

	err = files.Unmarshal(data, &privateNetworkModify, commandname)

	if err != nil {
		return nil, err
	}

	return privateNetworkModify.ToSdk(), nil
}
