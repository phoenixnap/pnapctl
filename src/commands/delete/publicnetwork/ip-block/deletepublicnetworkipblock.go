package ipblock

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "delete public-network ip-block"

var DeletePublicNetworkIpBlockCmd = &cobra.Command{
	Use:          "ip-block [ID]",
	Short:        "Delete an ip-block on a public network.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long:         `Delete an ip-block on a public network.`,
	Example: `# Delete an ip-block on a public network.
pnapctl delete public-network ip-block <NETWORK_ID> <IP_BLOCK_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		message, response, err := networks.Client.PublicNetworkIpBlockDelete(args[0], args[1])

		if err := utils.CheckForErrors(response, err, commandName); err != nil {
			return *err
		} else {
			fmt.Println(message)
		}
		return nil
	},
}

func init() {

}
