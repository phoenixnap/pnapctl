package privatenetwork

import (
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var DeletePrivateNetworkCmd = &cobra.Command{
	Use:          "private-network PRIVATE_NETWORK_ID",
	Short:        "Deletes a specific private network.",
	Long:         "Deletes a specific private network.",
	Example:      `pnapctl delete private-network <PRIVATE_NETWORK_ID>`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return deletePrivateNetwork(args[0])
	},
}

func deletePrivateNetwork(id string) error {
	log.Info().Msgf("Removing Private Network with ID [%s].", id)

	return networks.Client.PrivateNetworkDelete(id)
}
