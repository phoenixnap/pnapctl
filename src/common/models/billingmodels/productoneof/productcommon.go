package productoneof

import (
	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

// Uses the discriminator - productCategory
type Discriminator string

const (
	// ServerProduct
	SERVER Discriminator = "SERVER"

	// Product
	BANDWIDTH        Discriminator = "BANDWIDTH"
	OPERATING_SYSTEM Discriminator = "OPERATING_SYSTEM"
)

type ProductSdk interface {
	GetProductCode() string
	GetProductCategory() string
	GetPlans() []billingapi.PricingPlan
}

type ProductCommon struct {
	ProductCode     string
	ProductCategory Discriminator
	Plans           []PricingPlan
}

func ProductCommonFromSdkOneOf(sdk *billingapi.ProductsGet200ResponseInner) *ProductCommon {
	if sdk == nil {
		return nil
	}

	actual := sdk.GetActualInstance().(ProductSdk)

	return &ProductCommon{
		ProductCode:     actual.GetProductCode(),
		ProductCategory: Discriminator(actual.GetProductCategory()),
		Plans:           iterutils.Map(actual.GetPlans(), PricingPlanFromSdk),
	}
}

// Checks whether the struct type is actually any of the ones passed.
func (c *ProductCommon) IsActually(ds ...Discriminator) bool {
	return iterutils.Contains(ds, c.ProductCategory)
}

// Nested types
type PricingPlan struct {
	Sku                   string                   `json:"sku" yaml:"sku"`
	SkuDescription        *string                  `json:"skuDescription,omitempty" yaml:"skuDescription,omitempty"`
	Location              string                   `json:"location" yaml:"location"`
	PricingModel          string                   `json:"pricingModel" yaml:"pricingModel"`
	Price                 float32                  `json:"price" yaml:"price"`
	PriceUnit             billingapi.PriceUnitEnum `json:"priceUnit" yaml:"priceUnit"`
	CorrelatedProductCode *string                  `json:"correlatedProductCode,omitempty" yaml:"correlatedProductCode,omitempty"`
	PackageQuantity       *float32                 `json:"packageQuantity,omitempty" yaml:"packageQuantity,omitempty"`
	PackageUnit           *string                  `json:"packageUnit,omitempty" yaml:"packageUnit,omitempty"`
}

func PricingPlanFromSdk(sdk billingapi.PricingPlan) PricingPlan {
	return PricingPlan{
		Sku:                   sdk.Sku,
		SkuDescription:        sdk.SkuDescription,
		Location:              sdk.Location,
		PricingModel:          sdk.PricingModel,
		Price:                 sdk.Price,
		PriceUnit:             sdk.PriceUnit,
		CorrelatedProductCode: sdk.CorrelatedProductCode,
		PackageQuantity:       sdk.PackageQuantity,
		PackageUnit:           sdk.PackageUnit,
	}
}

func PricingPlanToTableString(plan PricingPlan) string {
	return fmt.Sprintf("Sku: %s\nPrice: %f\nPrice Unit: %s", plan.Sku, plan.Price, plan.PriceUnit)
}
