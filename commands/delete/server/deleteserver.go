package server

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"phoenixnap.com/pnap-cli/common/client"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

const commandName = "delete server"

var DeleteServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Deletes a specific server.",
	Long:         "Deletes a specific server.",
	Example:      `pnapctl delete server 5da891e90ab0c59bd28e34ad`,
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// var response, err = client.MainClient.PerformDelete(resource)
		result, response, err := client.BmcApiClient.ServersServerIdDelete(context.Background(), args[0]).Execute()

		if err != nil {
			return err
		} else if response.StatusCode != 200 {
			return ctlerrors.HandleBMCError(response, commandName)
		}

		fmt.Println(result.Result, result.ServerId)
		return nil
	},
}
