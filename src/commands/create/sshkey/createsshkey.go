package sshkey

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/bmcapimodels"
	"phoenixnap.com/pnap-cli/common/printer"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "create ssh-key"

var Full bool

// CreateSshKeyCmd is the command for creating an ssh key.
var CreateSshKeyCmd = &cobra.Command{
	Use:          "ssh-key",
	Short:        "Create a new ssh-key.",
	Args:         cobra.ExactArgs(0),
	SilenceUsage: true,
	Long: `Create a new ssh-key.

Requires a file (yaml or json) containing the information needed to create the ssh-key.`,
	Example: `# create a new ssh-key as described in sshKeyCreate.yaml
pnapctl create ssh-key --filename <FILE_PATH> [--full] [--output <OUTPUT_TYPE>]

# sshKeyCreate.yaml
default: true
name: default ssh key
key: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCyVGaw1PuEl98f4/7Kq3O9ZIvDw2OFOSXAFVqilSFNkHlefm1iMtPeqsIBp2t9cbGUf55xNDULz/bD/4BCV43yZ5lh0cUYuXALg9NI29ui7PEGReXjSpNwUD6ceN/78YOK41KAcecq+SS0bJ4b4amKZIJG3JWmDKljtv1dmSBCrTmEAQaOorxqGGBYmZS7NQumRe4lav5r6wOs8OACMANE1ejkeZsGFzJFNqvr5DuHdDL5FAudW23me3BDmrM9ifUzzjl1Jwku3bnRaCcjaxH8oTumt1a00mWci/1qUlaVFft085yvVq7KZbF2OPPbl+erDW91+EZ2FgEi+v1/CSJ5 test2@test`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sshKeyCreate, err := bmcapimodels.CreateSshKeyCreateRequestFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		// Create the ssh key
		response, httpResponse, err := bmcapi.Client.SshKeyPost(*sshKeyCreate)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else if httpResponse.StatusCode == 201 {
			return printer.PrintSshKeyResponse(response, Full, commandName)
		} else {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}
	},
}

func init() {
	CreateSshKeyCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all ssh key details")
	CreateSshKeyCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	CreateSshKeyCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	CreateSshKeyCmd.MarkFlagRequired("filename")
}
