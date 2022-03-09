package ip_blocks

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/ip"
	"phoenixnap.com/pnapctl/common/models/ipmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "create ip-block"

// CreateIpBlockCmd is the command for creating an ssh key.
var CreateIpBlockCmd = &cobra.Command{
	Use:          "ip-block",
	Short:        "Create a new ip-block.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a new ip-block.

Requires a file (yaml or json) containing the information needed to create the ip-block.`,
	Example: `# Create a new ip-block as described in ipblockcreate.yaml
pnapctl create ip-block --filename <FILE_PATH> [--output <OUTPUT_TYPE>]

# ipblockreate.yaml
cidrBlockSize: /28
location: PHX`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ipBlockCreate, err := ipmodels.CreateIpBlockRequestFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		// Create the ssh key
		response, httpResponse, err := ip.Client.IpBlockPost(*ipBlockCreate)
		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintIpBlockResponse(response, commandName)
		}
	},
}

func init() {
	CreateIpBlockCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreateIpBlockCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreateIpBlockCmd.MarkFlagRequired("filename")
}
