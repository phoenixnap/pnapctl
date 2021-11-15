package cluster

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/rancher"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

const commandName = "delete cluster"

var DeleteClusterCmd = &cobra.Command{
	Use:          "cluster CLUSTER_ID",
	Short:        "Deletes a specific cluster.",
	Long:         "Deletes a specific cluster.",
	Example:      `pnapctl delete cluster 5da891e90ab0c59bd28e34ad`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, httpResponse, err := rancher.Client.ClusterDelete(args[0])

		if err != nil {
			return err
		} else if httpResponse.StatusCode != 200 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}

		fmt.Println(result.Result, result.ClusterId)
		return nil
	},
}
