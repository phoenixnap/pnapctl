package printer

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"phoenixnap.com/pnapctl/common/models/tables"
)

func PrintQuotaResponse(quota *bmcapisdk.Quota) error {
	quotaToPrint := PrepareQuotaForPrinting(*quota)
	return MainPrinter.PrintOutput(quotaToPrint)
}

func PrintQuotaListResponse(quotas []bmcapisdk.Quota) error {
	quotaListToPrint := PrepareQuotaListForPrinting(quotas)
	return MainPrinter.PrintOutput(quotaListToPrint)
}

func PrepareQuotaListForPrinting(quotas []bmcapisdk.Quota) []interface{} {
	var quotaList []interface{}

	for _, bmcQuota := range quotas {
		quotaList = append(quotaList, PrepareQuotaForPrinting(bmcQuota))
	}

	return quotaList
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
