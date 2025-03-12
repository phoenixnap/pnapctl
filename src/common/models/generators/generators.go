package generators

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
	"phoenixnap.com/pnapctl/common/models/oneof/product"
	"phoenixnap.com/pnapctl/common/models/oneof/ratedusage"
)

/*
	GENERATORS
	Some types - like ones using AllOf and/or OneOf, are more complicated and
		cannot be easily generated directly.
	These methods assist with crafting these types -- usually they are wrappers
		around generation utilities, using UPDATERS to ensure that the data is valid.
*/

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

var GenerateBandwidthRecordSdk = OneOfGenerator(ratedusage.BandwidthRecordToInner,
	UpdateRatedUsageRecord[*billingapi.BandwidthRecord])

var GenerateOperatingSystemRecordSdk = OneOfGenerator(ratedusage.OperatingSystemRecordToInner,
	UpdateRatedUsageRecord[*billingapi.OperatingSystemRecord])

var GeneratePublicSubnetRecordSdk = OneOfGenerator(ratedusage.PublicSubnetRecordToInner,
	UpdateRatedUsageRecord[*billingapi.PublicSubnetRecord])

var GenerateServerRecordSdk = OneOfGenerator(ratedusage.ServerRecordToInner,
	UpdateRatedUsageRecord[*billingapi.ServerRecord])

var GenerateStorageRecordSdk = OneOfGenerator(ratedusage.StorageRecordToInner,
	UpdateRatedUsageRecord[*billingapi.StorageRecord])

// Product One Of
func GenerateProductSdkList() []billingapi.ProductsGet200ResponseInner {
	return []billingapi.ProductsGet200ResponseInner{
		GenerateBandwidthProduct(),
		GenerateOperatingSystemProduct(),
		GenerateStorageProduct(),
		GenerateServerProduct(),
	}
}

var GenerateBandwidthProduct = OneOfGenerator(product.BandwidthProductToInner,
	UpdatePricingPlans[*billingapi.Product])

var GenerateOperatingSystemProduct = OneOfGenerator(product.OperatingSystemProductToInner,
	UpdatePricingPlans[*billingapi.Product])

var GenerateStorageProduct = OneOfGenerator(product.StorageProductToInner,
	UpdatePricingPlans[*billingapi.Product])

var GenerateServerProduct = OneOfGenerator(product.ServerProductToInner,
	UpdatePricingPlans[*billingapi.ServerProduct])
