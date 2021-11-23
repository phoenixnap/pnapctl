package sshkey

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

const commandName = "delete ssh-key"

var DeleteSshKeyCmd = &cobra.Command{
	Use:          "ssh-key SSH_KEY_ID",
	Short:        "Deletes a specific SSH Key.",
	Long:         "Deletes a specific SSH Key.",
	Example:      `pnapctl delete ssh-key 5da891e90ab0c59bd28e34ad`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, httpResponse, err := bmcapi.Client.SshKeyDelete(args[0])

		if err != nil {
			return err
		} else if httpResponse.StatusCode != 200 {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}

		fmt.Println(result.Result, result.SshKeyId)
		return nil
	},
}
