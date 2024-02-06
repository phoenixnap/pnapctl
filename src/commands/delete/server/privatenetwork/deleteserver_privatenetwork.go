package privatenetwork

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

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
		cmdname.SetCommandName(cmd)
		return deletePrivateNetworkFromServer(args[0], args[1])
	},
}

func deletePrivateNetworkFromServer(serverId, privateNetworkId string) error {
	log.Info().Msgf("Removing Server with ID [%s] from Private Network with ID [%s].", serverId, privateNetworkId)

	result, err := bmcapi.Client.ServerPrivateNetworkDelete(serverId, privateNetworkId)
	if err != nil {
		return err
	} else {
		fmt.Println(result)
		return nil
	}
}
