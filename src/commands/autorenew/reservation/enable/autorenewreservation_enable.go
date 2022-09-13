package enable

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

var commandName = "auto-renew reservation enable"

var AutoRenewEnableReservationCmd = &cobra.Command{
	Use:          "enable [RESERVATION_ID]",
	Short:        "Enable auto-renew for a reservation",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	Long:         `Enable auto-renew for a reservation.`,
	Example: `
# Enable auto-renew for a specific reservation
pnapctl auto-renew reservation enable <RESERVATION_ID>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		response, httpResponse, err := billing.Client.ReservationEnableAutoRenew(args[0])
		generatedError := utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			return printer.PrintReservationResponse(response, Full, commandName)
		}
	},
}

var (
	Full bool
)

func init() {
	utils.SetupFullFlag(AutoRenewEnableReservationCmd, &Full, "reservation")
	utils.SetupOutputFlag(AutoRenewEnableReservationCmd)
}
