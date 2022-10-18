package billing

import (
	"fmt"
	"regexp"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/ctlerrors"
)

type RatedUsageGetQueryParams struct {
	FromYearMonth   string
	ToYearMonth     string
	ProductCategory *billingapisdk.ProductCategoryEnum
}

func NewRatedUsageGetQueryParams(fromYearMonth string, toYearMonth string, productCategory string) (*RatedUsageGetQueryParams, error) {
	validYearMonth := regexp.MustCompile("^([0-9]{4})(-){1}(0[1-9]|1[0-2]){1}$")

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

func (queries RatedUsageGetQueryParams) AttachToRequest(request billingapisdk.ApiRatedUsageGetRequest) billingapisdk.ApiRatedUsageGetRequest {
	request = request.FromYearMonth(queries.FromYearMonth)
	request = request.ToYearMonth(queries.ToYearMonth)

	if queries.ProductCategory != nil {
		request = request.ProductCategory(*queries.ProductCategory)
	}

	return request
}

type RatedUsageMonthToDateGetQueryParams struct {
	ProductCategory *billingapisdk.ProductCategoryEnum
}

func NewRatedUsageGetMonthToDateQueryParams(productCategory string) (*RatedUsageMonthToDateGetQueryParams, error) {
	validProductCategory, err := parseProductCategory(productCategory)

	if err != nil {
		return nil, err
	}

	return &RatedUsageMonthToDateGetQueryParams{
		ProductCategory: validProductCategory,
	}, nil
}

func (queries RatedUsageMonthToDateGetQueryParams) AttachToRequest(request billingapisdk.ApiRatedUsageMonthToDateGetRequest) billingapisdk.ApiRatedUsageMonthToDateGetRequest {
	if queries.ProductCategory != nil {
		request = request.ProductCategory(*queries.ProductCategory)
	}
	return request
}

// Private methods
func parseProductCategory(productCategory string) (validCategory *billingapisdk.ProductCategoryEnum, err error) {
	if productCategory != "" {
		validCategory, err = billingapisdk.NewProductCategoryEnumFromValue(productCategory)

		if err != nil {
			err = ctlerrors.InvalidFlagValuePassedError("category", productCategory, billingapisdk.AllowedProductCategoryEnumEnumValues)
		}
	}

	return
}
