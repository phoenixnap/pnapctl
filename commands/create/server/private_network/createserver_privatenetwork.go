package private_network

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/bmcapimodels"
	"phoenixnap.com/pnap-cli/common/printer"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "create server-private-network"

// CreateServerCmd is the command for creating a server.
var CreateServerPrivateNetworkCmd = &cobra.Command{
	Use:          "server-private-network",
	Short:        "Create a new private network for server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Create a new private network forserver.

Requires a file (yaml or json) containing the information needed to create the server.`,
	Example: `# create a new server as described in server.yaml
pnapctl create server --filename ./createPrivateNetwork.yaml

# createPrivateNetwork.yaml
id: 5ff5cc9bc1acf144d9106233,
ips: [],
dhcp: false,
statusDescription": in-progress
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		serverPrivateNetwork, err := bmcapimodels.CreateServerPrivateNetworkFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		// Create the server
		response, httpResponse, err := bmcapi.Client.ServerPrivateNetworkPost(args[0], *serverPrivateNetwork)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else if httpResponse.StatusCode == 200 {
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
