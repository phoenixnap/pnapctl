package shutdown

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/shutdown/server"
)

var ShutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Perform a soft shutdown on the resource.",
	Long:  `Perform a soft shutdown on the resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	ShutdownCmd.AddCommand(server.ShutdownCmd)
}
