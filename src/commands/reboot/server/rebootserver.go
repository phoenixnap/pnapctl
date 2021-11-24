package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

const commandName string = "reboot server"

var RebootCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Perform a soft reboot on a specific server.",
	Long:         "Perform a soft reboot on a specific server.",
	Example:      "pnapctl reboot server 5da891e90ab0c59bd28e34ad",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, httpResponse, err := bmcapi.Client.ServerReboot(args[0])

		if err != nil {
			// TODO - Process error from SDK in ctlerrors.
			return err
		} else if httpResponse.StatusCode != 200 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}

		fmt.Println(result.Result)
		return nil
	},
}

func init() {}