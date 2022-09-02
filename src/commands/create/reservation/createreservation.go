package reservation

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "create reservation"

var CreateReservationCmd = &cobra.Command{
	Use:          "reservation [RESERVATION_ID]",
	Short:        "Create a reservation",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(0),
	Long: `Create a reservation.
	
// ADD FURTHER NOTES`,
	Example: `
# Create a specific reservation
pnapctl create reservation [RESERVATION_ID]`,
	RunE: createReservation,
}

func createReservation(cmd *cobra.Command, args []string) error {
	// TODO: IMPLEMENT
	return nil
}

func init() {
	utils.SetupOutputFlag(CreateReservationCmd)
}
