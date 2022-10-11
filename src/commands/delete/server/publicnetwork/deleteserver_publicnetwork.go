package publicnetwork

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "delete server-public-network"

// DeleteServerPublicNetworkCmd is the command for creating a server.
var DeleteServerPublicNetworkCmd = &cobra.Command{
	Use:          "server-public-network SERVER_ID PUBLIC_NETWORK_ID",
	Short:        "Remove a server from a public network.",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	Long: `Remove a server from a public network.

Requires two IDs passed as arguments. First one being the server id and second being the public network id. `,
	Example: `# remove a server from a public network 
pnapctl delete server-public-network <SERVER_ID> <PUBLIC_NETWORK_ID>
`,
	RunE: func(_ *cobra.Command, args []string) error {
		return deletePublicNetworkFromServer(args[0], args[1])
	},
}

func deletePublicNetworkFromServer(serverId, publicNetworkId string) error {
	result, httpResponse, err := bmcapi.Client.ServerPublicNetworkDelete(serverId, publicNetworkId)
	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		fmt.Println(result)
		return nil
	}
}
