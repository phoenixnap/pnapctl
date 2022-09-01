package printer

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/models/tables"
)

// Rated Usage
func PrintRatedUsageResponse(ratedUsage *billingapi.RatedUsageGet200ResponseInner, full bool, commandName string) error {
	ratedUsageToPrint := PrepareRatedUsageForPrinting(*ratedUsage, full)
	return MainPrinter.PrintOutput(ratedUsageToPrint, commandName)
}

func PrintRatedUsageListResponse(ratedUsages []billingapi.RatedUsageGet200ResponseInner, full bool, commandName string) error {
	ratedUsagesToPrint := prepareOneOfWith(ratedUsages, withFull(full, PrepareRatedUsageForPrinting))
	return MainPrinter.PrintOutput(ratedUsagesToPrint, commandName)
}

func PrepareRatedUsageForPrinting(ratedUsage billingapi.RatedUsageGet200ResponseInner, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case table && full:
		return tables.RatedUsageRecordTableFromSdk(ratedUsage)
	case table:
		return tables.ShortRatedUsageRecordFromSdk(ratedUsage)
	case full:
		return billingmodels.RatedUsageActualFromSdk(ratedUsage)
	default:
		return billingmodels.ShortRatedUsageActualFromSdk(ratedUsage)
	}
}

// Products
func PrintProductResponse(product *billingapi.ProductsGet200ResponseInner, full bool, commandName string) error {
	productToPrint := PrepareProductForPrinting(*product)
	return MainPrinter.PrintOutput(productToPrint, commandName)
}

func PrintProductListResponse(products []billingapi.ProductsGet200ResponseInner, commandName string) error {
	productsToPrint := prepareOneOfWith(products, PrepareProductForPrinting)
	return MainPrinter.PrintOutput(productsToPrint, commandName)
}

func PrepareProductForPrinting(product billingapi.ProductsGet200ResponseInner) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ProductTableFromSdk(product)
	default:
		return billingmodels.ProductActualFromSdk(product)
	}
}
