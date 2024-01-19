package enable

import (
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/billing"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	Full bool
)

func init() {
	utils.SetupFullFlag(AutoRenewEnableReservationCmd, &Full, "reservation")
	utils.SetupOutputFlag(AutoRenewEnableReservationCmd)
}

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
		cmdname.SetCommandName(cmd)
		return enableAutoRenewForReservation(args[0])
	},
}

func enableAutoRenewForReservation(id string) error {
	log.Info().Msgf("Enabling auto renewal for Reservation with ID [%s].", id)

	response, err := billing.Client.ReservationEnableAutoRenew(id)

	if err != nil {
		return err
	} else {
		return printer.PrintReservationResponse(response, Full)
	}
}
