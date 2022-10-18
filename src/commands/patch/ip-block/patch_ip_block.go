package ipblock

import (
	"github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/ip"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var Filename string

var Full bool

const commandName = "patch ip-block"

var PatchIpBlockCmd = &cobra.Command{
	Use:          "ip-block IP_BLOCK_ID",
	Short:        "Updates a specific ip-block.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Patch an existing ip-block.

Requires a file (yaml or json) containing the information needed to update the ip-block.`,
	Example: `# Update an existing ip-block with request body as described in ipblockpatch.yaml
	pnapctl patch ip-block <IP_BLOCK_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]
	
	# ipblockpatch.yaml
	description: ip block description`,

	RunE: func(cmd *cobra.Command, args []string) error {

		ipBlockPatch, err := models.CreateRequestFromFile[ipapi.IpBlockPatch](Filename, commandName)

		if err != nil {
			return err
		}

		response, httpResponse, err := ip.Client.IpBlocksIpBlockIdPatch(args[0], *ipBlockPatch)

		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintIpBlockResponse(response, Full, commandName)
		}
	},
}

func init() {
	PatchIpBlockCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	PatchIpBlockCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for patch")
	PatchIpBlockCmd.MarkFlagRequired("filename")
	PatchIpBlockCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all ip-block details")
}
