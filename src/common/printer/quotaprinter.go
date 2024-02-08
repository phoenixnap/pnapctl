package printer

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func PrintQuotaResponse(quota *bmcapisdk.Quota) error {
	quotaToPrint := PrepareQuotaForPrinting(*quota)
	return MainPrinter.PrintOutput(quotaToPrint)
}

func PrintQuotaListResponse(quotas []bmcapisdk.Quota) error {
	quotaListToPrint := iterutils.Map(quotas, PrepareQuotaForPrinting)
	return MainPrinter.PrintOutput(quotaListToPrint)
}

func PrepareQuotaForPrinting(quota bmcapisdk.Quota) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ToQuotaTable(quota)
	default:
		return quota
	}
}
