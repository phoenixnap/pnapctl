package tables

import (
	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels/productoneof"
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
	ProductCode     string                     `header:"Product Code"`
	ProductCategory productoneof.Discriminator `header:"Product Category"`
	Plans           []string                   `header:"Plans"`

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
	productCommon := productoneof.ProductCommonFromSdkOneOf(&sdk)

	if productCommon == nil {
		return nil
	}

	return &ProductTable{
		ProductCode:     productCommon.ProductCode,
		ProductCategory: productCommon.ProductCategory,
		Plans:           iterutils.Map(productCommon.Plans, productoneof.PricingPlanToTableString),
	}
}

func (p *ProductTable) attachUnique(sdk billingapisdk.ProductsGet200ResponseInner) {
	switch p.ProductCategory {
	case productoneof.BANDWIDTH, productoneof.OPERATING_SYSTEM, productoneof.STORAGE:
		return
	case productoneof.SERVER:
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
