package printer

import (
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/common/models/bmcapimodels"
	"phoenixnap.com/pnap-cli/common/models/tables"
)

func PrintServerResponse(server bmcapisdk.Server, full bool, commandName string) error {
	serverToPrint := PrepareServerForPrinting(server, full)
	return MainPrinter.PrintOutput(serverToPrint, commandName)
}

func PrintServerListResponse(servers []bmcapisdk.Server, full bool, commandName string) error {
	serverListToPrint := PrepareServerListForPrinting(servers, full)
	return MainPrinter.PrintOutput(serverListToPrint, commandName)
}

func PrintServerPrivateNetwork(serverPrivateNetwork bmcapisdk.ServerPrivateNetwork, commandName string) error {
	table := OutputIsTable()

	if table {
		return MainPrinter.PrintOutput(tables.ToServerPrivateNetworkTable(serverPrivateNetwork), commandName)
	} else {
		return MainPrinter.PrintOutput(serverPrivateNetwork, commandName)
	}
}

func PrepareServerForPrinting(server bmcapisdk.Server, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case full && table:
		return tables.ToLongServerTable(server)
	case !full && table:
		return tables.ToShortServerTable(server)
	case full:
		return bmcapimodels.ToFullServer(server)
	default:
		return bmcapimodels.ToShortServer(server)
	}
}

func PrepareServerListForPrinting(servers []bmcapisdk.Server, full bool) []interface{} {
	var serverList []interface{}

	for _, bmcServer := range servers {
		serverList = append(serverList, PrepareServerForPrinting(bmcServer, full))
	}

	return serverList
}
