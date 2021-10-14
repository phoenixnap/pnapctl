package server

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

const commandName = "power-off server"

var PowerOffServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Perform a hard shutdown on a specific server.",
	Long:         "Perform a hard shutdown on a specific server.",
	Example:      "pnapctl power-off server 5da891e90ab0c59bd28e34ad",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, response, err := client.BmcApiClient.ServersServerIdActionsPowerOffPost(context.Background(), args[0]).Execute()

		if err != nil {
			// TODO - Process error from SDK in ctlerrors.
			return err
		} else if response.StatusCode != 200 {
			return ctlerrors.HandleBMCError(response, commandName)
		}

		fmt.Println(result.Result)
		return nil
	},
}

func init() {
}
