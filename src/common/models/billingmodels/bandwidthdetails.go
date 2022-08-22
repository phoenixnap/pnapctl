package billingmodels

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

type BandwidthDetails struct {
}

func BandwidthDetailsFromSdk(bandwidthDetails *billingapi.BandwidthDetails) *BandwidthDetails {
	return nil
}

func BandwidthDetailsToTableString(bandwidthDetails *billingapi.BandwidthDetails) string {
	if bandwidthDetails == nil {
		return ""
	}

	return ""
}
