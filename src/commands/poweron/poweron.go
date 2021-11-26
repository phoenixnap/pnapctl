package poweron

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/poweron/server"
)

var PowerOnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Power on a resource.",
	Long:  `Power on a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	PowerOnCmd.AddCommand(server.PowerOnServerCmd)
}
