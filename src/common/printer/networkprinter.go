package printer

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func PrintPrivateNetworkResponse(network *networksdk.PrivateNetwork) error {
	networkToPrint := PreparePrivateNetworkForPrinting(*network)
	return MainPrinter.PrintOutput(networkToPrint)
}

func PrintPrivateNetworkListResponse(networks []networksdk.PrivateNetwork) error {
	networkListToPrint := PreparePrivateNetworkListForPrinting(networks)
	return MainPrinter.PrintOutput(networkListToPrint)
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

func PrintPublicNetworkResponse(network *networksdk.PublicNetwork) error {
	networkToPrint := PreparePublicNetworkForPrinting(*network)
	return MainPrinter.PrintOutput(networkToPrint)
}

func PrintPublicNetworkListResponse(network []networksdk.PublicNetwork) error {
	networksToPrint := iterutils.Map(network, PreparePublicNetworkForPrinting)
	return MainPrinter.PrintOutput(networksToPrint)
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

func PrintPublicNetworkIpBlockResponse(ipBlock *networksdk.PublicNetworkIpBlock) error {
	networkToPrint := PreparePublicNetworkIpBlockForPrinting(*ipBlock)
	return MainPrinter.PrintOutput(networkToPrint)
}

func PrintPublicNetworkIpBlockListResponse(ipBlocks []networksdk.PublicNetworkIpBlock) error {
	ipBlocksToPrint := iterutils.Map(ipBlocks, PreparePublicNetworkIpBlockForPrinting)
	return MainPrinter.PrintOutput(ipBlocksToPrint)
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
