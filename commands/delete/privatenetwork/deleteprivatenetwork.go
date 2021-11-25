package privatenetwork

import (
	"github.com/spf13/cobra"

	"phoenixnap.com/pnap-cli/common/client/networks"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

const commandName = "delete private-network"

var DeletePrivateNetworkCmd = &cobra.Command{
	Use:          "private-network SERVER_ID",
	Short:        "Deletes a specific private network.",
	Long:         "Deletes a specific private network.",
	Example:      `pnapctl delete private-network 5da891e90ab0c59bd28e34ad`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		httpResponse, err := networks.Client.PrivateNetworkDelete(args[0])

		if httpResponse != nil && httpResponse.StatusCode != 200 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		} else if err != nil {
			return err
		}

		return nil
	},
}
