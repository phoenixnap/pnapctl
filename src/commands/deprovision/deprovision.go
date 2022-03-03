package deprovision

import (
	"os"
	"phoenixnap.com/pnapctl/commands/deprovision/server"

	"github.com/spf13/cobra"
)

var DeprovisionCmd = &cobra.Command{
	Use:   "deprovision",
	Short: "Deprovision a resource.",
	Long:  `Deprovision a resource`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	DeprovisionCmd.AddCommand(server.DeprovisionServerCmd)
}
