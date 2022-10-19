package ipblock

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var DeletePublicNetworkIpBlockCmd = &cobra.Command{
	Use:          "ip-block [ID]",
	Short:        "Delete an ip-block on a public network.",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	Long:         `Delete an ip-block on a public network.`,
	Example: `# Delete an ip-block on a public network.
pnapctl delete public-network ip-block <NETWORK_ID> <IP_BLOCK_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return deleteIpBlockFromPublicNetwork(args[0], args[1])
	},
}

func deleteIpBlockFromPublicNetwork(networkId, ipBlockId string) error {
	message, err := networks.Client.PublicNetworkIpBlockDelete(networkId, ipBlockId)

	if err != nil {
		return err
	} else {
		fmt.Println(message)
	}
	return nil
}
