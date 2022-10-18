package product

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

/*
One of types
*/
var (
	Bandwidth       = "BANDWIDTH"
	OperatingSystem = "OPERATING_SYSTEM"
	Server          = "SERVER"
	Storage         = "STORAGE"
)

func BandwidthProductToInner(product billingapi.Product) billingapi.ProductsGet200ResponseInner {
	product.ProductCategory = Bandwidth
	return billingapi.ProductAsProductsGet200ResponseInner(&product)
}

func OperatingSystemProductToInner(product billingapi.Product) billingapi.ProductsGet200ResponseInner {
	product.ProductCategory = OperatingSystem
	return billingapi.ProductAsProductsGet200ResponseInner(&product)
}

func StorageProductToInner(product billingapi.Product) billingapi.ProductsGet200ResponseInner {
	product.ProductCategory = Storage
	return billingapi.ProductAsProductsGet200ResponseInner(&product)
}

func ServerProductToInner(serverProduct billingapi.ServerProduct) billingapi.ProductsGet200ResponseInner {
	serverProduct.ProductCategory = Server
	return billingapi.ServerProductAsProductsGet200ResponseInner(&serverProduct)
}
