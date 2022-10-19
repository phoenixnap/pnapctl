package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func getQueryParams() (string, string, string, string) {
	return ProductCode, ProductCategory, SkuCode, Location
}

func TestGetAllProducts_FullTable(test_framework *testing.T) {
	responseList := generators.GenerateProductSdkList()
	products := iterutils.DerefInterface(iterutils.MapInterface(responseList, tables.ProductTableFromSdk))

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductsGet(getQueryParams()).
		Return(responseList, nil)

	ExpectToPrintSuccess(test_framework, products)

	err := GetProductsCmd.RunE(GetProductsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllProducts_ClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductsGet(getQueryParams()).
		Return(nil, testutil.TestError)

	err := GetProductsCmd.RunE(GetProductsCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestError, err)
}

func TestGetAllProducts_PrinterFailure(test_framework *testing.T) {
	responseList := generators.GenerateProductSdkList()
	products := iterutils.DerefInterface(iterutils.MapInterface(responseList, tables.ProductTableFromSdk))

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductsGet(getQueryParams()).
		Return(responseList, nil)

	expectedErr := ExpectToPrintFailure(test_framework, products)

	err := GetProductsCmd.RunE(GetProductsCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
