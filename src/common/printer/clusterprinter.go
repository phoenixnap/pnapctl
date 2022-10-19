package printer

import (
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func PrintClusterResponse(cluster *ranchersdk.Cluster) error {
	clusterToPrint := PrepareClusterForPrinting(*cluster)
	return MainPrinter.PrintOutput(clusterToPrint)
}

func PrintClusterListResponse(clusters []ranchersdk.Cluster) error {
	clusterListToPrint := iterutils.Map(clusters, PrepareClusterForPrinting)
	return MainPrinter.PrintOutput(clusterListToPrint)
}

func PrepareClusterForPrinting(cluster ranchersdk.Cluster) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ClusterFromSdk(cluster)
	default:
		return cluster
	}
}
