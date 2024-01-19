package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var RebootCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Perform a soft reboot on a specific server.",
	Long:         "Perform a soft reboot on a specific server.",
	Example:      "pnapctl reboot server <SERVER_ID>",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return rebootServer(args[0])
	},
}

func rebootServer(id string) error {
	log.Info().Msgf("Rebooting Server with ID [%s].", id)

	result, err := bmcapi.Client.ServerReboot(id)
	if err != nil {
		return err
	} else {
		fmt.Println(result.Result)
		return nil
	}
}
