package networkmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type PublicNetworkModify struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (cli *PublicNetworkModify) ToSdk() *networkapi.PublicNetworkModify {
	return &networkapi.PublicNetworkModify{
		Name:        cli.Name,
		Description: cli.Description,
	}
}

func CreatePublicNetworkModifyFromFile(filename string, commandName string) (*networkapi.PublicNetworkModify, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandName)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var publicNetworkModify PublicNetworkModify

	err = files.Unmarshal(data, &publicNetworkModify, commandName)

	if err != nil {
		return nil, err
	}

	return publicNetworkModify.ToSdk(), nil
}
