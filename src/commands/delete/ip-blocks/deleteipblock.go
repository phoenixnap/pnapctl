package ip_blocks

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/ip"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName = "delete ip-block"

var DeleteIpBlockCmd = &cobra.Command{
	Use:          "ip-block IP_BLOCK_ID",
	Short:        "Deletes a specific ip-block.",
	Long:         "Deletes a specific ip-block.",
	Example:      `pnapctl delete ip-block <IP_BLOCK_ID>`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(_ *cobra.Command, args []string) error {
		return deleteIpBlock(args[0])
	},
}

func deleteIpBlock(id string) error {
	result, httpResponse, err := ip.Client.IpBlocksIpBlockIdDelete(id)
	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		fmt.Println(result.Result, result.IpBlockId)
		return nil
	}
}
