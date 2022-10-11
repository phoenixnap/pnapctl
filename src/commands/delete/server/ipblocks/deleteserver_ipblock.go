package ipblocks

import (
	"fmt"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

func init() {
	utils.SetupFilenameFlag(DeleteServerIpBlockCmd, &Filename, utils.DELETION)
}

// DeleteServerIpBlockCmd is the command for deleting a server ip block.
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
		cmdname.SetCommandName(cmd)
		return deleteIpBlockFromServer(args[0], args[1])
	},
}

func deleteIpBlockFromServer(serverId, ipBlockId string) error {
	relinquishIpBlockRequest, err := models.CreateRequestFromFile[bmcapisdk.RelinquishIpBlock](Filename)

	if err != nil {
		return err
	}

	result, httpResponse, err := bmcapi.Client.ServerIpBlockDelete(serverId, ipBlockId, *relinquishIpBlockRequest)
	var generatedError = utils.CheckForErrors(httpResponse, err)

	if *generatedError != nil {
		return *generatedError
	} else {
		fmt.Println(result)
		return nil
	}
}
