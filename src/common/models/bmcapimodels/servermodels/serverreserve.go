package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type ServerReserve struct {
	PricingModel string `json:"pricingModel" yaml:"pricingModel"`
}

func CreateReserveRequestFromFile(filename string, commandname string) (*bmcapisdk.ServerReserve, error) {
	files.ExpandPath(&filename)
	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var serverReserve ServerReserve

	err = files.Unmarshal(data, &serverReserve, commandname)

	if err != nil {
		return nil, err
	}

	return serverReserve.toSdk(), nil
}

func (serverReserve ServerReserve) toSdk() *bmcapisdk.ServerReserve {
	return &bmcapisdk.ServerReserve{
		PricingModel: serverReserve.PricingModel,
	}
}
