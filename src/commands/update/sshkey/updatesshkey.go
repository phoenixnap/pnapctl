package sshkey

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

// Filename is the filename from which to retrieve the request body
var Filename string

var commandName = "update ssh-key"

var Full bool

func init() {
	UpdateSshKeyCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all ssh key details")
	UpdateSshKeyCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	UpdateSshKeyCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	UpdateSshKeyCmd.MarkFlagRequired("filename")
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
	RunE: func(_ *cobra.Command, args []string) error {
		return updateSshKey(args[0])
	},
}

func updateSshKey(id string) error {
	sshKeyUpdate, err := models.CreateRequestFromFile[bmcapisdk.SshKeyUpdate](Filename, commandName)

	if err != nil {
		return err
	}

	// update the ssh key
	response, httpResponse, err := bmcapi.Client.SshKeyPut(id, *sshKeyUpdate)
	var generatedError = utils.CheckForErrors(httpResponse, err, commandName)

	if *generatedError != nil {
		return *generatedError
	} else {
		return printer.PrintSshKeyResponse(response, Full, commandName)
	}
}
