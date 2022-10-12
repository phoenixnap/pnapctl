package printer

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

// Rated Usage
func PrintRatedUsageResponse(ratedUsage *billingapi.RatedUsageGet200ResponseInner, full bool) error {
	ratedUsageToPrint := PrepareRatedUsageForPrinting(*ratedUsage, full)
	return MainPrinter.PrintOutput(ratedUsageToPrint)
}

func PrintRatedUsageListResponse(ratedUsages []billingapi.RatedUsageGet200ResponseInner, full bool) error {
	ratedUsagesToPrint := prepareOneOfWith(ratedUsages, withFull(full, PrepareRatedUsageForPrinting))
	return MainPrinter.PrintOutput(ratedUsagesToPrint)
}

func PrepareRatedUsageForPrinting(ratedUsage billingapi.RatedUsageGet200ResponseInner, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case table && full:
		return tables.RatedUsageRecordTableFromSdk(ratedUsage)
	case table:
		return tables.ShortRatedUsageRecordFromSdk(ratedUsage)
	default:
		return ratedUsage
	}
}

// Products
func PrintProductResponse(product *billingapi.ProductsGet200ResponseInner) error {
	productToPrint := PrepareProductForPrinting(*product)
	return MainPrinter.PrintOutput(productToPrint)
}

func PrintProductListResponse(products []billingapi.ProductsGet200ResponseInner) error {
	productsToPrint := prepareOneOfWith(products, PrepareProductForPrinting)
	return MainPrinter.PrintOutput(productsToPrint)
}

func PrepareProductForPrinting(product billingapi.ProductsGet200ResponseInner) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ProductTableFromSdk(product)
	default:
		return product
	}
}

// Reservations
func PrintReservationResponse(reservation *billingapi.Reservation, full bool) error {
	reservationToPrint := PrepareReservationForPrinting(*reservation, full)
	return MainPrinter.PrintOutput(reservationToPrint)
}

func PrintReservationListResponse(reservations []billingapi.Reservation, full bool) error {
	reservationsToPrint := iterutils.Map(reservations, withFull(full, PrepareReservationForPrinting))
	return MainPrinter.PrintOutput(reservationsToPrint)
}

func PrepareReservationForPrinting(reservation billingapi.Reservation, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case table && full:
		return tables.ReservationTableFromSdk(reservation)
	case table:
		return tables.ShortReservationTableFromSdk(reservation)
	default:
		return reservation
	}
}

// Configuration Details
func PrintConfigurationDetailsResponse(configurationDetails *billingapi.ConfigurationDetails) error {
	configurationDetailsToPrint := PrepareConfigurationDetailsForPrinting(*configurationDetails)
	return MainPrinter.PrintOutput(configurationDetailsToPrint)
}

func PrintConfigurationDetailsListResponse(configurationDetails []billingapi.ConfigurationDetails) error {
	configurationDetailsToPrint := iterutils.Map(configurationDetails, PrepareConfigurationDetailsForPrinting)
	return MainPrinter.PrintOutput(configurationDetailsToPrint)
}

func PrepareConfigurationDetailsForPrinting(configurationDetails billingapi.ConfigurationDetails) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ConfigurationDetailsTableFromSdk(configurationDetails)
	default:
		return configurationDetails
	}
}

// Product Availability
func PrintProductAvailabilityResponse(productAvailability *billingapi.ProductAvailability) error {
	productAvailabilityToPrint := PrepareProductAvailabilityForPrinting(*productAvailability)
	return MainPrinter.PrintOutput(productAvailabilityToPrint)
}

func PrintProductAvailabilityListResponse(productAvailability []billingapi.ProductAvailability) error {
	productAvailabilityToPrint := iterutils.Map(productAvailability, PrepareProductAvailabilityForPrinting)
	return MainPrinter.PrintOutput(productAvailabilityToPrint)
}

func PrepareProductAvailabilityForPrinting(productAvailability billingapi.ProductAvailability) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ProductAvailabilityTableFromSdk(productAvailability)
	default:
		return productAvailability
	}
}
