package billingmodels

import (
	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ProductsGetQueryParams struct {
	ProductCode     *string
	ProductCategory *string
	SkuCode         *string
	Location        *string
}

func NewProductsGetQueryParams(productCode string, productCategory string, skuCode string, location string) ProductsGetQueryParams {
	return ProductsGetQueryParams{
		ProductCode:     stringToParam(productCode),
		ProductCategory: stringToParam(productCategory),
		SkuCode:         stringToParam(skuCode),
		Location:        stringToParam(location),
	}
}

func (queries ProductsGetQueryParams) AttachToRequest(request billingapisdk.ApiProductsGetRequest) billingapisdk.ApiProductsGetRequest {
	if queries.ProductCode != nil {
		request.ProductCode(*queries.ProductCode)
	}
	if queries.ProductCategory != nil {
		request.ProductCategory(*queries.ProductCategory)
	}
	if queries.SkuCode != nil {
		request.SkuCode(*queries.SkuCode)
	}
	if queries.Location != nil {
		request.Location(*queries.Location)
	}
	return request
}

func stringToParam(param string) *string {
	if param != "" {
		return &param
	}
	return nil
}
