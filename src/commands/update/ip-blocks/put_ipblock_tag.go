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

	RunE: func(cmd *cobra.Command, args []string) error {

		ipBlockPutTag, err := ipmodels.PutIpBlockTagRequestFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		response, httpResponse, err := ip.Client.IpBlocksIpBlockIdTagsPut(args[0], *ipBlockPutTag)

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
