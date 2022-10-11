package cluster

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/rancher"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName = "delete cluster"

var DeleteClusterCmd = &cobra.Command{
	Use:          "cluster CLUSTER_ID",
	Short:        "Deletes a specific cluster.",
	Long:         "Deletes a specific cluster.",
	Example:      `pnapctl delete cluster <CLUSTER_ID>`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(_ *cobra.Command, args []string) error {
		return deleteCluster(args[0])
	},
}

func deleteCluster(id string) error {
	result, httpResponse, err := rancher.Client.ClusterDelete(id)
	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		fmt.Println(result.Result, result.ClusterId)
		return nil
	}
}
