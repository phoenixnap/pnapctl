package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName = "power-off server"

var PowerOffServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Perform a hard shutdown on a specific server.",
	Long:         "Perform a hard shutdown on a specific server.",
	Example:      "pnapctl power-off server <SERVER_ID>",
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, httpResponse, err := bmcapi.Client.ServerPowerOff(args[0])
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
