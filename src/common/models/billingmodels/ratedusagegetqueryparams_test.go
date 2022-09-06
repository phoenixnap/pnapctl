package billingmodels

import (
	"fmt"
	"testing"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
)

func TestRatedUsageGetQueryParams_valid(test_framework *testing.T) {
	from := "2020-11"
	to := "2021-11"
	productCategory := billingapisdk.SERVER

	queryParams, err := NewRatedUsageGetQueryParams(from, to, string(productCategory))

	assert.NotNil(test_framework, queryParams)
	assert.Nil(test_framework, err)

	assert.Equal(test_framework, from, queryParams.FromYearMonth)
	assert.Equal(test_framework, to, queryParams.ToYearMonth)
	assert.Equal(test_framework, &productCategory, queryParams.ProductCategory)
}

func TestRatedUsageGetQueryParams_invalidFrom(test_framework *testing.T) {
	from := "2020-18"
	to := "2021-11"
	productCategory := billingapisdk.SERVER

	queryParams, err := NewRatedUsageGetQueryParams(from, to, string(productCategory))

	assert.Nil(test_framework, queryParams)
	assert.NotNil(test_framework, err)

	expectedErr := fmt.Errorf("'FromYearMonth' (%s) is not in the valid format (YYYY-MM)", from)

	assert.Equal(test_framework, err, expectedErr)
}

func TestRatedUsageGetQueryParams_invalidTo(test_framework *testing.T) {
	from := "2020-11"
	to := "2021-18"
	productCategory := billingapisdk.SERVER

	queryParams, err := NewRatedUsageGetQueryParams(from, to, string(productCategory))

	assert.Nil(test_framework, queryParams)
	assert.NotNil(test_framework, err)

	expectedErr := fmt.Errorf("'ToYearMonth' (%s) is not in the valid format (YYYY-MM)", to)

	assert.Equal(test_framework, err, expectedErr)
}

func TestRatedUsageGetQueryParams_invalidProductCategory(test_framework *testing.T) {
	from := "2020-11"
	to := "2021-11"
	var productCategory billingapisdk.ProductCategoryEnum = "NotValid"

	queryParams, err := NewRatedUsageGetQueryParams(from, to, string(productCategory))

	assert.Nil(test_framework, queryParams)
	assert.NotNil(test_framework, err)

	expectedErr := ctlerrors.InvalidFlagValuePassedError("category", "NotValid", billingapisdk.AllowedProductCategoryEnumEnumValues)

	assert.Equal(test_framework, expectedErr, err)
}

// Month to date
func TestRatedUsageMonthToDateGetQueryParams_valid(test_framework *testing.T) {
	productCategory := billingapisdk.SERVER

	queryParams, err := NewRatedUsageGetMonthToDateQueryParams(string(productCategory))

	assert.NotNil(test_framework, queryParams)
	assert.Nil(test_framework, err)

	assert.Equal(test_framework, &productCategory, queryParams.ProductCategory)
}

func TestRatedUsageMonthToDateGetQueryParams_invalidProductCategory(test_framework *testing.T) {
	var productCategory billingapisdk.ProductCategoryEnum = "NotValid"

	queryParams, err := NewRatedUsageGetMonthToDateQueryParams(string(productCategory))

	assert.Nil(test_framework, queryParams)
	assert.NotNil(test_framework, err)

	expectedErr := ctlerrors.InvalidFlagValuePassedError("category", "NotValid", billingapisdk.AllowedProductCategoryEnumEnumValues)

	assert.Equal(test_framework, expectedErr, err)
}
