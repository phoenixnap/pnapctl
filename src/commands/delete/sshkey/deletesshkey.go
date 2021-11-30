package sshkey

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName = "delete ssh-key"

var DeleteSshKeyCmd = &cobra.Command{
	Use:          "ssh-key SSH_KEY_ID",
	Short:        "Deletes a specific SSH Key.",
	Long:         "Deletes a specific SSH Key.",
	Example:      `pnapctl delete ssh-key <SSH_KEY_ID>`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, httpResponse, err := bmcapi.Client.SshKeyDelete(args[0])

		if httpResponse != nil && !utils.Is2xxSuccessful(httpResponse.StatusCode) {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		} else if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else {
			fmt.Println(result.Result, result.SshKeyId)
			return nil
		}
	},
}
