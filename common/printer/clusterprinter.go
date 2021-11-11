package printer

import (
	ranchersdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/ranchersolutionapi"
	"phoenixnap.com/pnap-cli/common/models/ranchermodels"
	"phoenixnap.com/pnap-cli/common/models/tables"
)

func PrintClusterResponse(cluster ranchersdk.Cluster, commandName string) error {
	clusterToPrint := PrepareClusterForPrinting(cluster)
	return MainPrinter.PrintOutput(clusterToPrint, commandName)
}

func PrintClusterListResponse(clusters []ranchersdk.Cluster, commandName string) error {
	clusterListToPrint := PrepareClusterListForPrinting(clusters)
	return MainPrinter.PrintOutput(clusterListToPrint, commandName)
}

func PrepareClusterForPrinting(cluster ranchersdk.Cluster) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ClusterFromSdk(cluster)
	default:
		return ranchermodels.ClusterFromSdk(cluster)
	}
}

func PrepareClusterListForPrinting(clusters []ranchersdk.Cluster) []interface{} {
	var clusterList []interface{}

	for _, cluster := range clusters {
		clusterList = append(clusterList, PrepareClusterForPrinting(cluster))
	}

	return clusterList
}
