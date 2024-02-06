package disable

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

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
	utils.SetupFullFlag(AutoRenewDisableReservationCmd, &Full, "reservation")
	utils.SetupOutputFlag(AutoRenewDisableReservationCmd)
	utils.SetupFilenameFlag(AutoRenewDisableReservationCmd, &Filename, utils.CREATION)
}

var AutoRenewDisableReservationCmd = &cobra.Command{
	Use:          "disable [RESERVATION_ID]",
	Short:        "Disable auto-renew for a reservation",
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	Long: `Disable auto-renew for a reservation.
	
Requires a file (yaml or json) containing the information needed to disable auto-renew.`,
	Example: `
# Disable auto-renew for a specific reservation
pnapctl auto-renew reservation disable <RESERVATION_ID> --filename=<FILENAME>

# reservationAutoRenewDisable.yaml
autoRenewDisableReasons: "disable reason"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return disableAutoRenewForReservation(args[0])
	},
}

func disableAutoRenewForReservation(id string) error {
	log.Info().Msgf("Disabling auto renewal for Reservation with ID [%s].", id)

	request, err := models.CreateRequestFromFile[billingapi.ReservationAutoRenewDisableRequest](Filename)
	if err != nil {
		return err
	}

	response, err := billing.Client.ReservationDisableAutoRenew(id, *request)
	if err != nil {
		return err
	} else {
		return printer.PrintReservationResponse(response, Full)
	}
}
