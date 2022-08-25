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

	validProductCategory, err := parseProductCategory(productCategory)
	if err != nil {
		return nil, err
	}

	return &RatedUsageGetQueryParams{
		FromYearMonth:   fromYearMonth,
		ToYearMonth:     toYearMonth,
		ProductCategory: validProductCategory,
	}, nil
}

func (queries RatedUsageGetQueryParams) AttachToRequest(request *billingapisdk.ApiRatedUsageGetRequest) {
	request.FromYearMonth(queries.FromYearMonth)
	request.ToYearMonth(queries.ToYearMonth)

	if queries.ProductCategory != nil {
		request.ProductCategory(*queries.ProductCategory)
	}
}

type RatedUsageGetMonthToDateQueryParams struct {
	ProductCategory *billingapisdk.ProductCategoryEnum
}

func NewRatedUsageGetMonthToDateQueryParams(productCategory string) (*RatedUsageGetMonthToDateQueryParams, error) {
	validProductCategory, err := parseProductCategory(productCategory)

	if err != nil {
		return nil, err
	}

	return &RatedUsageGetMonthToDateQueryParams{
		ProductCategory: validProductCategory,
	}, nil
}

func (queries RatedUsageGetMonthToDateQueryParams) AttachToRequest(request *billingapisdk.ApiRatedUsageMonthToDateGetRequest) {
	if queries.ProductCategory != nil {
		request.ProductCategory(*queries.ProductCategory)
	}
}

// Private methods
func parseProductCategory(productCategory string) (validCategory *billingapisdk.ProductCategoryEnum, err error) {
	if productCategory != "" {
		validCategory, err = billingapisdk.NewProductCategoryEnumFromValue(productCategory)

		if err != nil {
			err = fmt.Errorf("invalid ProductCategory '%s'. Valid values: %v", productCategory, billingapisdk.AllowedProductCategoryEnumEnumValues)
		}
	}

	return
}
