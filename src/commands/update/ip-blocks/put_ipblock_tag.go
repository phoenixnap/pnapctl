package ip_blocks

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/ip"
	"phoenixnap.com/pnapctl/common/models/ipmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var Filename string

const commandName = "put ip-block tag"

var PutIpBlockTagCmd = &cobra.Command{
	Use:          "ip-block IP_BLOCK_ID",
	Short:        "Updates a specific ip-block's tag.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Update an existing ip-block's tag.
	
	`,
	Example: `# Put a tag on an existing ip-block with request body as described in ipblockputtag.yaml
	pnapctl put ip-block tag <IP_BLOCK_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]
	
	# ipblockputtag.yaml
	name: ip block tag name
	value: ip block tag value
	`,

	RunE: func(cmd *cobra.Command, args []string) error {

		ipBlockPutTag, err := ipmodels.PutIpBlockTagRequestFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		response, httpResponse, err := ip.Client.IpBlocksIpBlockIdTagsPut(args[0], ipBlockPutTag)

		if err != nil {
			return err
		}

		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintIpBlockResponse(response, commandName)
		}
	},
}
