package reservation

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "auto-renew enable reservation"

var AutoRenewEnableReservationCmd = &cobra.Command{
	Use:          "reservation [RESERVATION_ID]",
	Short:        "Enable auto-renew for a reservation",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	Long: `Enable auto-renew for a reservation.
	
// ADD FURTHER NOTES`,
	Example: `
# Enable auto-renew for a specific reservation
pnapctl auto-renew enable reservation [RESERVATION_ID]`,
	RunE: enableAutoRenewForReservation,
}

func enableAutoRenewForReservation(cmd *cobra.Command, args []string) error {
	// TODO: IMPLEMENT
	return nil
}

func init() {
	utils.SetupOutputFlag(AutoRenewEnableReservationCmd)
}
