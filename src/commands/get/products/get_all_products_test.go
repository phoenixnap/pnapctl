package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/cmdname"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func getQueryParams() (string, string, string, string) {
	return ProductCode, ProductCategory, SkuCode, Location
}

func TestGetAllProducts_FullTable(test_framework *testing.T) {
	responseList := generators.GenerateProductSdkList()

	var products []interface{}

	for _, product := range responseList {
		products = append(products, *tables.ProductTableFromSdk(product))
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductsGet(getQueryParams()).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(products).
		Return(nil)

	err := GetProductsCmd.RunE(GetProductsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllProducts_KeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductsGet(getQueryParams()).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetProductsCmd.RunE(GetProductsCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllProducts_PrinterFailure(test_framework *testing.T) {
	responseList := generators.GenerateProductSdkList()

	var products []interface{}

	for _, product := range responseList {
		products = append(products, *tables.ProductTableFromSdk(product))
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductsGet(getQueryParams()).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(products).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetProductsCmd.RunE(GetProductsCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}

func TestGetAllProducts_ServerError(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductsGet(getQueryParams()).
		Return(nil, WithResponse(500, nil), nil)

	err := GetProductsCmd.RunE(GetProductsCmd, []string{})

	// Assertions
	expectedMessage := "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}
