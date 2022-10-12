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

func init() {
	utils.SetupOutputFlag(UpdatePrivateNetworkCmd)
	utils.SetupFilenameFlag(UpdatePrivateNetworkCmd, &Filename, utils.UPDATING)
}

// UpdatePrivateNetworkCmd is the command for creating a private network.
var UpdatePrivateNetworkCmd = &cobra.Command{
	Use:          "private-network PRIVATE_NETWORK_ID",
	Short:        "Update a private network.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Update a private network.

Requires a file (yaml or json) containing the information needed to modify the private-network.`,
	Example: `# Update a private network as per privateNetworkUpdate.yaml
pnapctl update private-network <PRIVATE_NETWORK_ID> --filename <FILENAME> [--output <OUTPUT_TYPE>]

# privateNetworkUpdate.yaml
name: Example CLI Network Updated,
description: Example CLI Network (Updated Description),
locationDefault: true`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return updatePrivateNetwork(args[0])
	},
}

func updatePrivateNetwork(id string) error {
	privateNetworkUpdate, err := models.CreateRequestFromFile[networkapi.PrivateNetworkModify](Filename)

	if err != nil {
		return err
	}

	// update the private network
	response, err := networks.Client.PrivateNetworkPut(id, *privateNetworkUpdate)

	if err != nil {
		return err
	} else {
		return printer.PrintPrivateNetworkResponse(response)
	}
}
