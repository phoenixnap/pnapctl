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

func ProductTableFromSdk(sdk billingapisdk.ProductsGet200ResponseInner) ProductTable {
	product := ProductTable{}
	product.parseCommon(sdk)
	product.attachMetadata(sdk)
	return product
}

func (p *ProductTable) parseCommon(sdk billingapisdk.ProductsGet200ResponseInner) {
	productCommon := productoneof.ProductCommonFromSdkOneOf(&sdk)

	p.ProductCode = productCommon.ProductCode
	p.ProductCategory = productCommon.ProductCategory
	p.Plans = iterutils.Map(productCommon.Plans, productoneof.PricingPlanToTableString)
}

func (p *ProductTable) attachMetadata(sdk billingapisdk.ProductsGet200ResponseInner) {
	switch p.ProductCategory {
	case productoneof.BANDWIDTH, productoneof.OPERATING_SYSTEM:
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
