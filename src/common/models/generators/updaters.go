package generators

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

/*
	UPDATERS are used to modify generated data for testing, in order to make it valid.
	Most of the time, they're needed to set ENUMS.
	Since ENUMS in Go are normal strings, the generator makes up a random string - which is
		almost always invalid. As a result, we use UPDATERS to reset them back to a valid value.
*/

/*
	UPDATER GROUPS
	these group multiple updaters together to avoid duplicating code.
	this makes them cleaner to use, at the cost of being less generic.
*/

func UpdateRatedUsageRecord[T interface {
	SetLocation(billingapi.LocationEnum)
	GetDiscountDetails() billingapi.DiscountDetails
	SetDiscountDetails(billingapi.DiscountDetails)
	GetCreditDetails() []billingapi.CreditDetails
	SetCreditDetails([]billingapi.CreditDetails)
}](item T) {
	UpdateLocation(item)
	UpdateDiscountDetails(item)
	UpdateCreditDetailsList(item)
}

/*
	GENERIC UPDATERS
	these are public updaters that are meant to be as generic as possible.
	for example, UpdateLocation works on any struct with the "Location" property of type "billingapi.LocationEnum".
	this way they are reusable across multiple types/structs.
*/

func UpdateLocation[T interface {
	SetLocation(billingapi.LocationEnum)
}](item T) {
	item.SetLocation(billingapi.LOCATIONENUM_PHX)
}

func UpdatePricingPlans[T interface {
	GetPlans() []billingapi.PricingPlan
	SetPlans([]billingapi.PricingPlan)
}](item T) {
	item.SetPlans(iterutils.Map(item.GetPlans(), updatePricingPlan))
}

func UpdateApplicableDiscounts[T interface {
	GetApplicableDiscounts() billingapi.ApplicableDiscounts
	SetApplicableDiscounts(billingapi.ApplicableDiscounts)
}](item T) {
	discounts := item.GetApplicableDiscounts()
	UpdateDiscountDetailsList(&discounts)
	item.SetApplicableDiscounts(discounts)
}

func UpdateDiscountDetails[T interface {
	GetDiscountDetails() billingapi.DiscountDetails
	SetDiscountDetails(billingapi.DiscountDetails)
}](item T) {
	discountDetails := updateDiscountDetails(item.GetDiscountDetails())
	item.SetDiscountDetails(discountDetails)
}

func UpdateDiscountDetailsList[T interface {
	GetDiscountDetails() []billingapi.DiscountDetails
	SetDiscountDetails([]billingapi.DiscountDetails)
}](item T) {
	item.SetDiscountDetails(iterutils.Map(item.GetDiscountDetails(), updateDiscountDetails))
}

func UpdateCreditDetailsList[T interface {
	GetCreditDetails() []billingapi.CreditDetails
	SetCreditDetails([]billingapi.CreditDetails)
}](item T) {
	item.SetCreditDetails(iterutils.Map(item.GetCreditDetails(), func(item billingapi.CreditDetails) billingapi.CreditDetails {
		item.PromoCreditDetails.Type = billingapi.CREDITTYPEENUM_PROMO
		item.SystemCreditDetails.Type = billingapi.CREDITTYPEENUM_SYSTEM
		return item
	}))
}

/*
	DIRECT UPDATERS
	these are used privately, and work with SPECIFIC types.
	for example "updatePricingPlan" only works on "billingapi.PricingPlan".
*/

func updatePricingPlan(sdk billingapi.PricingPlan) billingapi.PricingPlan {
	sdk.PriceUnit = billingapi.PRICEUNITENUM_GB
	sdk.PackageUnit = billingapi.PACKAGEUNITENUM_GB.Ptr()
	UpdateApplicableDiscounts(&sdk)

	return sdk
}

func updateDiscountDetails(sdk billingapi.DiscountDetails) billingapi.DiscountDetails {
	sdk.Type = billingapi.DISCOUNTTYPEENUM_GLOBAL_PERCENTAGE
	return sdk
}
