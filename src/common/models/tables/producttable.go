package tables

import (
	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi/v2"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/oneof/product"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

const (
	// Server Product
	RAM_IN_GB     = "Ram (GB)"
	CPU           = "CPU"
	CPU_COUNT     = "CPU Count"
	CORES_PER_CPU = "Cores per CPU"
	CPU_FREQUENCY = "CPU Frequency"
	NETWORK       = "Network"
	STORAGE       = "Storage"
)

type ProductTable struct {
	// Common
	ProductCode     string   `header:"Product Code"`
	ProductCategory string   `header:"Product Category"`
	Plans           []string `header:"Plans"`

	// ServerProduct
	Metadata map[string]interface{}
}

func ProductTableFromSdk(sdk billingapisdk.ProductsGet200ResponseInner) *ProductTable {
	product := parseCommonProduct(sdk)
	if product == nil {
		return nil
	}

	product.attachUnique(sdk)
	return product
}

func parseCommonProduct(sdk billingapisdk.ProductsGet200ResponseInner) *ProductTable {
	product := models.GetFromAllOf[billingapisdk.Product](sdk)
	if product == nil {
		return nil
	}

	return &ProductTable{
		ProductCode:     product.ProductCode,
		ProductCategory: product.ProductCategory,
		Plans:           iterutils.MapRef(product.Plans, models.PricingPlanToTableString),
	}
}

func (p *ProductTable) attachUnique(sdk billingapisdk.ProductsGet200ResponseInner) {
	switch p.ProductCategory {
	case product.Bandwidth, product.OperatingSystem:
		return
	case product.Server:
		p.Metadata = map[string]interface{}{
			RAM_IN_GB:     sdk.ServerProduct.Metadata.RamInGb,
			CPU:           sdk.ServerProduct.Metadata.Cpu,
			CPU_COUNT:     sdk.ServerProduct.Metadata.CpuCount,
			CORES_PER_CPU: sdk.ServerProduct.Metadata.CoresPerCpu,
			CPU_FREQUENCY: sdk.ServerProduct.Metadata.CpuFrequency,
			NETWORK:       sdk.ServerProduct.Metadata.Network,
			STORAGE:       sdk.ServerProduct.Metadata.Storage,
		}
	}
}
