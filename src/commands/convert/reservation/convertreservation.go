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
	utils.SetupFullFlag(ConvertReservationCmd, &Full, "reservation")
	utils.SetupOutputFlag(ConvertReservationCmd)
	utils.SetupFilenameFlag(ConvertReservationCmd, &Filename, utils.CONVERSION)
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
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return convertReservation(args[0])
	},
}

func convertReservation(id string) error {
	request, err := models.CreateRequestFromFile[billingapi.ReservationRequest](Filename)
	if err != nil {
		return err
	}

	response, httpResponse, err := billing.Client.ReservationConvert(id, *request)
	if err := utils.CheckErrs(httpResponse, err); err != nil {
		return err
	} else {
		return printer.PrintReservationResponse(response, Full)
	}
}
