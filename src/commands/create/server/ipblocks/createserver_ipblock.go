package ipblocks

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "create server-ip-block"

// CreateServerIpBlockCmd is the command for creating a server.
var CreateServerIpBlockCmd = &cobra.Command{
	Use:          "server-ip-block SERVER_ID",
	Short:        "Create a new ip-block for server.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Create a new ip-block for server.

Requires a file (yaml or json) containing the information needed to create the server ip-block.`,
	Example: `# Add an ip-block to a server defined in servercreateipblock.yaml
pnapctl create server-ip-block <SERVER_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# servercreateipblock.yaml
id: 5ff5cc9bc1acf144d9106233
vlanId: 11`,
	RunE: func(cmd *cobra.Command, args []string) error {
		serverIpBlock, err := servermodels.CreateServerIpBlockFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		// Create the server private network
		response, httpResponse, err := bmcapi.Client.ServerIpBlockPost(args[0], *serverIpBlock)

		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintServerIpBlock(response, commandName)
		}
	},
}

func init() {
	CreateServerIpBlockCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreateServerIpBlockCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreateServerIpBlockCmd.MarkFlagRequired("filename")
}
