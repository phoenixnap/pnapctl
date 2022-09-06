package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestReservationsGetQueryParamsValid_Success(test_framework *testing.T) {
	iterutils.Each(billingapi.AllowedProductCategoryEnumEnumValues,
		func(category billingapi.ProductCategoryEnum) {
			queryParams, err := NewReservationsGetQueryParams(string(category))

			assert.NoError(test_framework, err)
			assert.Equal(test_framework, category, *queryParams.ProductCategory)
		})
}

func TestReservationsGetQueryParamsEmpty_Success(test_framework *testing.T) {
	queryParams, err := NewReservationsGetQueryParams("")

	assert.NoError(test_framework, err)
	assert.Nil(test_framework, queryParams.ProductCategory)
}

func TestReservationsGetQueryParamsInvalid_Failure(test_framework *testing.T) {
	queryParams, err := NewReservationsGetQueryParams("INVALID")

	assert.Nil(test_framework, queryParams)
	assert.Error(test_framework,
		ctlerrors.InvalidFlagValuePassedError(
			"category",
			"INVALID",
			billingapi.AllowedProductCategoryEnumEnumValues,
		), err)
}
