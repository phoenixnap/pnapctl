package billingmodels

import (
	"math/rand"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestProductAvailabilityGetQueryParams_valid(test_framework *testing.T) {
	productCategory := allowedProductCategories
	productCode := []string{"SRVCODE"}
	showOnlyMinQuantityAvailable := false
	location := iterutils.Map(billingapi.AllowedLocationEnumEnumValues, func(loc billingapi.LocationEnum) string {
		return string(loc)
	})
	solution := allowedSolutions
	minQuantity := rand.Float32()

	queryParams, err := NewProductAvailabilityGetQueryParams(productCategory, productCode, showOnlyMinQuantityAvailable, location, solution, minQuantity)

	assert.Nil(test_framework, err)

	assert.Equal(test_framework, productCategory, queryParams.ProductCategory)
	assert.Equal(test_framework, productCode, queryParams.ProductCode)
	assert.Equal(test_framework, showOnlyMinQuantityAvailable, queryParams.ShowOnlyMinQuantityAvailable)
	assert.Equal(test_framework, billingapi.AllowedLocationEnumEnumValues, queryParams.Location)
	assert.Equal(test_framework, solution, queryParams.Solution)
	assert.Equal(test_framework, &minQuantity, queryParams.MinQuantity)
}

func TestProductAvailabilityGetQueryParams_invalidProductCategory(test_framework *testing.T) {
	productCategory := []string{"unknown"}
	productCode := []string{"SRVCODE"}
	showOnlyMinQuantityAvailable := false
	location := iterutils.Map(billingapi.AllowedLocationEnumEnumValues, func(loc billingapi.LocationEnum) string {
		return string(loc)
	})
	solution := allowedSolutions
	minQuantity := rand.Float32()

	queryParams, err := NewProductAvailabilityGetQueryParams(productCategory, productCode, showOnlyMinQuantityAvailable, location, solution, minQuantity)

	assert.Nil(test_framework, queryParams)
	assert.EqualError(test_framework, err, "category 'unknown' is invalid. Allowed values are [SERVER]")
}
