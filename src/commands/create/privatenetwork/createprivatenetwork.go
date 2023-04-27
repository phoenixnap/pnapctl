package privatenetwork

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var Force bool

func init() {
	utils.SetupOutputFlag(CreatePrivateNetworkCmd)
	utils.SetupFilenameFlag(CreatePrivateNetworkCmd, &Filename, utils.CREATION)
	CreatePrivateNetworkCmd.PersistentFlags().BoolVar(&Force, "force", false, "Controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. Defaults to false.")
}

// CreatePrivateNetworkCmd is the command for creating a private-network.
var CreatePrivateNetworkCmd = &cobra.Command{
	Use:          "private-network",
	Short:        "Create a new private network.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a new private-network.

Requires a file (yaml or json) containing the information needed to create the private network.`,
	Example: `# Create a new private network as per privateNetworkCreate.yaml
pnapctl create private-network --filename <FILE_PATH> [--output <OUTPUT_TYPE>] [--force]

# privateNetworkCreate.yaml
name: Example CLI Network,
location: PHX,
locationDefault: false,
description: Example CLI Network,
cidr: 10.0.0.0/24`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return createPrivateNetwork()
	},
}

func createPrivateNetwork() error {
	privateNetworkCreate, err := models.CreateRequestFromFile[networkapi.PrivateNetworkCreate](Filename)

	if err != nil {
		return err
	}

	// Create the private network
	response, err := networks.Client.PrivateNetworksPost(*privateNetworkCreate, Force)

	if err != nil {
		return err
	} else {
		return printer.PrintPrivateNetworkResponse(response)
	}
}
