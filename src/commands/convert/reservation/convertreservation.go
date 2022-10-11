package reservation

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "convert reservation"

var (
	Full     bool
	Filename string
)

func init() {
	utils.SetupFullFlag(ConvertReservationCmd, &Full, "reservation")
	utils.SetupOutputFlag(ConvertReservationCmd)

	ConvertReservationCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	ConvertReservationCmd.MarkFlagRequired("filename")
}

var ConvertReservationCmd = &cobra.Command{
	Use:          "reservation [RESERVATION_ID]",
	Short:        "Convert a reservation",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	Long: `Convert a reservation.
	
Requires a file (yaml or json) containing the information needed to convert the reservation`,
	Example: `
# Convert a specific reservation
pnapctl convert reservation <RESERVATION_ID> --filename=[FILENAME]

# convertReservation.yaml
sku: "SKU_CODE"`,
	RunE: func(_ *cobra.Command, args []string) error {
		return convertReservation(args[0])
	},
}

func convertReservation(id string) error {
	request, err := models.CreateRequestFromFile[billingapi.ReservationRequest](Filename, commandName)
	if err != nil {
		return err
	}

	response, httpResponse, err := billing.Client.ReservationConvert(id, *request)
	generatedError := utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintReservationResponse(response, Full, commandName)
	}
}
