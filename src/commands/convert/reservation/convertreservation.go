package reservation

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "convert reservation"

var ConvertReservationCmd = &cobra.Command{
	Use:          "reservation [RESERVATION_ID]",
	Short:        "Convert a reservation",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	Long: `Convert a reservation.
	
// ADD FURTHER NOTES`,
	Example: `
# Convert a specific reservation
pnapctl convert reservation [RESERVATION_ID]`,
	RunE: convertReservation,
}

func convertReservation(cmd *cobra.Command, args []string) error {
	// TODO: IMPLEMENT
	return nil
}

func init() {
	utils.SetupOutputFlag(ConvertReservationCmd)
}
