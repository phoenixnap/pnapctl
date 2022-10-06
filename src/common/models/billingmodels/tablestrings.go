package billingmodels

import (
	"fmt"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

func LocationAvailabilityDetailsToTableString(sdk billingapi.LocationAvailabilityDetail) string {
	return ""
}

func ThresholdConfigurationToTableString(sdk *billingapi.ThresholdConfigurationDetails) string {
	return fmt.Sprintf("%f", sdk.GetThresholdAmount())
}

func PricingPlanToTableString(sdk billingapi.PricingPlan) string {
	return ""
}
