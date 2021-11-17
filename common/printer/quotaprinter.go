package printer

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	quotaModel "phoenixnap.com/pnap-cli/common/models/bmcapimodels"
	"phoenixnap.com/pnap-cli/common/models/tables"
)

func PrintQuotaResponse(quota bmcapisdk.Quota, full bool, commandName string) error {
	quotaToPrint := PrepareQuotaForPrinting(quota, full)
	return MainPrinter.PrintOutput(quotaToPrint, commandName)
}

func PrintQuotaListResponse(quotas []bmcapisdk.Quota, full bool, commandName string) error {
	quotaListToPrint := PrepareQuotaListForPrinting(quotas, full)
	return MainPrinter.PrintOutput(quotaListToPrint, commandName)
}

func PrepareQuotaListForPrinting(quotas []bmcapisdk.Quota, full bool) []interface{} {
	var quotaList []interface{}

	for _, bmcQuota := range quotas {
		quotaList = append(quotaList, PrepareQuotaForPrinting(bmcQuota, full))
	}

	return quotaList
}

func PrepareQuotaForPrinting(quota bmcapisdk.Quota, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ToQuotaTable(quota)
	default:
		return quotaModel.QuotaSdkToDto(quota)
	}
}
