package reservation

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "auto-renew disable reservation"

var AutoRenewDisableReservationCmd = &cobra.Command{
	Use:          "reservation [RESERVATION_ID]",
	Short:        "Disable auto-renew for a reservation",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	Long: `Disable auto-renew for a reservation.
	
// ADD FURTHER NOTES`,
	Example: `
# Disable auto-renew for a specific reservation
pnapctl auto-renew disable reservation [RESERVATION_ID]`,
	RunE: disableAutoRenewForReservation,
}

func disableAutoRenewForReservation(cmd *cobra.Command, args []string) error {
	// TODO: IMPLEMENT
	return nil
}

func init() {
	utils.SetupOutputFlag(AutoRenewDisableReservationCmd)
}
