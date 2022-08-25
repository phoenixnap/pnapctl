package billingmodels

import (
	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ProductsGetQueryParams struct {
	productCode     *string
	productCategory *string
	skuCode         *string
	location        *string
}

func NewProductsGetQueryParams(productCode string, productCategory string, skuCode string, location string) ProductsGetQueryParams {
	return ProductsGetQueryParams{
		productCode:     stringToParam(productCode),
		productCategory: stringToParam(productCategory),
		skuCode:         stringToParam(skuCode),
		location:        stringToParam(location),
	}
}

func (queries ProductsGetQueryParams) AttachToRequest(request *billingapisdk.ApiProductsGetRequest) {
	if queries.productCode != nil {
		request.ProductCode(*queries.productCode)
	}
	if queries.productCategory != nil {
		request.ProductCategory(*queries.productCategory)
	}
	if queries.skuCode != nil {
		request.SkuCode(*queries.skuCode)
	}
	if queries.location != nil {
		request.Location(*queries.location)
	}
}

func stringToParam(param string) *string {
	if param != "" {
		return &param
	}
	return nil
}
