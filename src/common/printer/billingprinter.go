package printer

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/models/tables"
)

func PrintRatedUsageResponse(ratedUsage *billingapi.RatedUsageGet200ResponseInner, full bool, commandName string) error {
	clusterToPrint := PrepareRatedUsageForPrinting(*ratedUsage, full)
	return MainPrinter.PrintOutput(clusterToPrint, commandName)
}

func PrintRatedUsageListResponse(ratedUsages []billingapi.RatedUsageGet200ResponseInner, full bool, commandName string) error {
	clusterListToPrint := PrepareRatedUsageListForPrinting(ratedUsages, full)
	return MainPrinter.PrintOutput(clusterListToPrint, commandName)
}

func PrepareRatedUsageForPrinting(ratedUsage billingapi.RatedUsageGet200ResponseInner, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case table && full:
		return tables.RatedUsageRecordFromSdk(ratedUsage)
	case table:
		return tables.ShortRatedUsageRecordFromSdk(ratedUsage)
	default:
		return billingmodels.RatedUsageActualFromSdk(ratedUsage)
	}
}

func PrepareRatedUsageListForPrinting(ratedUsages []billingapi.RatedUsageGet200ResponseInner, full bool) []interface{} {
	var clusterList []interface{}

	for _, cluster := range ratedUsages {
		clusterList = append(clusterList, PrepareRatedUsageForPrinting(cluster, full))
	}

	return clusterList
}
