package quotas

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/bmcapi"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

func init() {
	utils.SetupOutputFlag(GetQuotasCmd)
}

var GetQuotasCmd = &cobra.Command{
	Use:          "quota [QUOTA_ID]",
	Short:        "Retrieve one or all quotas for your account.",
	Aliases:      []string{"quotas"},
	SilenceUsage: true,
	Args:         cobra.MaximumNArgs(1),
	Long: `Retrieve one or all quotas for your account.

Prints all information about the quotas assigned to your account.
By default, the data is printed in table format.

To print a specific quota, a quota ID needs to be passed as an argument.`,
	Example: `
# List all quotas in.
pnapctl get quotas [--output <OUTPUT_TYPE>]

# List a specific quota.
pnapctl get quota <QUOTA_ID> [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmdname.SetCommandName(cmd)
		if len(args) >= 1 {
			return getQuotaById(args[0])
		}
		return getQuotas()
	},
}

func getQuotas() error {
	quotas, err := bmcapi.Client.QuotasGet()

	if err != nil {
		return err
	} else {
		return printer.PrintQuotaListResponse(quotas)
	}
}

func getQuotaById(quotaId string) error {
	quota, err := bmcapi.Client.QuotaGetById(quotaId)

	if err != nil {
		return err
	} else {
		return printer.PrintQuotaResponse(quota)
	}
}
