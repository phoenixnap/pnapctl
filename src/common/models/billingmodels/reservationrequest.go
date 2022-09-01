package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ReservationRequest struct {
	Sku string `json:"sku" yaml:"sku"`
}

func ReservationRequestFromSdk(sdk *billingapi.ReservationRequest) *ReservationRequest {
	if sdk == nil {
		return nil
	}
	return &ReservationRequest{
		Sku: sdk.Sku,
	}
}
