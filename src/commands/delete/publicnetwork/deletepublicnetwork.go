package publicnetwork

import (
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	ipblock "phoenixnap.com/pnapctl/commands/delete/publicnetwork/ip-block"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

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
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return deletePublicNetwork(args[0])
	},
}

func deletePublicNetwork(id string) error {
	log.Info().Msgf("Deleting Public Network with ID [%s].", id)

	return networks.Client.PublicNetworkDelete(id)
}
