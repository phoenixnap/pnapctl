package ip_blocks

import (
	"phoenixnap.com/pnapctl/common/client/ip"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var tags []string
var Full bool

func init() {
	utils.SetupOutputFlag(GetIpBlockCmd)
	utils.SetupFullFlag(GetIpBlockCmd, &Full, "ip-block")
	GetIpBlockCmd.PersistentFlags().StringArrayVar(&tags, "tag", nil, "Filter by tag")
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
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		if len(args) >= 1 {
			return getIpBlockById(args[0])
		}
		return getIpBlocks()
	},
}

func getIpBlocks() error {
	log.Info().Msg("Retrieving list of Ip Blocks...")

	ipBlocks, err := ip.Client.IpBlocksGet(tags)

	if err != nil {
		return err
	} else {
		return printer.PrintIpBlockListResponse(ipBlocks, Full)
	}
}

func getIpBlockById(ipBlockId string) error {
	log.Info().Msgf("Retrieving Ip Block with ID [%s].", ipBlockId)

	ipBlock, err := ip.Client.IpBlocksGetById(ipBlockId)

	if err != nil {
		return err
	} else {
		return printer.PrintIpBlockResponse(ipBlock, Full)
	}
}
