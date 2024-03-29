package ipblocks

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

func init() {
	utils.SetupOutputFlag(CreateServerIpBlockCmd)
	utils.SetupFilenameFlag(CreateServerIpBlockCmd, &Filename, utils.CREATION)
}

// CreateServerIpBlockCmd is the command for creating a server.
var CreateServerIpBlockCmd = &cobra.Command{
	Use:          "server-ip-block SERVER_ID",
	Short:        "Create a new ip-block for server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Create a new ip-block for server.

Requires a file (yaml or json) containing the information needed to create the server ip-block.`,
	Example: `# Add an ip-block to a server defined in servercreateipblock.yaml
pnapctl create server-ip-block <SERVER_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# servercreateipblock.yaml
id: 5ff5cc9bc1acf144d9106233
vlanId: 11`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return createIpBlockForServer(args[0])
	},
}

func createIpBlockForServer(id string) error {
	log.Info().Msgf("Creating new Ip Block for Server with ID [%s].", id)

	serverIpBlock, err := models.CreateRequestFromFile[bmcapisdk.ServerIpBlock](Filename)

	if err != nil {
		return err
	}

	// Create the server ip block
	response, err := bmcapi.Client.ServerIpBlockPost(id, *serverIpBlock)

	if err != nil {
		return err
	} else {
		return printer.PrintServerIpBlock(response)
	}
}
