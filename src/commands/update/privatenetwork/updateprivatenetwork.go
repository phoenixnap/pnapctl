package privatenetwork

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/networks"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "update private-network"

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
		privateNetworkUpdate, err := models.CreateRequestFromFile[networkapi.PrivateNetworkModify](Filename, commandName)

		if err != nil {
			return err
		}

		// update the private network
		response, httpResponse, err := networks.Client.PrivateNetworkPut(args[0], *privateNetworkUpdate)

		if httpResponse != nil && httpResponse.StatusCode != 200 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		} else if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else {
			return printer.PrintPrivateNetworkResponse(response, commandName)
		}
	},
}

func init() {
	UpdatePrivateNetworkCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	UpdatePrivateNetworkCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	UpdatePrivateNetworkCmd.MarkFlagRequired("filename")
}
