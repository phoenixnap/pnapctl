package ip_blocks

import (
	"phoenixnap.com/pnapctl/common/client/ip"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName string = "get ip-blocks"

var tags []string
var Full bool

func init() {
	GetIpBlockCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	GetIpBlockCmd.PersistentFlags().StringArrayVar(&tags, "tag", nil, "Filter by tag")
	GetIpBlockCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all ip-block details")
}

var GetIpBlockCmd = &cobra.Command{
	Use:          "ip-block [IP_BLOCK_ID]",
	Short:        "Retrieve one or all ip-blocks for your account.",
	Aliases:      []string{"ip-blocks"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all ip-blocks for your account.

Prints all information about the ip-blocks owned by your account.
By default, the data is printed in table format.

To print a specific ip-block, an ip-block ID needs to be passed as an argument.`,
	Example: `
# List all ip-blocks.
pnapctl get ip-blocks [--output <OUTPUT_TYPE>]

# List a specific ip-block.
pnapctl get ip-block <IP_BLOCK_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) >= 1 {
			return getIpBlockById(args[0])
		}
		return getIpBlocks()
	},
}

func getIpBlocks() error {
	ipBlocks, httpResponse, err := ip.Client.IpBlocksGet(tags)

	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintIpBlockListResponse(ipBlocks, Full, commandName)
	}
}

func getIpBlockById(ipBlockId string) error {
	ipBlock, httpResponse, err := ip.Client.IpBlocksGetById(ipBlockId)

	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintIpBlockResponse(ipBlock, Full, commandName)
	}
}
