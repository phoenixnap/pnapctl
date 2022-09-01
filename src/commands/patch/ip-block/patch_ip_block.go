package ipblock

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/ip"
	"phoenixnap.com/pnapctl/common/models/ipmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var Filename string

const commandName = "patch ip-block"

var PatchIpBlockCmd = &cobra.Command{
	Use:          "ip-block IP_BLOCK_ID",
	Short:        "Updates a specific ip-block.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Patch an existing ip-block.

Requires a file (yaml or json) containing the information needed to update the ip-block.`,
	Example: `# Update an existing ip-block`,

	RunE: func(cmd *cobra.Command, args []string) error {

		ipBlockPatch, err := ipmodels.PatchIpBlockRequestFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		response, httpResponse, err := ip.Client.IpBlocksIpBlockIdPatch(args[0], *ipBlockPatch)
		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintIpBlockResponse(response, commandName)
		}
	},
}

func init() {
	PatchIpBlockCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	PatchIpBlockCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for patch")
	PatchIpBlockCmd.MarkFlagRequired("filename")
}
