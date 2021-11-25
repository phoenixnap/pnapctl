package quotas

import (
	netHttp "net/http"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/printer"
)

const commandName string = "get quotas"

var ID string

var GetQuotasCmd = &cobra.Command{
	Use:          "quota [QUOTA_ID]",
	Short:        "Retrieve one or all quotas for your account.",
	Aliases:      []string{"quotas"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all quotas for your account.

Prints all information about the quotas assigned to your account.
By default, the data is printed in table format.

To print a single quota, a quota ID needs to be passed as an argument.`,
	Example: `
# List all quotas in json format.
pnapctl get quotas [-output <OUTPUT_TYPE>]

# List all details of a desired quota in yaml format.
pnapctl get quota <QUOTA_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			ID = args[0]
			return getQuotas(ID)
		}
		return getQuotas("")
	},
}

func getQuotas(quotaId string) error {
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
			return printer.PrintQuotaListResponse(quotas, commandName)
		} else {
			return printer.PrintQuotaResponse(quota, commandName)
		}
	} else {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	}
}

func init() {
	GetQuotasCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
