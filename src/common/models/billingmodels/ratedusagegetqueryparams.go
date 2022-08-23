package billingmodels

import (
	"fmt"
	"regexp"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type RatedUsageGetQueryParams struct {
	FromYearMonth   string
	ToYearMonth     string
	ProductCategory *billingapisdk.ProductCategoryEnum
}

func NewRatedUsageGetQueryParams(fromYearMonth string, toYearMonth string, productCategory string) (*RatedUsageGetQueryParams, error) {
	validYearMonth := regexp.MustCompile("[0-9]{4}-0[1-9]|1[0-2]")

	if !validYearMonth.MatchString(fromYearMonth) {
		return nil, fmt.Errorf("'FromYearMonth' (%s) is not in the valid format (YYYY-MM)", fromYearMonth)
	}
	if !validYearMonth.MatchString(toYearMonth) {
		return nil, fmt.Errorf("'ToYearMonth' (%s) is not in the valid format (YYYY-MM)", toYearMonth)
	}

	var validProductCategory *billingapisdk.ProductCategoryEnum

	if productCategory != "" {
		parsedEnum, err := billingapisdk.NewProductCategoryEnumFromValue(productCategory)
		validProductCategory = parsedEnum

		if err != nil {
			return nil, fmt.Errorf("invalid ProductCategory '%s'. Valid values: %v", productCategory, billingapisdk.AllowedProductCategoryEnumEnumValues)
		}
	}

	return &RatedUsageGetQueryParams{
		FromYearMonth:   fromYearMonth,
		ToYearMonth:     toYearMonth,
		ProductCategory: validProductCategory,
	}, nil
}

func (queries RatedUsageGetQueryParams) AttachToRequest(request billingapisdk.ApiRatedUsageGetRequest) *billingapisdk.ApiRatedUsageGetRequest {
	request = request.FromYearMonth(queries.FromYearMonth)
	request = request.ToYearMonth(queries.ToYearMonth)

	if queries.ProductCategory != nil {
		request = request.ProductCategory(*queries.ProductCategory)
	}

	return &request
}
