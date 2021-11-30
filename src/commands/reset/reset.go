package reset

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/reset/server"
)

var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the resource to original state.",
	Long: `Reset the resource to the same state as it was originally created.
NOTE: Any data on the resource will be lost.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	ResetCmd.AddCommand(server.ResetServerCmd)
}
