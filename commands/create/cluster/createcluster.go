package cluster

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/rancher"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/ranchermodels"
	"phoenixnap.com/pnap-cli/common/printer"
)

var Filename string

var commandName = "create cluster"

var CreateClusterCmd = &cobra.Command{
	Use:          "cluster",
	Short:        "Create a new cluster.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a new cluster.
	
Requires a file (yaml or json) containing the information needed to create the cluster.`,
	Example: `# create a new cluster as described in cluster.yaml
pnapctl create cluster --filename ./cluster.yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cluster, err := ranchermodels.CreateClusterFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		response, httpResponse, err := rancher.Client.ClusterPost(*cluster)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else if httpResponse.StatusCode == 200 {
			return printer.PrintClusterResponse(response, commandName)
		} else {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}
	},
}

func init() {
	CreateClusterCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreateClusterCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreateClusterCmd.MarkFlagRequired("filename")
}
