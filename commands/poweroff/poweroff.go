package poweroff

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/poweroff/server"
)

var PowerOffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Perform a hard shutdown on the resource.",
	Long:  `Perform a hard shutdown on the resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	PowerOffCmd.AddCommand(server.PowerOffServerCmd)
}
