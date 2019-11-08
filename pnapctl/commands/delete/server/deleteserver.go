package server

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
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
		var resource = "servers/" + args[0]
		var response, err = client.MainClient.PerformDelete(resource)

		return client.HandleClientResponse(response, err, commandName)
	},
}
