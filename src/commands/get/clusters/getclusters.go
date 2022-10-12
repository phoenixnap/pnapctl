package clusters

import (
	"phoenixnap.com/pnapctl/common/client/rancher"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"

	"github.com/spf13/cobra"
)

var GetClustersCmd = &cobra.Command{
	Use:          "cluster [CLUSTER_ID]",
	Short:        "Retrieve one or all clusters.",
	Aliases:      []string{"clusters"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all clusters.
	
Prints information about the clusters.
By default, the data is printed in table format.

To print a specific cluster, an ID needs to be passed as an argument.`,
	Example: `
# List all clusters.
pnapctl get clusters [--output <OUTPUT_TYPE>]

# List a specific cluster.
pnapctl get cluster <CLUSTER_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		if len(args) >= 1 {
			return getClusterById(args[0])
		}
		return getClusters()
	},
}

func getClusters() error {
	clusters, httpResponse, err := rancher.Client.ClustersGet()

	var generatedError = utils.CheckErrs(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		return printer.PrintClusterListResponse(clusters)
	}
}

func getClusterById(clusterID string) error {
	cluster, httpResponse, err := rancher.Client.ClusterGetById(clusterID)

	var generatedError = utils.CheckErrs(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		return printer.PrintClusterResponse(cluster)
	}
}

func init() {
	GetClustersCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
