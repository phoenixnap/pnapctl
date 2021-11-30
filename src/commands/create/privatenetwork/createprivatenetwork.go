package privatenetwork

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
	"phoenixnap.com/pnapctl/common/printer"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "create private-network"

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
		privateNetworkCreate, err := networkmodels.CreatePrivateNetworkCreateFromFile(Filename, commandName)

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
	},
}

func init() {
	CreatePrivateNetworkCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreatePrivateNetworkCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreatePrivateNetworkCmd.MarkFlagRequired("filename")
}
