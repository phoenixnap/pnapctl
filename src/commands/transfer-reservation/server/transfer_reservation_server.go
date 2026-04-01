package server

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var Filename string
var Full bool

func init() {
	utils.SetupOutputFlag(TransferReservationServerCmd)
	utils.SetupFullFlag(TransferReservationServerCmd, &Full, "server")
	utils.SetupFilenameFlag(TransferReservationServerCmd, &Filename, utils.RESERVATION_TRANSFER)
}

var TransferReservationServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Transfer reservation of server elsewhere.",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Long: `Transfer reservation of server elsewhere.
	
Requires a file (yaml or json) containing the information needed to transfer a server's reservation.`,
	Example: `# Transfer a server's reservations using the contents of serverTransferReservation.yaml as request body. 
pnapctl transfer-reservation server <SERVER_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# serverTransferReservation.yaml
targetServerId: "<SERVER ID>"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return transferReservationServer(args[0])
	},
}

func transferReservationServer(id string) error {
	log.Info().Msgf("Transferring reservation of Server with ID [%s].", id)

	transferRequest, err := models.CreateRequestFromFile[bmcapisdk.ReservationTransferDetails](Filename)
	if err != nil {
		return err
	}

	serverResponse, err := bmcapi.Client.ServerTransferReservation(id, *transferRequest)
	if err != nil {
		return err
	} else {
		return printer.PrintServerResponse(serverResponse, Full)
	}
}
