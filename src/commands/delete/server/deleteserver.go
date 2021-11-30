package server

import (
	"fmt"

	"github.com/spf13/cobra"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/ctlerrors"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		result, httpResponse, err := bmcapi.Client.ServerDelete(args[0])

		if httpResponse != nil && !utils.Is2xxSuccessful(httpResponse.StatusCode) {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		} else if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else {
			fmt.Println(result.Result, result.ServerId)
			return nil
		}
	},
}
