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
	RunE: func(cmd *cobra.Command, args []string) error {
		result, httpResponse, err := rancher.Client.ClusterDelete(args[0])
		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			fmt.Println(result.Result, result.ClusterId)
			return nil
		}
	},
}
