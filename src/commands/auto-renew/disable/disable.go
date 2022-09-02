package disable

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/auto-renew/disable/reservation"
)

var DisableAutoRenewCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable auto-renew for a resource.",
	Long:  `Disable auto-renew for a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	DisableAutoRenewCmd.AddCommand(reservation.AutoRenewDisableReservationCmd)
}
