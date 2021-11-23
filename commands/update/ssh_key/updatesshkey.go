package ssh_key

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/bmcapimodels"
	"phoenixnap.com/pnap-cli/common/printer"
)

// Filename is the filename from which to retrieve a complex object
var Filename string

var commandName = "update ssh-key"

var Full bool

// UpdateSshKeyCmd is the command for creating a server.
var UpdateSshKeyCmd = &cobra.Command{
	Use:          "ssh-key SSH_KEY_ID",
	Short:        "Update an ssh-key.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	Long: `Update an ssh-key.

Requires a file (yaml or json) containing the information needed to modify the ssh-key.`,
	Example: `# update an ssh-key as described in sshKeyEdit.yaml
pnapctl update ssh-key 5da891e90ab0c59bd28e34ad --filename ~/sshKeyUpdate.yaml

# sshKeyUpdate.yaml
default: true
name: default ssh key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sshKeyUpdate, err := bmcapimodels.CreateSshKeyUpdateRequestFromFile(Filename, commandName)

		if err != nil {
			return err
		}

		// Create the server
		response, httpResponse, err := bmcapi.Client.SshKeyPut(args[0], *sshKeyUpdate)

		if err != nil {
			return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
		} else if httpResponse.StatusCode == 200 {
			return printer.PrintSshKeyResponse(response, Full, commandName)
		} else {
			return ctlerrors.HandleBMCError(httpResponse, commandName)
		}
	},
}

func init() {
	UpdateSshKeyCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	UpdateSshKeyCmd.Flags().StringVarP(&Filename, "filename", "f", "", "File containing required information for creation")
	UpdateSshKeyCmd.MarkFlagRequired("filename")
}
