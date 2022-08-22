package billingmodels

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

type OperatingSystemDetails struct {
}

func OperatingSystemDetailsFromSdk(operatingSystemDetails *billingapi.OperatingSystemDetails) *OperatingSystemDetails {
	return nil
}

func OperatingSystemDetailsToTableString(operatingSystemDetails *billingapi.OperatingSystemDetails) string {
	if operatingSystemDetails == nil {
		return ""
	}

	return ""
}
