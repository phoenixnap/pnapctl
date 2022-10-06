package printer

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
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
		return network
	}
}

func PreparePrivateNetworkListForPrinting(networks []networksdk.PrivateNetwork) []interface{} {
	var networkList []interface{}

	for _, network := range networks {
		networkList = append(networkList, PreparePrivateNetworkForPrinting(network))
	}

	return networkList
}

func PrintPublicNetworkResponse(network *networksdk.PublicNetwork, commandName string) error {
	networkToPrint := PreparePublicNetworkForPrinting(*network)
	return MainPrinter.PrintOutput(networkToPrint, commandName)
}

func PrintPublicNetworkListResponse(network []networksdk.PublicNetwork, commandName string) error {
	networksToPrint := iterutils.Map(network, PreparePublicNetworkForPrinting)
	return MainPrinter.PrintOutput(networksToPrint, commandName)
}

func PreparePublicNetworkForPrinting(network networksdk.PublicNetwork) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.PublicNetworkTableFromSdk(network)
	default:
		return network
	}
}

func PrintPublicNetworkIpBlockResponse(ipBlock *networksdk.PublicNetworkIpBlock, commandName string) error {
	networkToPrint := PreparePublicNetworkIpBlockForPrinting(*ipBlock)
	return MainPrinter.PrintOutput(networkToPrint, commandName)
}

func PrintPublicNetworkIpBlockListResponse(ipBlocks []networksdk.PublicNetworkIpBlock, commandName string) error {
	ipBlocksToPrint := iterutils.Map(ipBlocks, PreparePublicNetworkIpBlockForPrinting)
	return MainPrinter.PrintOutput(ipBlocksToPrint, commandName)
}

func PreparePublicNetworkIpBlockForPrinting(ipBlock networksdk.PublicNetworkIpBlock) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.PublicNetworkIpBlockTableFromSdk(ipBlock)
	default:
		return ipBlock
	}
}
