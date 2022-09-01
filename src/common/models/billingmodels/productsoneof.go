package billingmodels

import (
	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	. "phoenixnap.com/pnapctl/common/models/billingmodels/productoneof"
)

func ProductActualFromSdk(productOneOf billingapisdk.ProductsGet200ResponseInner) interface{} {
	productCommon := ProductCommonFromSdkOneOf(&productOneOf)

	if productCommon == nil {
		return nil
	}

	// Product is identical to ProductCommon for now - so we'll initialize here.
	productCli := Product{
		ProductCommon: *productCommon,
	}

	switch {
	case productCommon.IsActually(BANDWIDTH, OPERATING_SYSTEM):
		return &productCli
	case productCommon.IsActually(SERVER):
		return &ServerProduct{
			Product:  productCli,
			Metadata: ServerProductMetadataFromSdk(productOneOf.ServerProduct.Metadata),
		}
	}

	return nil
}
