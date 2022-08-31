package billingmodels

import (
	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ReservationsGetQueryParams struct {
	ProductCategory *billingapi.ProductCategoryEnum
}

func NewReservationsGetQueryParams(productCategory string) (params *ReservationsGetQueryParams, err error) {
	var validCategory *billingapi.ProductCategoryEnum

	if productCategory != "" {
		validCategory, err = billingapi.NewProductCategoryEnumFromValue(productCategory)

		if err != nil {
			return nil, fmt.Errorf("category '%s' is invalid. Allowed values are %v", productCategory, billingapi.AllowedProductCategoryEnumEnumValues)
		}
	}

	return &ReservationsGetQueryParams{
		ProductCategory: validCategory,
	}, nil
}

func (queryParams *ReservationsGetQueryParams) AttachToRequest(request billingapi.ApiReservationsGetRequest) billingapi.ApiReservationsGetRequest {
	if queryParams.ProductCategory != nil {
		request.ProductCategory(*queryParams.ProductCategory)
	}

	return request
}
