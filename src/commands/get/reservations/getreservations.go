package reservations

import (
	netHttp "net/http"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	qp "phoenixnap.com/pnapctl/common/models/queryparams/billing"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "get reservations"

var ID *string

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
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			ID = &args[0]
			return getReservations(ID)
		}
		return getReservations(nil)
	},
}

func getReservations(reservationId *string) error {
	var (
		httpResponse *netHttp.Response
		err          error
		reservation  *billingapi.Reservation
		reservations []billingapi.Reservation
	)

	queryParams, err := qp.NewReservationsGetQueryParams(productCategory)

	if err != nil {
		return err
	} else if reservationId == nil {
		reservations, httpResponse, err = billing.Client.ReservationsGet(*queryParams)
	} else {
		reservation, httpResponse, err = billing.Client.ReservationGetById(*reservationId)
	}

	generatedError := utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else if reservationId == nil {
		return printer.PrintReservationListResponse(reservations, Full, commandName)
	} else {
		return printer.PrintReservationResponse(reservation, Full, commandName)
	}
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
