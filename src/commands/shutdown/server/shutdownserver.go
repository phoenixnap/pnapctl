package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

const commandName = "shutdown server"

var ShutdownCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Perform a soft shutdown on a specific server.",
	Long:         "Perform a soft shutdown on a specific server.",
	Example:      "pnapctl shutdown server <SERVER_ID>",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return shutdownServer(args[0])
	},
}

func shutdownServer(id string) error {
	result, httpResponse, err := bmcapi.Client.ServerShutdown(id)
	var generatedError = utils.CheckForErrors(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		fmt.Println(result.Result)
		return err
	}
}
