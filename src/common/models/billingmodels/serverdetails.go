package billingmodels

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

type ServerDetails struct {
}

func ServerDetailsFromSdk(serverDetails *billingapi.ServerDetails) *ServerDetails {
	return nil
}

func ServerDetailsToTableString(serverDetails *billingapi.ServerDetails) string {
	if serverDetails == nil {
		return ""
	}

	return ""
}
