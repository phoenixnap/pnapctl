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

func NewRatedUsageGetQueryParams(fromYearMonth string, toYearMonth string, productCategory string) (params *RatedUsageGetQueryParams, err error) {
	validYearMonth := regexp.MustCompile("[0-9]{4}-0[1-9]|1[0-2]")

	if !validYearMonth.MatchString(fromYearMonth) {
		err = fmt.Errorf("'FromYearMonth' (%s) is not in the valid format (YYYY-MM)", fromYearMonth)
		return
	}
	if !validYearMonth.MatchString(toYearMonth) {
		err = fmt.Errorf("'ToYearMonth' (%s) is not in the valid format (YYYY-MM)", toYearMonth)
		return
	}

	validProductCategory, err := parseProductCategory(productCategory)

	params = &RatedUsageGetQueryParams{
		FromYearMonth:   fromYearMonth,
		ToYearMonth:     toYearMonth,
		ProductCategory: validProductCategory,
	}

	return
}

func (queries RatedUsageGetQueryParams) AttachToRequest(request billingapisdk.ApiRatedUsageGetRequest) *billingapisdk.ApiRatedUsageGetRequest {
	request = request.FromYearMonth(queries.FromYearMonth)
	request = request.ToYearMonth(queries.ToYearMonth)

	if queries.ProductCategory != nil {
		request = request.ProductCategory(*queries.ProductCategory)
	}

	return &request
}

type RatedUsageGetMonthToDateQueryParams struct {
	ProductCategory *billingapisdk.ProductCategoryEnum
}

func NewRatedUsageGetMonthToDateQueryParams(productCategory string) (params *RatedUsageGetMonthToDateQueryParams, err error) {
	validProductCategory, err := parseProductCategory(productCategory)

	params = &RatedUsageGetMonthToDateQueryParams{
		ProductCategory: validProductCategory,
	}

	return
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
