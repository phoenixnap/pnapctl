package poweron

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/commands/poweron/server"
)

var PowerOnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Power on resource.",
	Long:  `Power on resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	PowerOnCmd.AddCommand(server.PowerOnServerCmd)
}
