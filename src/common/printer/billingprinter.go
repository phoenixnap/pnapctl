package printer

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/models/tables"
)

func PrintRatedUsageResponse(ratedUsage *billingapi.RatedUsageGet200ResponseInner, full bool, commandName string) error {
	clusterToPrint, err := PrepareRatedUsageForPrinting(*ratedUsage, full, commandName)
	if err != nil {
		return err
	}

	return MainPrinter.PrintOutput(clusterToPrint, commandName)
}

func PrintRatedUsageListResponse(ratedUsages []billingapi.RatedUsageGet200ResponseInner, full bool, commandName string) error {
	clusterListToPrint, err := PrepareRatedUsageListForPrinting(ratedUsages, full, commandName)
	if err != nil {
		return err
	}

	return MainPrinter.PrintOutput(clusterListToPrint, commandName)
}

func PrepareRatedUsageForPrinting(ratedUsage billingapi.RatedUsageGet200ResponseInner, full bool, commandName string) (interface{}, error) {
	table := OutputIsTable()

	switch {
	case table && full:
		record, err := tables.RatedUsageRecordFromSdk(ratedUsage, commandName)
		if err != nil {
			return nil, err
		}
		return record, nil
	case table:
		return tables.ShortRatedUsageRecordFromSdk(ratedUsage), nil
	default:
		return billingmodels.RatedUsageActualFromSdk(ratedUsage), nil
	}
}

func PrepareRatedUsageListForPrinting(ratedUsages []billingapi.RatedUsageGet200ResponseInner, full bool, commandName string) ([]interface{}, error) {
	var clusterList []interface{}

	for _, cluster := range ratedUsages {
		record, err := PrepareRatedUsageForPrinting(cluster, full, commandName)
		if err != nil {
			return nil, err
		}
		clusterList = append(clusterList, record)
	}

	return clusterList, nil
}
