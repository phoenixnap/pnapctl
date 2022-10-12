package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var PowerOnServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Powers on a specific server.",
	Long:         "Powers on a specific server.",
	Example:      `pnapctl power-on server <SERVER_ID>`,
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return powerOnServer(args[0])
	},
}

func powerOnServer(id string) error {
	result, httpResponse, err := bmcapi.Client.ServerPowerOn(id)
	var generatedError = utils.CheckErrs(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		fmt.Println(result.Result)
		return nil
	}
}
