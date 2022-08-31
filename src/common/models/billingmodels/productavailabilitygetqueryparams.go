package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

// Retrieved from API spec
var allowedProductCategories = []string{
	"SERVER",
}
var allowedSolutions = []string{
	"SERVER_RANCHER",
}

type ProductAvailabilityGetQueryParams struct {
	ProductCategory              []string
	ProductCode                  []string
	ShowOnlyMinQuantityAvailable bool
	Location                     []billingapi.LocationEnum
	Solution                     []string
	MinQuantity                  *float32
}

func NewProductAvailabilityGetQueryParams(
	productCategory []string,
	productCode []string,
	showOnlyMinQuantityAvailable bool, // default true
	location []string,
	solution []string,
	minQuantity *float32,
) (queryParams *ProductAvailabilityGetQueryParams, err error) {

	if invalid := iterutils.FindElementThat(productCategory, productCategoryIsNotAllowed); invalid != nil {
		return nil, ctlerrors.InvalidFlagValuePassedError("category", *invalid, allowedProductCategories)
	}

	if invalid := iterutils.FindElementThat(location, locationIsNotAllowed); invalid != nil {
		return nil, ctlerrors.InvalidFlagValuePassedError("location", *invalid, billingapi.AllowedLocationEnumEnumValues)
	}

	parsedLocation := iterutils.Map(location, func(str string) billingapi.LocationEnum {
		return billingapi.LocationEnum(str)
	})

	if invalid := iterutils.FindElementThat(solution, solutionIsNotAllowed); invalid != nil {
		return nil, ctlerrors.InvalidFlagValuePassedError("solution", *invalid, allowedSolutions)
	}

	return &ProductAvailabilityGetQueryParams{
		ProductCategory:              productCategory,
		ProductCode:                  productCode,
		ShowOnlyMinQuantityAvailable: showOnlyMinQuantityAvailable,
		Location:                     parsedLocation,
		Solution:                     solution,
		MinQuantity:                  minQuantity,
	}, nil
}

func (queryParams *ProductAvailabilityGetQueryParams) AttachToRequest(request billingapi.ApiProductAvailabilityGetRequest) billingapi.ApiProductAvailabilityGetRequest {
	if queryParams.ProductCategory != nil {
		request.ProductCategory(queryParams.ProductCategory)
	}
	if queryParams.ProductCode != nil {
		request.ProductCode(queryParams.ProductCode)
	}
	request.ShowOnlyMinQuantityAvailable(queryParams.ShowOnlyMinQuantityAvailable)
	if queryParams.Location != nil {
		request.Location(queryParams.Location)
	}
	return request
}

// Predicates
func productCategoryIsNotAllowed(category string) bool {
	return !iterutils.Contains(allowedProductCategories, category)
}
func locationIsNotAllowed(location string) bool {
	return !billingapi.LocationEnum(location).IsValid()
}
func solutionIsNotAllowed(solution string) bool {
	return !iterutils.Contains(allowedSolutions, solution)
}
