package server

import (
	"bytes"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client"
	utils "phoenixnap.com/pnap-cli/helpers/utility"
)

const commandName string = "power-on server"

var PowerOnServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Powers on a specific server.",
	Long:         "Powers on a specific server.",
	Example:      `pnapctl power-on server 5da891e90ab0c59bd28e34ad`,
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var resource = "servers/" + args[0] + "/actions/power-on"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		return utils.HandleClientResponse(response, err, commandName)
	},
}

func init() {
}
