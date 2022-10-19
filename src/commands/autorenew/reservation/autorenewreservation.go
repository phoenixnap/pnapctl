package reservation

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/autorenew/reservation/disable"
	"phoenixnap.com/pnapctl/commands/autorenew/reservation/enable"
)

var AutoRenewReservationCmd = &cobra.Command{
	Use:   "reservation",
	Short: "autorenew for a resource.",
	Long:  `autorenew for a resource.`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	AutoRenewReservationCmd.AddCommand(disable.AutoRenewDisableReservationCmd)
	AutoRenewReservationCmd.AddCommand(enable.AutoRenewEnableReservationCmd)
}
