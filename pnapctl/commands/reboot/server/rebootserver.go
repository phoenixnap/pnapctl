package server

import (
	"bytes"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
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
		resource := "servers/" + args[0] + "/actions/reboot"
		response, err := client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		return client.HandleClientResponse(response, err, commandName)
	},
}

func init() {}
