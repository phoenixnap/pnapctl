package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/billingmodels/productoneof"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

// FromSdk tests
func TestProductActualFromSdk_BandwidthProduct(test_framework *testing.T) {
	bandwidthProduct := GenerateBandwidthProduct()
	ProductsResponse := billingapi.ProductsGet200ResponseInner{
		Product: bandwidthProduct,
	}

	actual := ProductActualFromSdk(ProductsResponse)
	assertEqualAsProduct(test_framework, actual, *bandwidthProduct)
}

func TestProductActualFromSdk_OperatingSystemProduct(test_framework *testing.T) {
	operatingSystemProduct := GenerateOperatingSystemProduct()
	ProductsResponse := billingapi.ProductsGet200ResponseInner{
		Product: operatingSystemProduct,
	}

	actual := ProductActualFromSdk(ProductsResponse)
	assertEqualAsProduct(test_framework, actual, *operatingSystemProduct)
}

func TestProductActualFromSdk_ServerProduct(test_framework *testing.T) {
	serverProduct := GenerateServerProduct()
	ProductsResponse := billingapi.ProductsGet200ResponseInner{
		ServerProduct: serverProduct,
	}

	actual := ProductActualFromSdk(ProductsResponse)
	assertEqualAsServerProduct(test_framework, actual, *serverProduct)
}

// Equality asserts

func assertEqualAsProduct(
	test_framework *testing.T,
	cliOneOf interface{},
	sdkProduct billingapi.Product,
) {
	cliProduct := cliOneOf.(*productoneof.Product)

	assert.Equal(test_framework, cliProduct.ProductCode, sdkProduct.ProductCode)
	assert.Equal(test_framework, string(cliProduct.ProductCategory), sdkProduct.ProductCategory)

	iterutils.AssertOnListElements(
		test_framework,
		cliProduct.Plans,
		sdkProduct.Plans,
		assertEqualPricingPlan,
	)
}

func assertEqualAsServerProduct(
	test_framework *testing.T,
	cliOneOf interface{},
	sdkServerProduct billingapi.ServerProduct,
) {
	cliServerProduct := cliOneOf.(*productoneof.ServerProduct)

	assert.Equal(test_framework, cliServerProduct.ProductCode, sdkServerProduct.ProductCode)
	assert.Equal(test_framework, string(cliServerProduct.ProductCategory), sdkServerProduct.ProductCategory)

	iterutils.AssertOnListElements(
		test_framework,
		cliServerProduct.Plans,
		sdkServerProduct.Plans,
		assertEqualPricingPlan,
	)
}

// Nested asserts
func assertEqualPricingPlan(
	test_framework *testing.T,
	cliPricingPlan productoneof.PricingPlan,
	sdkPricingPlan billingapi.PricingPlan,
) {
	assert.Equal(test_framework, cliPricingPlan.Sku, sdkPricingPlan.Sku)
	assert.Equal(test_framework, cliPricingPlan.SkuDescription, sdkPricingPlan.SkuDescription)
	assert.Equal(test_framework, cliPricingPlan.Location, sdkPricingPlan.Location)
	assert.Equal(test_framework, cliPricingPlan.PricingModel, sdkPricingPlan.PricingModel)
	assert.Equal(test_framework, cliPricingPlan.Price, sdkPricingPlan.Price)
	assert.Equal(test_framework, cliPricingPlan.PriceUnit, sdkPricingPlan.PriceUnit)
	assert.Equal(test_framework, cliPricingPlan.CorrelatedProductCode, sdkPricingPlan.CorrelatedProductCode)
	assert.Equal(test_framework, cliPricingPlan.PackageQuantity, sdkPricingPlan.PackageQuantity)
	assert.Equal(test_framework, cliPricingPlan.PackageUnit, sdkPricingPlan.PackageUnit)
}
