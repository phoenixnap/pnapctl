package sshkeys

import (
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var Full bool

func init() {
	utils.SetupOutputFlag(GetSshKeysCmd)
	utils.SetupFullFlag(GetSshKeysCmd, &Full, "ssh key")
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
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		if len(args) >= 1 {
			return getSshKeyById(args[0])
		}
		return getSshKeys()
	},
}

func getSshKeys() error {
	log.Info().Msg("Retrieving list of Ssh Keys...")

	sshKeys, err := bmcapi.Client.SshKeysGet()

	if err != nil {
		return err
	} else {
		return printer.PrintSshKeyListResponse(sshKeys, Full)
	}
}

func getSshKeyById(sshKeyId string) error {
	log.Info().Msgf("Retrieving Ssh Key with ID [%s].", sshKeyId)

	sshKey, err := bmcapi.Client.SshKeyGetById(sshKeyId)

	if err != nil {
		return err
	} else {
		return printer.PrintSshKeyResponse(sshKey, Full)
	}
}
