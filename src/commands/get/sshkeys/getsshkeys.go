package sshkeys

import (
	netHttp "net/http"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/printer"
	"phoenixnap.com/pnap-cli/common/utils"
)

const commandName string = "get ssh-keys"

var ID string
var Full bool

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
		if len(args) >= 1 {
			ID = args[0]
			return getSshKeys(ID)
		}
		return getSshKeys("")
	},
}

func getSshKeys(sshKeyId string) error {
	var httpResponse *netHttp.Response
	var err error
	var sshKey bmcapisdk.SshKey
	var sshKeys []bmcapisdk.SshKey

	if sshKeyId == "" {
		sshKeys, httpResponse, err = bmcapi.Client.SshKeysGet()
	} else {
		sshKey, httpResponse, err = bmcapi.Client.SshKeyGetById(sshKeyId)
	}

	if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else if utils.Is2xxSuccessful(httpResponse.StatusCode) {
		if sshKeyId == "" {
			return printer.PrintSshKeyListResponse(sshKeys, Full, commandName)
		} else {
			return printer.PrintSshKeyResponse(sshKey, Full, commandName)
		}
	} else {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	}
}

func init() {
	GetSshKeysCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all ssh key details")
	GetSshKeysCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
