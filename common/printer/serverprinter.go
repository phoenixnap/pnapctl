package printer

import (
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/common/models"
	"phoenixnap.com/pnap-cli/common/models/tables"
)

func PrintServerResponse(server bmcapi.Server, full bool, commandName string) error {
	serverToPrint := PrepareServerForPrinting(server, full)
	return MainPrinter.PrintOutput(serverToPrint, commandName)
}

func PrintServerListResponse(servers []bmcapi.Server, full bool, commandName string) error {
	serverListToPrint := PrepareServerListForPrinting(servers, full)
	return MainPrinter.PrintOutput(serverListToPrint, commandName)
}

func PrepareServerForPrinting(server bmcapi.Server, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case full && table:
		return tables.ToLongServerTable(server)
	case !full && table:
		return tables.ToShortServerTable(server)
	case full:
		return models.ToFullServer(server)
	default:
		return models.ToShortServer(server)
	}
}

func PrepareServerListForPrinting(servers []bmcapi.Server, full bool) []interface{} {
	var serverList []interface{}

	for _, bmcServer := range servers {
		serverList = append(serverList, PrepareServerForPrinting(bmcServer, full))
	}

	return serverList
}
