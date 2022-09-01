package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ReservationRequest struct {
	Sku string `json:"sku" yaml:"sku"`
}

func ReservationRequestFromSdk(sdk billingapi.ReservationRequest) ReservationRequest {
	return ReservationRequest{
		Sku: sdk.Sku,
	}
}
