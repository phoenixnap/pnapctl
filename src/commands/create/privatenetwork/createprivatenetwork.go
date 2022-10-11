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

var commandName = "create private-network"

func init() {
	utils.SetupOutputFlag(CreatePrivateNetworkCmd)
	utils.SetupFilenameFlag(CreatePrivateNetworkCmd, &Filename, utils.CREATION)
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
pnapctl create private-network --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# privateNetworkCreate.yaml
name: Example CLI Network,
location: PHX,
locationDefault: false,
description: Example CLI Network,
cidr: 10.0.0.0/24`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return createPrivateNetwork()
	},
}

func createPrivateNetwork() error {
	privateNetworkCreate, err := models.CreateRequestFromFile[networkapi.PrivateNetworkCreate](Filename, commandName)

	if err != nil {
		return err
	}

	// Create the private network
	response, httpResponse, err := networks.Client.PrivateNetworksPost(*privateNetworkCreate)

	if httpResponse != nil && httpResponse.StatusCode != 201 {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	} else if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else {
		return printer.PrintPrivateNetworkResponse(response, commandName)
	}
}
