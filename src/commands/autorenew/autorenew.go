package autorenew

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/autorenew/reservation"
)

var AutoRenewCmd = &cobra.Command{
	Use:   "auto-renew",
	Short: "Modify auto-renew for a resource.",
	Long:  `Modify auto-renew for a resource.`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	AutoRenewCmd.AddCommand(reservation.AutoRenewReservationCmd)
}
