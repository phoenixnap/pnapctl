package sshkey

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
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
	RunE: func(_ *cobra.Command, args []string) error {
		return deleteSshKey(args[0])
	},
}

func deleteSshKey(id string) error {
	result, httpResponse, err := bmcapi.Client.SshKeyDelete(id)
	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		fmt.Println(result.Result, result.SshKeyId)
		return nil
	}
}
