package generators

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/oneof/product"
	"phoenixnap.com/pnapctl/common/models/oneof/ratedusage"
)

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
