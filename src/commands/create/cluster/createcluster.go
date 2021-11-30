package cluster

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/rancher"
	"phoenixnap.com/pnapctl/common/models/ranchermodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
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
	Example: `# Create a new cluster as described in clusterCreate.yaml
pnapctl create cluster --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# clusterCreate.yaml
location: PHX
name: rancher-cluster-test
nodePools:
  - serverType: s1.c1.medium
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cluster, err := ranchermodels.CreateClusterFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		response, httpResponse, err := rancher.Client.ClusterPost(*cluster)
		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintClusterResponse(response, commandName)
		}
	},
}

func init() {
	CreateClusterCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreateClusterCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreateClusterCmd.MarkFlagRequired("filename")
}
