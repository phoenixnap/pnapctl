package private_network

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "delete server-private-network"

// DeleteServerPrivateNetworkCmd is the command for creating a server.
var DeleteServerPrivateNetworkCmd = &cobra.Command{
	Use:          "server-private-network [SEVER_ID] [PRIVATE_NETWORK_ID]",
	Short:        "Remove a server from a private network.",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	Long: `Remove a server from a private network.

Requires two IDs passed as arguments. First one being the server id and second being the private network id. `,
	Example: `# remove a server from a private network 
pnapctl delete server-private-network 5ff5cc9bc1acf144d910621f 5ff5cc9bc1acf144d9106233
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, httpResponse, err := bmcapi.Client.ServerPrivateNetworkDelete(args[0], args[1])

		if err != nil {
			return err
		} else if httpResponse.StatusCode != 202 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}

		fmt.Println(result)
		return nil
	},
}
