package publicnetwork

import (
	"fmt"

	"github.com/spf13/cobra"
	ipblock "phoenixnap.com/pnapctl/commands/delete/publicnetwork/ip-block"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "delete public-network"

func init() {
	DeletePublicNetworkCmd.AddCommand(ipblock.DeletePublicNetworkIpBlockCmd)
}

var DeletePublicNetworkCmd = &cobra.Command{
	Use:          "public-network [ID]",
	Short:        "Deletes a public network.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long:         `Delete a public network.`,
	Example: `# Delete a public network
pnapctl delete public-network <ID>`,
	RunE: func(_ *cobra.Command, args []string) error {
		return deletePublicNetwork(args[0])
	},
}

func deletePublicNetwork(id string) error {
	response, err := networks.Client.PublicNetworkDelete(id)

	generatedErr := utils.CheckForErrors(response, err, commandName)

	if *generatedErr != nil {
		return *generatedErr
	} else {
		fmt.Println("Successfully deleted.")
	}

	return nil
}
