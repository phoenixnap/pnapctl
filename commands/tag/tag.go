package tag

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/tag/server"
)

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Tag a resource.",
	Long:  `Tag a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	TagCmd.AddCommand(server.TagServerCmd)
}
