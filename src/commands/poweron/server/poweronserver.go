package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName string = "power-on server"

var PowerOnServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Powers on a specific server.",
	Long:         "Powers on a specific server.",
	Example:      `pnapctl power-on server <SERVER_ID>`,
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, httpResponse, err := bmcapi.Client.ServerPowerOn(args[0])
		var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

		if *generatedError != nil {
			return *generatedError
		} else {
			fmt.Println(result.Result)
			return nil
		}
	},
}

func init() {
}
