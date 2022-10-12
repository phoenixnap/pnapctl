package sshkey

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var DeleteSshKeyCmd = &cobra.Command{
	Use:          "ssh-key SSH_KEY_ID",
	Short:        "Deletes a specific SSH Key.",
	Long:         "Deletes a specific SSH Key.",
	Example:      `pnapctl delete ssh-key <SSH_KEY_ID>`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return deleteSshKey(args[0])
	},
}

func deleteSshKey(id string) error {
	result, httpResponse, err := bmcapi.Client.SshKeyDelete(id)
	var generatedError = utils.CheckErrs(httpResponse, err)

	if generatedError != nil {
		return generatedError
	} else {
		fmt.Println(result.Result, result.SshKeyId)
		return nil
	}
}
