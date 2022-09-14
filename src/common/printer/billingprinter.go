package printer

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
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

// Reservations
func PrintReservationResponse(reservation *billingapi.Reservation, full bool, commandName string) error {
	reservationToPrint := PrepareReservationForPrinting(*reservation, full)
	return MainPrinter.PrintOutput(reservationToPrint, commandName)
}

func PrintReservationListResponse(reservations []billingapi.Reservation, full bool, commandName string) error {
	reservationsToPrint := iterutils.Map(reservations, withFull(full, PrepareReservationForPrinting))
	return MainPrinter.PrintOutput(reservationsToPrint, commandName)
}

func PrepareReservationForPrinting(reservation billingapi.Reservation, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case table && full:
		return tables.ReservationTableFromSdk(reservation)
	case table:
		return tables.ShortReservationTableFromSdk(reservation)
	case full:
		return billingmodels.ReservationFromSdk(reservation)
	default:
		return billingmodels.ShortReservationFromSdk(reservation)
	}
}

// Configuration Details
func PrintConfigurationDetailsResponse(configurationDetails *billingapi.ConfigurationDetails, commandName string) error {
	configurationDetailsToPrint := PrepareConfigurationDetailsForPrinting(*configurationDetails)
	return MainPrinter.PrintOutput(configurationDetailsToPrint, commandName)
}

func PrintConfigurationDetailsListResponse(configurationDetails []billingapi.ConfigurationDetails, commandName string) error {
	configurationDetailsToPrint := iterutils.Map(configurationDetails, PrepareConfigurationDetailsForPrinting)
	return MainPrinter.PrintOutput(configurationDetailsToPrint, commandName)
}

func PrepareConfigurationDetailsForPrinting(configurationDetails billingapi.ConfigurationDetails) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ConfigurationDetailsTableFromSdk(configurationDetails)
	default:
		return billingmodels.ConfigurationDetailsFromSdk(configurationDetails)
	}
}

// Product Availability
func PrintProductAvailabilityResponse(productAvailability *billingapi.ProductAvailability, commandName string) error {
	productAvailabilityToPrint := PrepareProductAvailabilityForPrinting(*productAvailability)
	return MainPrinter.PrintOutput(productAvailabilityToPrint, commandName)
}

func PrintProductAvailabilityListResponse(productAvailability []billingapi.ProductAvailability, commandName string) error {
	productAvailabilityToPrint := iterutils.Map(productAvailability, PrepareProductAvailabilityForPrinting)
	return MainPrinter.PrintOutput(productAvailabilityToPrint, commandName)
}

func PrepareProductAvailabilityForPrinting(productAvailability billingapi.ProductAvailability) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ProductAvailabilityTableFromSdk(productAvailability)
	default:
		return billingmodels.ProductAvailabilityFromSdk(productAvailability)
	}
}
