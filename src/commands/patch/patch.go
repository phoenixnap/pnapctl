package patch

import (
	"os"

	"github.com/spf13/cobra"
	ip_block "phoenixnap.com/pnapctl/commands/patch/ip-block"
	"phoenixnap.com/pnapctl/commands/patch/publicnetwork"
	"phoenixnap.com/pnapctl/commands/patch/server"
	"phoenixnap.com/pnapctl/commands/patch/tag"
)

var PatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Modify a resource.",
	Long:  `Modify a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	PatchCmd.AddCommand(tag.PatchTagCmd)
	PatchCmd.AddCommand(server.PatchServerCmd)
	PatchCmd.AddCommand(publicnetwork.PatchPublicNetworkCmd)
	PatchCmd.AddCommand(ip_block.PatchIpBlockCmd)
}
