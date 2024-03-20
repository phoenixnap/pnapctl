package generators

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func UpdateLocation[T interface {
	SetLocation(billingapi.LocationEnum)
}](item T) {
	item.SetLocation(billingapi.LOCATIONENUM_PHX)
}

func updatePricingPlan(sdk billingapi.PricingPlan) billingapi.PricingPlan {
	sdk.PriceUnit = billingapi.PRICEUNITENUM_GB
	sdk.PackageUnit = billingapi.PACKAGEUNITENUM_GB.Ptr()
	return sdk
}

func UpdatePricingPlans[T interface {
	GetPlans() []billingapi.PricingPlan
	SetPlans([]billingapi.PricingPlan)
}](item T) {
	item.SetPlans(iterutils.Map(item.GetPlans(), updatePricingPlan))
}
