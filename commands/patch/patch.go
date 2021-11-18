package patch

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/patch/tag"
)

var PatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch a resource.",
	Long:  `Patch a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	PatchCmd.AddCommand(tag.PatchTagCmd)
}
