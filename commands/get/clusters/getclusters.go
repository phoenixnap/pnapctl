package clusters

import (
	netHttp "net/http"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"phoenixnap.com/pnap-cli/common/client/rancher"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/printer"

	log "github.com/sirupsen/logrus"
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

To print a single cluster, an ID needs to be passed as an argument.`,
	Example: `
# List all clusters in json format.
pnapctl get clusters -o json

# List a single cluster in yaml format.
pnapctl get cluster NDIid939dfkoDd -o yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			ID = args[0]
			return getClusters(ID)
		}
		return getClusters("")
	},
}

func getClusters(clusterID string) error {
	log.Debug("Getting clusters...")

	var httpResponse *netHttp.Response
	var err error
	var cluster ranchersdk.Cluster
	var clusters []ranchersdk.Cluster

	if clusterID == "" {
		clusters, httpResponse, err = rancher.Client.ClustersGet()
	} else {
		cluster, httpResponse, err = rancher.Client.ClusterGetById(clusterID)
	}

	if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else if httpResponse.StatusCode == 200 {
		if clusterID == "" {
			return printer.PrintClusterListResponse(clusters, commandName)
		} else {
			return printer.PrintClusterResponse(cluster, commandName)
		}
	} else {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	}
}

func init() {
	GetClustersCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
