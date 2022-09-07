package ip_blocks

import (
	"os"

	"github.com/spf13/cobra"
	tags "phoenixnap.com/pnapctl/commands/update/ip-blocks/tags"
)

var UpdateIpBlockCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an ip-block.",
	Long:  `Update an ip-block.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	UpdateIpBlockCmd.AddCommand(tags.PutIpBlockTagCmd)
}
