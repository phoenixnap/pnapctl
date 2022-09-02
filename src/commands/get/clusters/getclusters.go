package clusters

import (
	netHttp "net/http"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"phoenixnap.com/pnapctl/common/client/rancher"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"

	"github.com/spf13/cobra"
)

const commandName string = "get clusters"

var ID string

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
		if len(args) >= 1 {
			ID = args[0]
			return getClusters(ID)
		}
		return getClusters("")
	},
}

func getClusters(clusterID string) error {
	var httpResponse *netHttp.Response
	var err error
	var cluster *ranchersdk.Cluster
	var clusters []ranchersdk.Cluster

	if clusterID == "" {
		clusters, httpResponse, err = rancher.Client.ClustersGet()
	} else {
		cluster, httpResponse, err = rancher.Client.ClusterGetById(clusterID)
	}

	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		if clusterID == "" {
			return printer.PrintClusterListResponse(clusters, commandName)
		} else {
			return printer.PrintClusterResponse(cluster, commandName)
		}
	}
}

func init() {
	GetClustersCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
