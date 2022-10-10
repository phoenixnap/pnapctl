package generators

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/queryparams/billing"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
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

var GenerateBandwidthRecordSdk = Generator(func(sdk *billingapi.BandwidthRecord) {
	sdk.Location = "PHX"
	sdk.ProductCategory = RatedUsageBandwidth
})
var GenerateOperatingSystemRecordSdk = Generator(func(sdk *billingapi.OperatingSystemRecord) {
	sdk.Location = "PHX"
	sdk.ProductCategory = RatedUsageOperatingSystem
})
var GeneratePublicSubnetRecordSdk = Generator(func(sdk *billingapi.PublicSubnetRecord) {
	sdk.Location = "PHX"
	sdk.ProductCategory = RatedUsagePublicSubnet
})
var GenerateServerRecordSdk = Generator(func(sdk *billingapi.ServerRecord) {
	sdk.Location = "PHX"
	sdk.ProductCategory = RatedUsageServer
})

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

func updatePricingPlan(sdk billingapi.PricingPlan) billingapi.PricingPlan {
	sdk.PriceUnit = billingapi.GB
	return sdk
}

var GenerateBandwidthProduct = Generator(func(sdk *billingapi.Product) {
	sdk.Plans = iterutils.Map(sdk.Plans, updatePricingPlan)
	sdk.ProductCategory = ProductBandwidth
})
var GenerateOperatingSystemProduct = Generator(func(sdk *billingapi.Product) {
	sdk.Plans = iterutils.Map(sdk.Plans, updatePricingPlan)
	sdk.ProductCategory = ProductOperatingSystem
})
var GenerateServerProduct = Generator(func(sdk *billingapi.ServerProduct) {
	sdk.Plans = iterutils.Map(sdk.Plans, updatePricingPlan)
	sdk.ProductCategory = ProductServer
})

var GenerateProductAvailabilityGetQueryParams = Generator(func(params *billing.ProductAvailabilityGetQueryParams) {
	params.ProductCategory = []string{"SERVER"}
	params.Location = billingapi.AllowedLocationEnumEnumValues
	params.Solution = []string{"SERVER_RANCHER"}
})

var GenerateReservationGetQueryParams = Generator(func(params *billing.ReservationsGetQueryParams) {
	params.ProductCategory = billingapi.BANDWIDTH.Ptr()
})
