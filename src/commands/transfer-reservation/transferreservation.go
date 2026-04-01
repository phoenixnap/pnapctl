package transferreservation

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/transfer-reservation/server"
)

var TransferReservationCmd = &cobra.Command{
	Use:   "transfer-reservation",
	Short: "Transfer a reservation from one point to another(??)",
	Long:  "REPLACE",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	TransferReservationCmd.AddCommand(server.TransferReservationServerCmd)
}
