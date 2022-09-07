package publicnetwork

import (
	"fmt"

	"github.com/spf13/cobra"
	ipblock "phoenixnap.com/pnapctl/commands/delete/publicnetwork/ip-block"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "delete public-network"

var DeletePublicNetworkCmd = &cobra.Command{
	Use:          "public-network [ID]",
	Short:        "Delete a public network.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long:         `Delete a public network.`,
	Example: `# Delete a public network
pnapctl delete public-network <ID>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := networks.Client.PublicNetworkDelete(args[0])

		generatedErr := utils.CheckForErrors(response, err, commandName)

		if *generatedErr != nil {
			return *generatedErr
		} else {
			fmt.Println("Successfully deleted.")
		}

		return nil
	},
}

func init() {
	DeletePublicNetworkCmd.AddCommand(ipblock.DeletePublicNetworkIpBlockCmd)
}
