package tables

import (
	"fmt"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

// Full version
func TestProductActualFromSdk_BandwidthProduct(test_framework *testing.T) {
	bandwidthProduct := billingmodels.GenerateBandwidthProduct()
	ProductsResponse := billingapi.ProductsGet200ResponseInner{
		Product: bandwidthProduct,
	}

	actual := ProductTableFromSdk(ProductsResponse)
	assertEqualAsProduct(test_framework, *bandwidthProduct, *actual)
}

func TestProductActualFromSdk_OperatingSystemProduct(test_framework *testing.T) {
	operatingSystemProduct := billingmodels.GenerateOperatingSystemProduct()
	ProductsResponse := billingapi.ProductsGet200ResponseInner{
		Product: operatingSystemProduct,
	}

	actual := ProductTableFromSdk(ProductsResponse)
	assertEqualAsProduct(test_framework, *operatingSystemProduct, *actual)
}

func TestProductActualFromSdk_StorageProduct(test_framework *testing.T) {
	storageProduct := billingmodels.GenerateStorageProduct()
	ProductsResponse := billingapi.ProductsGet200ResponseInner{
		Product: storageProduct,
	}

	actual := ProductTableFromSdk(ProductsResponse)
	assertEqualAsProduct(test_framework, *storageProduct, *actual)
}

func TestProductActualFromSdk_ServerProduct(test_framework *testing.T) {
	serverProduct := billingmodels.GenerateServerProduct()
	ProductsResponse := billingapi.ProductsGet200ResponseInner{
		ServerProduct: serverProduct,
	}

	actual := ProductTableFromSdk(ProductsResponse)
	assertEqualAsServerProduct(test_framework, *serverProduct, *actual)
}

// Assertions
func assertEqualAsProduct(
	test_framework *testing.T,
	sdkProduct billingapi.Product,
	cliTable ProductTable,
) {
	assert.Equal(test_framework, sdkProduct.ProductCode, cliTable.ProductCode)
	assert.Equal(test_framework, sdkProduct.ProductCategory, string(cliTable.ProductCategory))

	sdkAsTableStrings := iterutils.Map(sdkProduct.Plans, sdkPricingPlanToTableString)

	assert.Equal(test_framework, sdkAsTableStrings, cliTable.Plans)
}

func assertEqualAsServerProduct(
	test_framework *testing.T,
	sdkProduct billingapi.ServerProduct,
	cliTable ProductTable,
) {
	assert.Equal(test_framework, sdkProduct.ProductCode, cliTable.ProductCode)
	assert.Equal(test_framework, sdkProduct.ProductCategory, string(cliTable.ProductCategory))

	sdkAsTableStrings := iterutils.Map(sdkProduct.Plans, sdkPricingPlanToTableString)
	assert.Equal(test_framework, cliTable.Plans, sdkAsTableStrings)

	assert.Equal(test_framework, sdkProduct.Metadata.RamInGb, cliTable.Metadata[RAM_IN_GB])
	assert.Equal(test_framework, sdkProduct.Metadata.Cpu, cliTable.Metadata[CPU])
	assert.Equal(test_framework, sdkProduct.Metadata.CpuCount, cliTable.Metadata[CPU_COUNT])
	assert.Equal(test_framework, sdkProduct.Metadata.CoresPerCpu, cliTable.Metadata[CORES_PER_CPU])
	assert.Equal(test_framework, sdkProduct.Metadata.CpuFrequency, cliTable.Metadata[CPU_FREQUENCY])
	assert.Equal(test_framework, sdkProduct.Metadata.Network, cliTable.Metadata[NETWORK])
	assert.Equal(test_framework, sdkProduct.Metadata.Storage, cliTable.Metadata[STORAGE])
}

func sdkPricingPlanToTableString(plan billingapi.PricingPlan) string {
	return fmt.Sprintf("Sku: %s\nPrice: %f\nPrice Unit: %s", plan.Sku, plan.Price, plan.PriceUnit)
}
