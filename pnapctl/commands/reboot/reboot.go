package reboot

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/commands/reboot/server"
)

var RebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Perform a soft reboot on resource.",
	Long:  `Perform a soft reboot on resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	RebootCmd.AddCommand(server.RebootCmd)
}
