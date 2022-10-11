package convert

import (
	"os"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/commands/convert/reservation"
)

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a resource.",
	Long:  `Convert a resource.`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	ConvertCmd.AddCommand(reservation.ConvertReservationCmd)
}
