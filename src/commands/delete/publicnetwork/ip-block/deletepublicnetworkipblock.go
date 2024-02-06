package ipblock

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	"phoenixnap.com/pnapctl/common/utils"
)

var (
	force bool
)

func init() {
	utils.SetupOutputFlag(DeletePublicNetworkIpBlockCmd)

	DeletePublicNetworkIpBlockCmd.Flags().BoolVar(&force, "force", false, "Controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. Defaults to false.")
}

var DeletePublicNetworkIpBlockCmd = &cobra.Command{
	Use:          "ip-block [ID]",
	Short:        "Delete an ip-block on a public network.",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	Long:         `Delete an ip-block on a public network.`,
	Example: `# Delete an ip-block on a public network.
pnapctl delete public-network ip-block <NETWORK_ID> <IP_BLOCK_ID> [--output <OUTPUT_TYPE>] [--force=false]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return deleteIpBlockFromPublicNetwork(args[0], args[1])
	},
}

func deleteIpBlockFromPublicNetwork(networkId, ipBlockId string) error {
	log.Info().Msgf("Removing Ip Block with ID [%s] in Public Network with ID [%s].", ipBlockId, networkId)

	message, err := networks.Client.PublicNetworkIpBlockDelete(networkId, ipBlockId, force)

	if err != nil {
		return err
	} else {
		fmt.Println(message)
	}
	return nil
}
