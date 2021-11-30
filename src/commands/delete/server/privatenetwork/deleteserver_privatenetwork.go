package privatenetwork

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "delete server-private-network"

// DeleteServerPrivateNetworkCmd is the command for creating a server.
var DeleteServerPrivateNetworkCmd = &cobra.Command{
	Use:          "server-private-network SERVER_ID PRIVATE_NETWORK_ID",
	Short:        "Remove a server from a private network.",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	Long: `Remove a server from a private network.

Requires two IDs passed as arguments. First one being the server id and second being the private network id. `,
	Example: `# remove a server from a private network 
pnapctl delete server-private-network <SERVER_ID> <PRIVATE_NETWORK_ID>
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, httpResponse, err := bmcapi.Client.ServerPrivateNetworkDelete(args[0], args[1])

		if httpResponse != nil && !utils.Is2xxSuccessful(httpResponse.StatusCode) {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		} else if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else {
			fmt.Println(result)
			return nil
		}
	},
}
