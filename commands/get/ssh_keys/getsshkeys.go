package ssh_keys

import (
	netHttp "net/http"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/printer"
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

Prints all ssh-keys assigned to your account.
By default, the data is printed in table format.

To print a single ssh-key, an ID linked to the resource needs to be passed as an argument.`,
	Example: `
# List all ssh-keys in json format.
pnapctl get ssh-keys -o json

# List all details of a desired quota in yaml format.
pnapctl get ssh-key 619605811954a568606eb71a -o yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			ID = args[0]
			return getSshKeys(ID)
		}
		return getSshKeys("")
	},
}

func getSshKeys(sshKeyId string) error {
	log.Debug("Getting quotas...")

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
	} else if httpResponse.StatusCode == 200 {
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
	GetSshKeysCmd.PersistentFlags().BoolVar(&Full, "full", false, "Shows all server details")
	GetSshKeysCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
