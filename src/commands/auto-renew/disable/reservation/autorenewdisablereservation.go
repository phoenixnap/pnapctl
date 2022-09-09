package reservation

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "auto-renew disable reservation"

var AutoRenewDisableReservationCmd = &cobra.Command{
	Use:          "reservation [RESERVATION_ID]",
	Short:        "Disable auto-renew for a reservation",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	Long: `Disable auto-renew for a reservation.
	
Requires a file (yaml or json) containing the information needed to disable auto-renew.`,
	Example: `
# Disable auto-renew for a specific reservation
pnapctl auto-renew disable reservation <RESERVATION_ID> --filename=<FILENAME>

# reservationAutoRenewDisable.yaml
autoRenewDisableReasons: "disable reason"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		request, err := billingmodels.CreateReservationAutoRenewDisableRequestFromFile(Filename, commandName)
		if err != nil {
			return err
		}

		response, httpResponse, err := billing.Client.ReservationDisableAutoRenew(args[0], *request)
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
	utils.SetupFullFlag(AutoRenewDisableReservationCmd, &Full, "reservation")
	utils.SetupOutputFlag(AutoRenewDisableReservationCmd)

	AutoRenewDisableReservationCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	AutoRenewDisableReservationCmd.MarkFlagRequired("filename")
}