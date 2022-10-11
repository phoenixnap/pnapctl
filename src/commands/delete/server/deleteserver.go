package server

import (
	"fmt"

	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName = "delete server"

var DeleteServerCmd = &cobra.Command{
	Use:          "server SERVER_ID",
	Short:        "Deletes a specific server.",
	Long:         "Deletes a specific server.",
	Example:      `pnapctl delete server <SERVER_ID>`,
	Args:         cobra.ExactArgs(1),
	Aliases:      []string{"srv"},
	SilenceUsage: true,
	Deprecated:   "Use the deprovision command instead: pnapctl deprovision server <SERVER_ID> --filename <FILE_PATH>",
	RunE: func(_ *cobra.Command, args []string) error {
		return deleteServer(args[0])
	},
}

func deleteServer(id string) error {
	result, httpResponse, err := bmcapi.Client.ServerDelete(id)
	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		fmt.Println(result.Result, result.ServerId)
		return nil
	}
}
