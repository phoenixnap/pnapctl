package billingmodels

import (
	"math/rand"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	"phoenixnap.com/pnapctl/testsupport/testutil"
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
	location := testutil.AsStrings(billingapi.AllowedLocationEnumEnumValues)
	solution := allowedSolutions
	minQuantity := rand.Float32()

	queryParams, err := NewProductAvailabilityGetQueryParams(productCategory, productCode, showOnlyMinQuantityAvailable, location, solution, minQuantity)

	assert.Nil(test_framework, queryParams)
	assert.EqualError(test_framework, err, "category 'unknown' is invalid. Allowed values are [SERVER]")
}

func TestProductAvailabilityGetQueryParams_invalidLocation(test_framework *testing.T) {
	productCategory := allowedProductCategories
	productCode := []string{"SRVCODE"}
	showOnlyMinQuantityAvailable := false
	location := []string{"unknown"}
	solution := allowedSolutions
	minQuantity := rand.Float32()

	queryParams, err := NewProductAvailabilityGetQueryParams(productCategory, productCode, showOnlyMinQuantityAvailable, location, solution, minQuantity)

	assert.Nil(test_framework, queryParams)
	assert.EqualError(test_framework, err, "location 'unknown' is invalid. Allowed values are [PHX ASH NLD SGP CHI SEA AUS]")
}

func TestProductAvailabilityGetQueryParams_invalidSolution(test_framework *testing.T) {
	productCategory := allowedProductCategories
	productCode := []string{"SRVCODE"}
	showOnlyMinQuantityAvailable := false
	location := testutil.AsStrings(billingapi.AllowedLocationEnumEnumValues)
	solution := []string{"unknown"}
	minQuantity := rand.Float32()

	queryParams, err := NewProductAvailabilityGetQueryParams(productCategory, productCode, showOnlyMinQuantityAvailable, location, solution, minQuantity)

	assert.Nil(test_framework, queryParams)
	assert.EqualError(test_framework, err, "solution 'unknown' is invalid. Allowed values are [SERVER_RANCHER]")
}
