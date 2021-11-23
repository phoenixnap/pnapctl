package reserve

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/commands/reserve/server"
)

var ReserveCmd = &cobra.Command{
	Use:   "reserve",
	Short: "Reserve the resource for future use.",
	Long:  `Reserve the resource to be used later on.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	ReserveCmd.AddCommand(server.ReserveServerCmd)
}
