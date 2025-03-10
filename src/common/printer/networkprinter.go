package printer

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func PrintPrivateNetworkResponse(network *networksdk.PrivateNetwork) error {
	networkToPrint := PreparePrivateNetworkForPrinting(*network)
	return MainPrinter.PrintOutput(networkToPrint)
}

func PrintPrivateNetworkListResponse(networks []networksdk.PrivateNetwork) error {
	networkListToPrint := iterutils.Map(networks, PreparePrivateNetworkForPrinting)
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

func PrintBgpPeerGroupResponse(bgpPeerGroup *networksdk.BgpPeerGroup) error {
	bgpPeerGroupToPrint := PrepareBgpPeerGroupForPrinting(*bgpPeerGroup)
	return MainPrinter.PrintOutput(bgpPeerGroupToPrint)
}

func PrintBgpPeerGroupListResponse(bgpPeerGroups []networksdk.BgpPeerGroup) error {
	bgpPeerGroupsToPrint := iterutils.Map(bgpPeerGroups, PrepareBgpPeerGroupForPrinting)
	return MainPrinter.PrintOutput(bgpPeerGroupsToPrint)
}

func PrepareBgpPeerGroupForPrinting(bgpPeerGroup networksdk.BgpPeerGroup) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.BgpPeerGroupFromSdk(bgpPeerGroup)
	default:
		return bgpPeerGroup
	}
}
