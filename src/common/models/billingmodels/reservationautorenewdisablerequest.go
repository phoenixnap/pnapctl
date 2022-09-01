package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ReservationAutoRenewDisableRequest struct {
	AutoRenewDisableReason *string `json:"autoRenewDisableReason,omitempty" yaml:"autoRenewDisableReason,omitempty"`
}

func ReservationAutoRenewDisableRequestFromSdk(sdk *billingapi.ReservationAutoRenewDisableRequest) *ReservationAutoRenewDisableRequest {
	if sdk == nil {
		return nil
	}
	return &ReservationAutoRenewDisableRequest{
		AutoRenewDisableReason: sdk.AutoRenewDisableReason,
	}
}
