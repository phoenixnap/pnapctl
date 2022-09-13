package publicnetwork

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "create server-public-network"

var CreateServerPublicNetworkCmd = &cobra.Command{
	Use:          "server-public-network SERVER_ID",
	Short:        "Create a new public network for server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Create a new public network for server.

Requires a file (yaml or json) containing the information needed to create the server public network.`,
	Example: `# Add a server to a public network as defined in serverCreatePublicNetwork.yaml
pnapctl create server-public-network <SERVER_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# serverCreatePublicNetwork.yaml
id: 5ff5cc9bc1acf144d9106233
ips: 
  - 10.0.0.1
  - 10.0.0.2
statusDescription: in-progress
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		serverPublicNetwork, err := servermodels.CreateServerPublicNetworkFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		// Create the server private network
		response, httpResponse, err := bmcapi.Client.ServerPublicNetworkPost(args[0], *serverPublicNetwork)

		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintServerPublicNetwork(response, commandName)
		}
	},
}

func init() {
	CreateServerPublicNetworkCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreateServerPublicNetworkCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreateServerPublicNetworkCmd.MarkFlagRequired("filename")
}