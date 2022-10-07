package generators

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/queryparams/billing"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

// one-of-types
var (
	RatedUsageBandwidth       = "bandwidth"
	RatedUsageOperatingSystem = "operating-system"
	RatedUsagePublicSubnet    = "public-ip"
	RatedUsageServer          = "bmc-server"

	ProductBandwidth       = "BANDWIDTH"
	ProductOperatingSystem = "OPERATING_SYSTEM"
	ProductServer          = "SERVER"
)

// Rated Usage
var GenerateRatedUsageGetQueryParams = Generator(func(params *billing.RatedUsageGetQueryParams) {
	params.FromYearMonth = "2020-10"
	params.ToYearMonth = "2021-10"
	params.ProductCategory = billingapi.BANDWIDTH.Ptr()
})

var GenerateRatedUsageMonthToDateGetQueryParams = Generator(func(params *billing.RatedUsageMonthToDateGetQueryParams) {
	params.ProductCategory = billingapi.BANDWIDTH.Ptr()
})

// Rated Usage One Of
func GenerateRatedUsageRecordSdkList() []billingapi.RatedUsageGet200ResponseInner {
	return []billingapi.RatedUsageGet200ResponseInner{
		{
			BandwidthRecord: testutil.AsPointer(GenerateBandwidthRecordSdk()),
		},
		{
			OperatingSystemRecord: testutil.AsPointer(GenerateOperatingSystemRecordSdk()),
		},
		{
			PublicSubnetRecord: testutil.AsPointer(GeneratePublicSubnetRecordSdk()),
		},
		{
			ServerRecord: testutil.AsPointer(GenerateServerRecordSdk()),
		},
	}
}

func updateRatedUsageWithCategory[T interface {
	HasProductCategory
	HasLocation
}](location string) (func(T), func(T)) {
	return updateLocation[T]("PHX"), updateProductCategory[T](location)
}

var GenerateBandwidthRecordSdk = Generator(updateRatedUsageWithCategory[*billingapi.BandwidthRecord](RatedUsageBandwidth))
var GenerateOperatingSystemRecordSdk = Generator(updateRatedUsageWithCategory[*billingapi.OperatingSystemRecord](RatedUsageOperatingSystem))
var GeneratePublicSubnetRecordSdk = Generator(updateRatedUsageWithCategory[*billingapi.PublicSubnetRecord](RatedUsagePublicSubnet))
var GenerateServerRecordSdk = Generator(updateRatedUsageWithCategory[*billingapi.ServerRecord](RatedUsageServer))

var GenerateReservationAutoRenewDisableRequestSdk = Generator[billingapi.ReservationAutoRenewDisableRequest]()
var GenerateReservationRequestSdk = Generator[billingapi.ReservationRequest]()
var GenerateProductsGetQueryParams = Generator[billing.ProductsGetQueryParams]()

// Product One Of
func GenerateProductSdkList() []billingapi.ProductsGet200ResponseInner {
	return []billingapi.ProductsGet200ResponseInner{
		{
			Product: testutil.AsPointer(GenerateBandwidthProduct()),
		},
		{
			Product: testutil.AsPointer(GenerateOperatingSystemProduct()),
		},
		{
			ServerProduct: testutil.AsPointer(GenerateServerProduct()),
		},
	}
}

func updateProductWithCategory[T interface {
	HasProductCategory
	HasPricingPlans
}](category string) (func(T), func(T)) {
	return updatePlans[T]("GB"), updateProductCategory[T](category)
}

var GenerateBandwidthProduct = Generator(updateProductWithCategory[*billingapi.Product](ProductBandwidth))
var GenerateOperatingSystemProduct = Generator(updateProductWithCategory[*billingapi.Product](ProductOperatingSystem))
var GenerateServerProduct = Generator(updateProductWithCategory[*billingapi.ServerProduct](ProductServer))

var GenerateConfigurationDetails = Generator[billingapi.ConfigurationDetails]()
var GenerateProductAvailability = Generator[billingapi.ProductAvailability]()

var GenerateProductAvailabilityGetQueryParams = Generator(func(params *billing.ProductAvailabilityGetQueryParams) {
	params.ProductCategory = []string{"SERVER"}
	params.Location = billingapi.AllowedLocationEnumEnumValues
	params.Solution = []string{"SERVER_RANCHER"}
})

var GenerateReservation = Generator[billingapi.Reservation]()
var GenerateReservationGetQueryParams = Generator(func(params *billing.ReservationsGetQueryParams) {
	params.ProductCategory = billingapi.BANDWIDTH.Ptr()
})
