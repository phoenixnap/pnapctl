package quotas

import (
	netHttp "net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/printer"
)

const commandName string = "get quotas"

var Full bool
var ID string

var GetQuotasCmd = &cobra.Command{
	Use:          "quota [QUOTA_ID]",
	Short:        "Retrieve one or all quotas for your account.",
	Aliases:      []string{"quotas"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all quotas for your account.

Prints brief or detailed information about the quotas assigned to your account.
By default, the data is printed in table format.

To print a single quota, a quota ID needs to be passed as an argument.`,
	Example: `
# List all quotas in json format.
pnapctl get quotas -o json

# List all details of a desired quota in yaml format.
pnapctl get quota bmc.servers.max_count -o yaml --full`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			ID = args[0]
			return getQuotas(ID)
		}
		return getQuotas("")
	},
}

func getQuotas(quotaId string) error {
	log.Debug("Getting servers...")

	var httpResponse *netHttp.Response
	var err error
	var quota bmcapisdk.Quota
	var quotas []bmcapisdk.Quota

	if quotaId == "" {
		quotas, httpResponse, err = bmcapi.Client.QuotasGet()
	} else {
		quota, httpResponse, err = bmcapi.Client.QuotaGetById(quotaId)
	}

	if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else if httpResponse.StatusCode == 200 {
		if quotaId == "" {
			return printer.PrintQuotaListResponse(quotas, Full, commandName)
		} else {
			return printer.PrintQuotaResponse(quota, Full, commandName)
		}
	} else {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	}
}

func init() {
	GetQuotasCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
