package publicnetwork

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
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
	utils.SetupOutputFlag(CreateServerPublicNetworkCmd)
	utils.SetupFilenameFlag(CreateServerPublicNetworkCmd, &Filename, utils.CREATION)
}

var CreateServerPublicNetworkCmd = &cobra.Command{
	Use:          "server-public-network SERVER_ID",
	Short:        "Create a new public network for server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Create a new public network for server.

Requires a file (yaml or json) containing the information needed to create the server public network.`,
	Example: `# Add a server to a public network as defined in serverCreatePublicNetwork.yaml
pnapctl create server-public-network <SERVER_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# serverCreatePublicNetwork.yaml
id: 6322c9ec9da56569d0ca4add
ips: 
  - 10.111.24.25
  - 10.111.24.26
statusDescription: in-progress
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return createPublicNetworkForServer(args[0])
	},
}

func createPublicNetworkForServer(id string) error {
	serverPublicNetwork, err := models.CreateRequestFromFile[bmcapisdk.ServerPublicNetwork](Filename)

	if err != nil {
		return err
	}

	// Create the server private network
	response, err := bmcapi.Client.ServerPublicNetworkPost(id, *serverPublicNetwork)

	if err != nil {
		return err
	} else {
		return printer.PrintServerPublicNetwork(response)
	}
}
