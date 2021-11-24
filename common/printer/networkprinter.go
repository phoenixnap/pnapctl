package printer

import (
	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"phoenixnap.com/pnap-cli/common/models/networkmodels"
	"phoenixnap.com/pnap-cli/common/models/tables"
)

func PrintPrivateNetworkResponse(cluster networksdk.PrivateNetwork, commandName string) error {
	clusterToPrint := PreparePrivateNetworkForPrinting(cluster)
	return MainPrinter.PrintOutput(clusterToPrint, commandName)
}

func PrintPrivateNetworkListResponse(clusters []networksdk.PrivateNetwork, commandName string) error {
	clusterListToPrint := PreparePrivateNetworkListForPrinting(clusters)
	return MainPrinter.PrintOutput(clusterListToPrint, commandName)
}

func PreparePrivateNetworkForPrinting(cluster networksdk.PrivateNetwork) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.PrivateNetworkFromSdk(cluster)
	default:
		return networkmodels.PrivateNetworkFromSdk(cluster)
	}
}

func PreparePrivateNetworkListForPrinting(clusters []networksdk.PrivateNetwork) []interface{} {
	var clusterList []interface{}

	for _, cluster := range clusters {
		clusterList = append(clusterList, PreparePrivateNetworkForPrinting(cluster))
	}

	return clusterList
}
