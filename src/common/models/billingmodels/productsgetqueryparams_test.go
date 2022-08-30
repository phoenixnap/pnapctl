package billingmodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductsGetQueryParams_allSet(test_framework *testing.T) {
	productCode := "CODE"
	productCategory := "CATEGORY"
	skuCode := "SKU CODE"
	location := "LOCATION"

	queryParams := NewProductsGetQueryParams(productCode, productCategory, skuCode, location)

	assert.Equal(test_framework, &productCode, queryParams.ProductCode)
	assert.Equal(test_framework, &productCategory, queryParams.ProductCategory)
	assert.Equal(test_framework, &skuCode, queryParams.SkuCode)
	assert.Equal(test_framework, &location, queryParams.Location)
}

func TestProductsGetQueryParams_noneSet(test_framework *testing.T) {
	productCode := ""
	productCategory := ""
	skuCode := ""
	location := ""

	queryParams := NewProductsGetQueryParams(productCode, productCategory, skuCode, location)

	assert.Nil(test_framework, queryParams.ProductCode)
	assert.Nil(test_framework, queryParams.ProductCategory)
	assert.Nil(test_framework, queryParams.SkuCode)
	assert.Nil(test_framework, queryParams.Location)
}
