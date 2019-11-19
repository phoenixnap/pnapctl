package server

import (
	"bytes"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
	utils "phoenixnap.com/pnap-cli/pnapctl/utility"
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
		var resource = "servers/" + args[0] + "/actions/power-off"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		return utils.HandleClientResponse(response, err, commandName)
	},
}

func init() {
}
