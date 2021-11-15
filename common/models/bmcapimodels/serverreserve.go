package bmcapimodels

import (
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	files "phoenixnap.com/pnap-cli/common/fileprocessor"
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

	return serverReserveToSDK(&serverReserve), nil
}

func serverReserveToSDK(reserveRequest *ServerReserve) *bmcapisdk.ServerReserve {
	return &bmcapisdk.ServerReserve{
		PricingModel: reserveRequest.PricingModel,
	}
}
