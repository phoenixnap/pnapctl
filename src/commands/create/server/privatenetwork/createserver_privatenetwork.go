package privatenetwork

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "create server-private-network"

// CreateServerPrivateNetworkCmd is the command for creating a server.
var CreateServerPrivateNetworkCmd = &cobra.Command{
	Use:          "server-private-network SERVER_ID",
	Short:        "Create a new private network for server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Create a new private network for server.

Requires a file (yaml or json) containing the information needed to create the server private network.`,
	Example: `# Add a server to a private network as defined in serverPrivateNetworkCreate.yaml
pnapctl create server-private-network <SERVER_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# serverPrivateNetworkCreate.yaml
id: 5ff5cc9bc1acf144d9106233
ips: 
  - 10.0.0.1
  - 10.0.0.2
dhcp: false
statusDescription: in-progress
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		serverPrivateNetwork, err := bmcapimodels.CreateServerPrivateNetworkFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		// Create the server private network
		response, httpResponse, err := bmcapi.Client.ServerPrivateNetworkPost(args[0], *serverPrivateNetwork)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else if utils.Is2xxSuccessful(httpResponse.StatusCode) {
			return printer.PrintServerPrivateNetwork(response, commandName)
		} else {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}
	},
}

func init() {
	CreateServerPrivateNetworkCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreateServerPrivateNetworkCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreateServerPrivateNetworkCmd.MarkFlagRequired("filename")
}
