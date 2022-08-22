package ratedusageoneof

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

type OperatingSystemDetails struct {
	Cores         int32  `json:"cores" yaml:"cores"`
	CorrelationId string `json:"correlationId" yaml:"correlationId"`
}

func OperatingSystemDetailsFromSdk(operatingSystemDetails *billingapi.OperatingSystemDetails) *OperatingSystemDetails {
	return &OperatingSystemDetails{
		Cores:         operatingSystemDetails.Cores,
		CorrelationId: operatingSystemDetails.CorrelationId,
	}
}
