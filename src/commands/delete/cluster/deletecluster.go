package cluster

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/rancher"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var DeleteClusterCmd = &cobra.Command{
	Use:          "cluster CLUSTER_ID",
	Short:        "Deletes a specific cluster.",
	Long:         "Deletes a specific cluster.",
	Example:      `pnapctl delete cluster <CLUSTER_ID>`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return deleteCluster(args[0])
	},
}

func deleteCluster(id string) error {
	log.Info().Msgf("Deleting Cluster with ID [%s].", id)

	result, err := rancher.Client.ClusterDelete(id)
	if err != nil {
		return err
	} else {
		fmt.Println(result.Result, result.ClusterId)
		return nil
	}
}
