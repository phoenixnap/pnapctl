package printer

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func PrintServerResponse(server *bmcapisdk.Server, full bool) error {
	serverToPrint := PrepareServerForPrinting(*server, full)
	return MainPrinter.PrintOutput(serverToPrint)
}

func PrintServerListResponse(servers []bmcapisdk.Server, full bool) error {
	serverListToPrint := iterutils.Map(servers, withFull(full, PrepareServerForPrinting))
	return MainPrinter.PrintOutput(serverListToPrint)
}

func PrintServerPrivateNetwork(serverPrivateNetwork *bmcapisdk.ServerPrivateNetwork) error {
	table := OutputIsTable()

	if table {
		return MainPrinter.PrintOutput(tables.ToServerPrivateNetworkTable(*serverPrivateNetwork))
	} else {
		return MainPrinter.PrintOutput(serverPrivateNetwork)
	}
}

func PrintServerPublicNetwork(serverPublicNetwork *bmcapisdk.ServerPublicNetwork) error {
	table := OutputIsTable()

	if table {
		return MainPrinter.PrintOutput(tables.ToServerPublicNetworkTable(*serverPublicNetwork))
	} else {
		return MainPrinter.PrintOutput(serverPublicNetwork)
	}
}

func PrintServerIpBlock(serverIpBlock *bmcapisdk.ServerIpBlock) error {
	table := OutputIsTable()

	if table {
		return MainPrinter.PrintOutput(tables.ToServerIpBlockTable(*serverIpBlock))
	} else {
		return MainPrinter.PrintOutput(serverIpBlock)
	}
}

func PrepareServerForPrinting(server bmcapisdk.Server, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case full && table:
		return tables.ToLongServerTable(server)
	case !full && table:
		return tables.ToShortServerTable(server)
	default:
		return server
	}
}
