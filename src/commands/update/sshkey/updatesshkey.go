package sshkey

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var Full bool

func init() {
	utils.SetupOutputFlag(UpdateSshKeyCmd)
	utils.SetupFullFlag(UpdateSshKeyCmd, &Full, "ssh key")
	utils.SetupFilenameFlag(UpdateSshKeyCmd, &Filename, utils.UPDATING)
}

// UpdateSshKeyCmd is the command for creating a server.
var UpdateSshKeyCmd = &cobra.Command{
	Use:          "ssh-key SSH_KEY_ID",
	Short:        "Update an ssh-key.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Update an ssh-key.

Requires a file (yaml or json) containing the information needed to modify the ssh-key.`,
	Example: `# Update an ssh-key as per sshKeyUpdate.yaml
pnapctl update ssh-key <SSH_KEY_ID> --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# sshKeyUpdate.yaml
default: true
name: default ssh key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		return updateSshKey(args[0])
	},
}

func updateSshKey(id string) error {
	log.Info().Msgf("Updating Ssh Key with ID [%s].", id)

	sshKeyUpdate, err := models.CreateRequestFromFile[bmcapisdk.SshKeyUpdate](Filename)

	if err != nil {
		return err
	}

	// update the ssh key
	response, err := bmcapi.Client.SshKeyPut(id, *sshKeyUpdate)
	if err != nil {
		return err
	} else {
		return printer.PrintSshKeyResponse(response, Full)
	}
}
