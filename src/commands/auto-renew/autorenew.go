package autorenew

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/auto-renew/disable"
	"phoenixnap.com/pnapctl/commands/auto-renew/enable"
)

var AutoRenewCmd = &cobra.Command{
	Use:   "auto-renew",
	Short: "Modify auto-renew for a resource.",
	Long:  `Modify auto-renew for a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	AutoRenewCmd.AddCommand(disable.DisableAutoRenewCmd)
	AutoRenewCmd.AddCommand(enable.EnableAutoRenewCmd)
}
