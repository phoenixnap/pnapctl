package enable

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/auto-renew/enable/reservation"
)

var EnableAutoRenewCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable auto-renew for a resource.",
	Long:  `Enable auto-renew for a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	EnableAutoRenewCmd.AddCommand(reservation.AutoRenewEnableReservationCmd)
}
