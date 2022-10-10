package generators

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/queryparams/billing"
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
		GenerateBandwidthRecordSdk(),
		GenerateOperatingSystemRecordSdk(),
		GeneratePublicSubnetRecordSdk(),
		GeneratePublicSubnetRecordSdk(),
	}
}

var GenerateBandwidthRecordSdk = OneOfGenerator(func(sdk billingapi.BandwidthRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(RatedUsageBandwidth)
	return billingapi.BandwidthRecordAsRatedUsageGet200ResponseInner(&sdk)
}, UpdateLocation[*billingapi.BandwidthRecord])

var GenerateOperatingSystemRecordSdk = OneOfGenerator(func(sdk billingapi.OperatingSystemRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(RatedUsageOperatingSystem)
	return billingapi.OperatingSystemRecordAsRatedUsageGet200ResponseInner(&sdk)
}, UpdateLocation[*billingapi.OperatingSystemRecord])

var GeneratePublicSubnetRecordSdk = OneOfGenerator(func(sdk billingapi.PublicSubnetRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(RatedUsagePublicSubnet)
	return billingapi.PublicSubnetRecordAsRatedUsageGet200ResponseInner(&sdk)
}, UpdateLocation[*billingapi.PublicSubnetRecord])

var GenerateServerRecordSdk = OneOfGenerator(func(sdk billingapi.ServerRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(RatedUsageServer)
	return billingapi.ServerRecordAsRatedUsageGet200ResponseInner(&sdk)
}, UpdateLocation[*billingapi.ServerRecord])

// Product One Of
func GenerateProductSdkList() []billingapi.ProductsGet200ResponseInner {
	return []billingapi.ProductsGet200ResponseInner{
		GenerateBandwidthProduct(),
		GenerateOperatingSystemProduct(),
		GenerateServerProduct(),
	}
}

var GenerateBandwidthProduct = OneOfGenerator(func(sdk billingapi.Product) billingapi.ProductsGet200ResponseInner {
	sdk.ProductCategory = ProductBandwidth
	return billingapi.ProductAsProductsGet200ResponseInner(&sdk)
}, UpdatePricingPlans[*billingapi.Product])

var GenerateOperatingSystemProduct = OneOfGenerator(func(sdk billingapi.Product) billingapi.ProductsGet200ResponseInner {
	sdk.ProductCategory = ProductOperatingSystem
	return billingapi.ProductAsProductsGet200ResponseInner(&sdk)
}, UpdatePricingPlans[*billingapi.Product])

var GenerateServerProduct = OneOfGenerator(func(sdk billingapi.ServerProduct) billingapi.ProductsGet200ResponseInner {
	sdk.ProductCategory = ProductServer
	return billingapi.ServerProductAsProductsGet200ResponseInner(&sdk)
}, UpdatePricingPlans[*billingapi.ServerProduct])

var GenerateProductAvailabilityGetQueryParams = Generator(func(params *billing.ProductAvailabilityGetQueryParams) {
	params.ProductCategory = []string{"SERVER"}
	params.Location = billingapi.AllowedLocationEnumEnumValues
	params.Solution = []string{"SERVER_RANCHER"}
})

var GenerateReservationGetQueryParams = Generator(func(params *billing.ReservationsGetQueryParams) {
	params.ProductCategory = billingapi.BANDWIDTH.Ptr()
})
