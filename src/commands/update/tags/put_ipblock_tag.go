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
	Short:        "Updates an ip block's tags.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Update an existing ip-block's tag.
	
	Requires a file (yaml or json) containing the information needed to update the ip-block's tags.
	`,
	Example: `# Put a tag on an existing ip-block with request body as described in ipblockputtag.yaml
	pnapctl put ip-block tag <IP_BLOCK_ID> --filename <FILE_PATH> [--output <OUTPUT_TYPE>]
	
	# ipblockputtag.yaml
	---
	tags:
		- name: ip block tag name
  		  value: ip block tag value
	`,

	RunE: func(cmd *cobra.Command, args []string) error {

		ipBlockPutTag, err := ipmodels.PutIpBlockTagRequestFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		response, httpResponse, err := ip.Client.IpBlocksIpBlockIdTagsPut(args[0], ipBlockPutTag)

		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintIpBlockResponse(response, commandName)
		}
	},
}

func init() {
	PutIpBlockTagCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	PutIpBlockTagCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	PutIpBlockTagCmd.MarkFlagRequired("filename")
}
