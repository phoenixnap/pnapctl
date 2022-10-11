package privatenetwork

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "update private-network"

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
	RunE: func(_ *cobra.Command, args []string) error {
		return updatePrivateNetwork(args[0])
	},
}

func updatePrivateNetwork(id string) error {
	privateNetworkUpdate, err := models.CreateRequestFromFile[networkapi.PrivateNetworkModify](Filename, commandName)

	if err != nil {
		return err
	}

	// update the private network
	response, httpResponse, err := networks.Client.PrivateNetworkPut(id, *privateNetworkUpdate)

	if httpResponse != nil && httpResponse.StatusCode != 200 {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	} else if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else {
		return printer.PrintPrivateNetworkResponse(response, commandName)
	}
}
