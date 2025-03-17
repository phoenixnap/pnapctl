package ip_blocks

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/ip"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var DeleteIpBlockCmd = &cobra.Command{
	Use:          "ip-block IP_BLOCK_ID",
	Short:        "Deletes a specific ip-block.",
	Long:         "Deletes a specific ip-block.",
	Example:      `pnapctl delete ip-block <IP_BLOCK_ID>`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return deleteIpBlock(args[0])
	},
}

func deleteIpBlock(id string) error {
	log.Info().Msgf("Deleting Ip Block with ID [%s].", id)

	result, err := ip.Client.IpBlocksIpBlockIdDelete(id)
	if err != nil {
		return err
	} else {
		fmt.Println(*result.Result, *result.IpBlockId)
		return nil
	}
}
