package sshkeys

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName string = "get ssh-keys"

var Full bool

func init() {
	GetSshKeysCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all ssh key details")
	GetSshKeysCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}

var GetSshKeysCmd = &cobra.Command{
	Use:          "ssh-key [SSH_KEY_ID]",
	Short:        "Retrieve one or all ssh-keys for your account.",
	Aliases:      []string{"ssh-keys"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all ssh-keys for your account.

Prints one or all ssh-keys assigned to your account.
By default, the data is printed in table format.

To print a specific ssh-key, an ID linked to the resource needs to be passed as an argument.`,
	Example: `
# List all ssh-keys.
pnapctl get ssh-keys [--full] [--output <OUTPUT_TYPE>]

# List a specific ssh-key.
pnapctl get ssh-key <SSH_KEY_ID> [--full] [--output <OUTPUT_TYPE>]`,
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) >= 1 {
			return getSshKeyById(args[0])
		}
		return getSshKeys()
	},
}

func getSshKeys() error {
	sshKeys, httpResponse, err := bmcapi.Client.SshKeysGet()

	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintSshKeyListResponse(sshKeys, Full, commandName)
	}
}

func getSshKeyById(sshKeyId string) error {
	sshKey, httpResponse, err := bmcapi.Client.SshKeyGetById(sshKeyId)

	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintSshKeyResponse(sshKey, Full, commandName)
	}
}
