package printer

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
	"phoenixnap.com/pnapctl/common/models/tables"
)

func PrintPrivateNetworkResponse(network *networksdk.PrivateNetwork, commandName string) error {
	networkToPrint := PreparePrivateNetworkForPrinting(*network)
	return MainPrinter.PrintOutput(networkToPrint, commandName)
}

func PrintPrivateNetworkListResponse(networks []networksdk.PrivateNetwork, commandName string) error {
	networkListToPrint := PreparePrivateNetworkListForPrinting(networks)
	return MainPrinter.PrintOutput(networkListToPrint, commandName)
}

func PreparePrivateNetworkForPrinting(network networksdk.PrivateNetwork) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.PrivateNetworkFromSdk(network)
	default:
		return networkmodels.PrivateNetworkFromSdk(network)
	}
}

func PreparePrivateNetworkListForPrinting(networks []networksdk.PrivateNetwork) []interface{} {
	var networkList []interface{}

	for _, network := range networks {
		networkList = append(networkList, PreparePrivateNetworkForPrinting(network))
	}

	return networkList
}
