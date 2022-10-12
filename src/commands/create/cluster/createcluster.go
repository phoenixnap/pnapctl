package cluster

import (
	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/rancher"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var Filename string

func init() {
	utils.SetupOutputFlag(CreateClusterCmd)
	utils.SetupFilenameFlag(CreateClusterCmd, &Filename, utils.CREATION)
}

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
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return createCluster()
	},
}

func createCluster() error {
	cluster, err := models.CreateRequestFromFile[ranchersolutionapi.Cluster](Filename)

	if err != nil {
		return err
	}

	response, httpResponse, err := rancher.Client.ClusterPost(*cluster)
	var generatedError = utils.CheckErrs(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		return printer.PrintClusterResponse(response)
	}
}
