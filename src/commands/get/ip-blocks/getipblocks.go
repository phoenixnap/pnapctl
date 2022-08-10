package ip_blocks

import (
	netHttp "net/http"

	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi"
	"phoenixnap.com/pnapctl/common/client/ip"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName string = "get ip-blocks"

var ID string

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
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			ID = args[0]
			return getIpBlocks(ID)
		}
		return getIpBlocks("")
	},
}

func getIpBlocks(ipBlockId string) error {
	var httpResponse *netHttp.Response
	var err error
	var ipBlock *ipapisdk.IpBlock
	var ipBlocks []ipapisdk.IpBlock

	if ipBlockId == "" {
		ipBlocks, httpResponse, err = ip.Client.IpBlocksGet()
	} else {
		ipBlock, httpResponse, err = ip.Client.IpBlocksGetById(ipBlockId)
	}

	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		if ipBlockId == "" {
			return printer.PrintIpBlockListResponse(ipBlocks, commandName)
		} else {
			return printer.PrintIpBlockResponse(*ipBlock, commandName)
		}
	}
}

func init() {
	GetIpBlockCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
