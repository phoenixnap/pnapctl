package reservations

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	qp "phoenixnap.com/pnapctl/common/models/queryparams/billing"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "get reservations"

var (
	Full            bool
	productCategory string
)

func init() {
	utils.SetupOutputFlag(GetReservationsCmd)
	utils.SetupFullFlag(GetReservationsCmd, &Full, "reservation")

	GetReservationsCmd.Flags().StringVar(&productCategory, "category", "", "Product category to filter reservations by.")
}

var GetReservationsCmd = &cobra.Command{
	Use:          "reservation [RESERVATION_ID]",
	Short:        "Retrieve one or all reservations",
	Aliases:      []string{"reservations"},
	SilenceUsage: true,
	Args:         cobra.RangeArgs(0, 1),
	Long:         `Retrieve one or all reservations.`,
	Example: `
# Retrieve all reservations
pnapctl get reservations [--category=<CATEGORY>] [--full] [--output=<OUTPUT_TYPE>]

# Retrieve a specific reservation
pnapctl get reservation <RESERVATION_ID> [--full] [--output=<OUTPUT_TYPE>]`,
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) >= 1 {
			return getReservationById(args[0])
		}
		return getReservations()
	},
}

func getReservations() error {
	queryParams, err := qp.NewReservationsGetQueryParams(productCategory)

	if err != nil {
		return err
	}

	reservations, httpResponse, err := billing.Client.ReservationsGet(*queryParams)

	generatedError := utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintReservationListResponse(reservations, Full, commandName)
	}
}

func getReservationById(reservationId string) error {
	reservation, httpResponse, err := billing.Client.ReservationGetById(reservationId)

	generatedError := utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintReservationResponse(reservation, Full, commandName)
	}
}
