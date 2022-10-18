package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/queryparams/billing"
	"phoenixnap.com/pnapctl/common/models/tables"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllProducts_FullTable(test_framework *testing.T) {
	responseList := generators.GenerateProductSdkList()
	queryParams := generators.Generate[billing.ProductsGetQueryParams]()
	setQueryParams(queryParams)

	var products []interface{}

	for _, product := range responseList {
		products = append(products, *tables.ProductTableFromSdk(product))
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductsGet(queryParams).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(products, "get products").
		Return(nil)

	err := GetProductsCmd.RunE(GetProductsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllProducts_KeycloakFailure(test_framework *testing.T) {
	queryParams := generators.Generate[billing.ProductsGetQueryParams]()
	setQueryParams(queryParams)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductsGet(queryParams).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetProductsCmd.RunE(GetProductsCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllProducts_PrinterFailure(test_framework *testing.T) {
	responseList := generators.GenerateProductSdkList()
	queryParams := generators.Generate[billing.ProductsGetQueryParams]()
	setQueryParams(queryParams)

	var products []interface{}

	for _, product := range responseList {
		products = append(products, *tables.ProductTableFromSdk(product))
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductsGet(queryParams).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(products, "get products").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetProductsCmd.RunE(GetProductsCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}

func TestGetAllProducts_ServerError(test_framework *testing.T) {
	queryParams := generators.Generate[billing.ProductsGetQueryParams]()
	setQueryParams(queryParams)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductsGet(queryParams).
		Return(nil, WithResponse(500, nil), nil)

	err := GetProductsCmd.RunE(GetProductsCmd, []string{})

	// Assertions
	expectedMessage := "Command 'get products' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func setQueryParams(queryparams billing.ProductsGetQueryParams) {
	ProductCategory = *queryparams.ProductCategory
	ProductCode = *queryparams.ProductCode
	SkuCode = *queryparams.SkuCode
	Location = *queryparams.Location
}
