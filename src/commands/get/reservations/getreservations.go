package reservations

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "get reservations"

var GetReservationsCmd = &cobra.Command{
	Use:          "reservation [RESERVATION_ID]",
	Short:        "Retrieve one or all reservations",
	Aliases:      []string{"reservations"},
	SilenceUsage: true,
	Args:         cobra.RangeArgs(0, 1),
	Long: `Retrieve one or all reservations.
	
// ADD FURTHER NOTES`,
	Example: `
# Retrieve all reservation
pnapctl get reservations

# Retrieve a specific reservation
pnapctl get reservation [RESERVATION_ID]`,
	RunE: getReservations,
}

func getReservations(cmd *cobra.Command, args []string) error {
	queryParams, err := billingmodels.NewReservationsGetQueryParams(productCategory)

	if err != nil {
		return err
	}

	return utils.UseIdFor[billingapi.Reservation](args).
		IfPresent(
			billing.Client.ReservationGetById,
			utils.UsingFull(Full, printer.PrintReservationResponse),
		).
		Else(
			utils.DoRequestWith(billing.Client.ReservationsGet, *queryParams),
			utils.UsingFull(Full, printer.PrintReservationListResponse),
		).
		Execute(commandName)
}

var (
	Full            bool
	productCategory string
)

func init() {
	utils.SetupOutputFlag(GetReservationsCmd)
	utils.SetupFullFlag(GetReservationsCmd, &Full, "reservation")

	GetReservationsCmd.Flags().StringVar(&productCategory, "category", "", "Product category to filter reservations by.")
}
