package billingmodels

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

type PublicSubnetDetails struct {
}

func PublicSubnetDetailsFromSdk(publicSubnetDetails *billingapi.PublicSubnetDetails) *PublicSubnetDetails {
	return nil
}

func PublicSubnetDetailsToTableString(publicSubnetDetails *billingapi.PublicSubnetDetails) string {
	if publicSubnetDetails == nil {
		return ""
	}

	return ""
}
