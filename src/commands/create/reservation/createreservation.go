package reservation

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	Full     bool
	Filename string
)

func init() {
	utils.SetupFullFlag(CreateReservationCmd, &Full, "reservation")
	utils.SetupOutputFlag(CreateReservationCmd)
	utils.SetupFilenameFlag(CreateReservationCmd, &Filename, utils.CREATION)
}

var CreateReservationCmd = &cobra.Command{
	Use:          "reservation [RESERVATION_ID]",
	Short:        "Create a new reservation.",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(0),
	Long: `Create a new reservation.
	
Requires a file (yaml or json) containing the information needed to create the reservation.`,
	Example: `
# Create a specific reservation
pnapctl create reservation <RESERVATION_ID> --filename=<FILENAME>

# reservationCreate.yaml
sku: "skuCode"`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return createReservation()
	},
}

func createReservation() error {
	reservationCreate, err := models.CreateRequestFromFile[billingapi.ReservationRequest](Filename)

	if err != nil {
		return err
	}

	// Create the server
	response, httpResponse, err := billing.Client.ReservationsPost(*reservationCreate)
	if err := utils.CheckErrs(httpResponse, err); err != nil {
		return err
	} else {
		return printer.PrintReservationResponse(response, Full)
	}
}
