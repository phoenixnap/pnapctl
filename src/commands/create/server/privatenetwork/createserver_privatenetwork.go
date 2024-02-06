package privatenetwork

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var (
	Filename string
	force bool
)

func init() {
	utils.SetupOutputFlag(CreateServerPrivateNetworkCmd)
	utils.SetupFilenameFlag(CreateServerPrivateNetworkCmd, &Filename, utils.CREATION)

	CreateServerPrivateNetworkCmd.Flags().BoolVar(&force, "force", false, "Controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. Defaults to false.")
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
pnapctl create server-private-network <SERVER_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>] [--force=false]

# serverCreatePrivateNetwork.yaml
id: 5ff5cc9bc1acf144d9106233
ips: 
  - 10.0.0.1
  - 10.0.0.2
dhcp: false
statusDescription: in-progress
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return createPrivateNetworkForServer(args[0])
	},
}

func createPrivateNetworkForServer(id string) error {
	log.Info().Msgf("Creating new Private Network for Server with ID [%s].", id)

	serverPrivateNetwork, err := models.CreateRequestFromFile[bmcapisdk.ServerPrivateNetwork](Filename)

	if err != nil {
		return err
	}

	// Create the server private network
	response, err := bmcapi.Client.ServerPrivateNetworkPost(id, *serverPrivateNetwork, force)

	if err != nil {
		return err
	} else {
		return printer.PrintServerPrivateNetwork(response)
	}
}
