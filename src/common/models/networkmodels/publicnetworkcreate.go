package networkmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type PublicNetworkCreate struct {
	Name        string                 `json:"name" yaml:"name"`
	Description *string                `json:"description,omitempty" yaml:"description,omitempty"`
	Location    string                 `json:"location" yaml:"location"`
	IpBlocks    []PublicNetworkIpBlock `json:"ipBlocks,omitempty" yaml:"ipBlocks,omitempty"`
}

func (cli *PublicNetworkCreate) ToSdk() *networkapi.PublicNetworkCreate {
	return &networkapi.PublicNetworkCreate{
		Name:        cli.Name,
		Description: cli.Description,
		Location:    cli.Location,
		IpBlocks: iterutils.Map(cli.IpBlocks, func(ipBlock PublicNetworkIpBlock) networkapi.PublicNetworkIpBlock {
			return *ipBlock.ToSdk()
		}),
	}
}

func CreatePublicNetworkCreateFromFile(filename string, commandName string) (*networkapi.PublicNetworkCreate, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandName)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var publicNetworkCreate PublicNetworkCreate

	err = files.Unmarshal(data, &publicNetworkCreate, commandName)

	if err != nil {
		return nil, err
	}

	return publicNetworkCreate.ToSdk(), nil
}
