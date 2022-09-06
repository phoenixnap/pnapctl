package reservation

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "create reservation"

var CreateReservationCmd = &cobra.Command{
	Use:          "reservation [RESERVATION_ID]",
	Short:        "Create a reservation",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(0),
	Long: `Create a reservation.
	
Requires a file (yaml or json) containing the information needed to create the reservation.`,
	Example: `
# Create a specific reservation
pnapctl create reservation <RESERVATION_ID> --filename=<FILENAME>

# reservationCreate.yaml
sku: "skuCode"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		reservationCreate, err := billingmodels.CreateReservationRequestFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		// Create the server
		response, httpResponse, err := billing.Client.ReservationsPost(*reservationCreate)
		generatedError := utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintReservationResponse(response, Full, commandName)
		}
	},
}

var (
	Full     bool
	Filename string
)

func init() {
	utils.SetupFullFlag(CreateReservationCmd, &Full, "reservation")
	utils.SetupOutputFlag(CreateReservationCmd)

	CreateReservationCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreateReservationCmd.MarkFlagRequired("filename")
}
