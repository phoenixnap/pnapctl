package generators

import (
	"time"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/oneof/product"
	"phoenixnap.com/pnapctl/common/models/oneof/ratedusage"
	"phoenixnap.com/pnapctl/common/models/queryparams/audit"
	"phoenixnap.com/pnapctl/common/models/queryparams/billing"
	"phoenixnap.com/pnapctl/common/models/queryparams/network"
)

/*
	Network API
*/

var GeneratePublicNetworksGetQueryParams = Generator(func(item *network.PublicNetworksGetQueryParams) {
	item.Location = &network.AllowedLocations[0]
})

/*
	Audit API
*/

var GenerateQueryParamsSdk = Generator(func(event *audit.EventsGetQueryParams) {
	event.Order = "ASC"
	event.Verb = "POST"
	t, _ := time.Parse(time.RFC3339, event.From.Format(time.RFC3339))
	event.From = &t
	event.To = &t
})

/*
	Billing API
*/

// Rated Usage One Of
func GenerateRatedUsageRecordSdkList() []billingapi.RatedUsageGet200ResponseInner {
	return []billingapi.RatedUsageGet200ResponseInner{
		GenerateBandwidthRecordSdk(),
		GenerateOperatingSystemRecordSdk(),
		GeneratePublicSubnetRecordSdk(),
		GenerateServerRecordSdk(),
		GenerateStorageRecordSdk(),
	}
}

var GenerateBandwidthRecordSdk = OneOfGenerator(ratedusage.BandwidthRecordToInner, UpdateLocation[*billingapi.BandwidthRecord])

var GenerateOperatingSystemRecordSdk = OneOfGenerator(ratedusage.OperatingSystemRecordToInner, UpdateLocation[*billingapi.OperatingSystemRecord])

var GeneratePublicSubnetRecordSdk = OneOfGenerator(ratedusage.PublicSubnetRecordToInner, UpdateLocation[*billingapi.PublicSubnetRecord])

var GenerateServerRecordSdk = OneOfGenerator(ratedusage.ServerRecordToInner, UpdateLocation[*billingapi.ServerRecord])

var GenerateStorageRecordSdk = OneOfGenerator(ratedusage.StorageRecordToInner, UpdateLocation[*billingapi.StorageRecord])

// Product One Of
func GenerateProductSdkList() []billingapi.ProductsGet200ResponseInner {
	return []billingapi.ProductsGet200ResponseInner{
		GenerateBandwidthProduct(),
		GenerateOperatingSystemProduct(),
		GenerateStorageProduct(),
		GenerateServerProduct(),
	}
}

var GenerateBandwidthProduct = OneOfGenerator(product.BandwidthProductToInner, UpdatePricingPlans[*billingapi.Product])

var GenerateOperatingSystemProduct = OneOfGenerator(product.OperatingSystemProductToInner, UpdatePricingPlans[*billingapi.Product])

var GenerateStorageProduct = OneOfGenerator(product.StorageProductToInner, UpdatePricingPlans[*billingapi.Product])

var GenerateServerProduct = OneOfGenerator(product.ServerProductToInner, UpdatePricingPlans[*billingapi.ServerProduct])

// Query Parameters
var GenerateRatedUsageGetQueryParams = Generator(func(params *billing.RatedUsageGetQueryParams) {
	params.FromYearMonth = "2020-10"
	params.ToYearMonth = "2021-10"
	params.ProductCategory = billingapi.BANDWIDTH.Ptr()
})

var GenerateRatedUsageMonthToDateGetQueryParams = Generator(func(params *billing.RatedUsageMonthToDateGetQueryParams) {
	params.ProductCategory = billingapi.BANDWIDTH.Ptr()
})

var GenerateProductAvailabilityGetQueryParams = Generator(func(params *billing.ProductAvailabilityGetQueryParams) {
	params.ProductCategory = []string{"SERVER"}
	params.Location = billingapi.AllowedLocationEnumEnumValues
	params.Solution = []string{"SERVER_RANCHER"}
})

var GenerateReservationGetQueryParams = Generator(func(params *billing.ReservationsGetQueryParams) {
	params.ProductCategory = billingapi.BANDWIDTH.Ptr()
})
