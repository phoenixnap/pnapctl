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
	utils.SetupOutputFlag(CreateSshKeyCmd)
	utils.SetupFullFlag(CreateSshKeyCmd, &Full, "ssh key")
	utils.SetupFilenameFlag(CreateSshKeyCmd, &Filename, utils.CREATION)
}

// CreateSshKeyCmd is the command for creating an ssh key.
var CreateSshKeyCmd = &cobra.Command{
	Use:          "ssh-key",
	Short:        "Create a new ssh-key.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a new ssh-key.

Requires a file (yaml or json) containing the information needed to create the ssh-key.`,
	Example: `# Create a new ssh-key as described in sshKeyCreate.yaml
pnapctl create ssh-key --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# sshKeyCreate.yaml
default: true
name: default ssh key
key: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 test2@test`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return createSshKey()
	},
}

func createSshKey() error {
	log.Info().Msg("Creating new Ssh Key...")

	sshKeyCreate, err := models.CreateRequestFromFile[bmcapisdk.SshKeyCreate](Filename)

	if err != nil {
		return err
	}

	// Create the ssh key
	response, err := bmcapi.Client.SshKeyPost(*sshKeyCreate)
	if err != nil {
		return err
	} else {
		return printer.PrintSshKeyResponse(response, Full)
	}
}
