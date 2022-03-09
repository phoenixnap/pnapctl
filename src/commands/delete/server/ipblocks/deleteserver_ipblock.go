package ipblocks

import (
	"fmt"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "delete server-ip-block"

// DeleteServerIpBlockCmd is the command for creating a server.
var DeleteServerIpBlockCmd = &cobra.Command{
	Use:          "server-ip-block SERVER_ID IP_BLOCK_ID",
	Short:        "Remove an ip-block from a server.",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	Long: `Remove an ip-block from a server.

Requires two IDs passed as arguments and a file (yaml or json) containing the information needed. First one being the server id and second being the ip-block id. `,
	Example: `# Remove an ip-block from a server. 
pnapctl delete server-ip-block <SERVER_ID> <IP_BLOCK_ID> --filename <FILE_PATH>

# serveripblockdelete.yaml
deleteIpBlocks: false`,
	RunE: func(cmd *cobra.Command, args []string) error {
		relinquishIpBlockRequest, err := servermodels.CreateRelinquishIpBlockRequestFromFile(Filename, commandName)
		result, httpResponse, err := bmcapi.Client.ServerIpBlockDelete(args[0], args[1], *relinquishIpBlockRequest)
		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			fmt.Println(result)
			return nil
		}
	},
}

func init() {
	DeleteServerIpBlockCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	DeleteServerIpBlockCmd.MarkFlagRequired("filename")
}
