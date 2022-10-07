package generators

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type HasProductCategory interface {
	GetProductCategory() string
	SetProductCategory(string)
}

func updateProductCategory[T HasProductCategory](category string) func(T) {
	return func(item T) {
		item.SetProductCategory(category)
	}
}

type HasPricingPlans interface {
	GetPlans() []billingapi.PricingPlan
	SetPlans([]billingapi.PricingPlan)
}

func updatePlans[T HasPricingPlans](priceUnit string) func(T) {
	return func(item T) {
		item.SetPlans(iterutils.Map(item.GetPlans(), func(plan billingapi.PricingPlan) billingapi.PricingPlan {
			plan.PriceUnit = billingapi.GB
			return plan
		}))
	}
}

type HasLocation interface {
	GetLocation() billingapi.LocationEnum
	SetLocation(billingapi.LocationEnum)
}

func updateLocation[T HasLocation](location string) func(T) {
	return func(item T) {
		item.SetLocation(billingapi.PHX)
	}
}
