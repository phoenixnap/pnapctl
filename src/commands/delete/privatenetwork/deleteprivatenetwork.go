package privatenetwork

import (
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/ctlerrors"
)

const commandName = "delete private-network"

var DeletePrivateNetworkCmd = &cobra.Command{
	Use:          "private-network PRIVATE_NETWORK_ID",
	Short:        "Deletes a specific private network.",
	Long:         "Deletes a specific private network.",
	Example:      `pnapctl delete private-network <PRIVATE_NETWORK_ID>`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return deletePrivateNetwork(args[0])
	},
}

func deletePrivateNetwork(id string) error {
	httpResponse, err := networks.Client.PrivateNetworkDelete(id)

	if httpResponse != nil && httpResponse.StatusCode != 204 {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	} else if err != nil {
		return err
	}

	return nil
}
