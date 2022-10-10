package generators

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func UpdateLocation[T interface{ SetLocation(billingapi.LocationEnum) }](item T) {
	item.SetLocation(billingapi.PHX)
}

func updatePricingPlan(sdk billingapi.PricingPlan) billingapi.PricingPlan {
	sdk.PriceUnit = billingapi.GB
	return sdk
}

func UpdatePricingPlans[T interface {
	GetPlans() []billingapi.PricingPlan
	SetPlans([]billingapi.PricingPlan)
}](item T) {
	item.SetPlans(iterutils.Map(item.GetPlans(), updatePricingPlan))
}
