package privatenetwork

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "create server-private-network"

func init() {
	utils.SetupOutputFlag(CreateServerPrivateNetworkCmd)
	utils.SetupFilenameFlag(CreateServerPrivateNetworkCmd, &Filename, utils.CREATION)
}

// CreateServerPrivateNetworkCmd is the command for creating a server.
var CreateServerPrivateNetworkCmd = &cobra.Command{
	Use:          "server-private-network SERVER_ID",
	Short:        "Create a new private network for server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Create a new private network for server.

Requires a file (yaml or json) containing the information needed to create the server private network.`,
	Example: `# Add a server to a private network as defined in serverCreatePrivateNetwork.yaml
pnapctl create server-private-network <SERVER_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# serverCreatePrivateNetwork.yaml
id: 5ff5cc9bc1acf144d9106233
ips: 
  - 10.0.0.1
  - 10.0.0.2
dhcp: false
statusDescription: in-progress
`,

	RunE: func(_ *cobra.Command, args []string) error {
		return createPrivateNetworkForServer(args[0])
	},
}

func createPrivateNetworkForServer(id string) error {
	serverPrivateNetwork, err := models.CreateRequestFromFile[bmcapisdk.ServerPrivateNetwork](Filename, commandName)

	if err != nil {
		return err
	}

	// Create the server private network
	response, httpResponse, err := bmcapi.Client.ServerPrivateNetworkPost(id, *serverPrivateNetwork)

	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintServerPrivateNetwork(response, commandName)
	}
}
