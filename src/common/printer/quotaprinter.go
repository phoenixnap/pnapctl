package printer

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/quotamodels"
	"phoenixnap.com/pnapctl/common/models/tables"
)

func PrintQuotaResponse(quota bmcapisdk.Quota, commandName string) error {
	quotaToPrint := PrepareQuotaForPrinting(quota)
	return MainPrinter.PrintOutput(quotaToPrint, commandName)
}

func PrintQuotaListResponse(quotas []bmcapisdk.Quota, commandName string) error {
	quotaListToPrint := PrepareQuotaListForPrinting(quotas)
	return MainPrinter.PrintOutput(quotaListToPrint, commandName)
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
		return quotamodels.QuotaSdkToDto(quota)
	}
}
