package patch

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/patch/server"
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
	PatchCmd.AddCommand(server.PatchServerCmd)
}
