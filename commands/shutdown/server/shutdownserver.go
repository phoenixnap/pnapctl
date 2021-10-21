package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

const commandName = "shutdown server"

var ShutdownCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Perform a soft shutdown on a specific server.",
	Long:         "Perform a soft shutdown on a specific server.",
	Example:      "pnapctl shutdown server 5da891e90ab0c59bd28e34ad",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, response, err := bmcapi.Client.ServerShutdown(args[0])

		if err != nil {
			return err
		} else if response.StatusCode != 200 {
			return ctlerrors.HandleBMCError(response, commandName)
		}

		fmt.Println(result.Result)
		return nil
	},
}
