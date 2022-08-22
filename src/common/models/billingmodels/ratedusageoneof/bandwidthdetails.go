package ratedusageoneof

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

type BandwidthDetails struct {
	IngressGb       float32  `json:"ingressGb" yaml:"ingressGb"`
	EgressGb        float32  `json:"egressGb" yaml:"egressGb"`
	PackageQuantity *float32 `json:"packageQuantity,omitempty" yaml:"packageQuantity,omitempty"`
	PackageUnit     *string  `json:"packageUnit,omitempty" yaml:"packageUnit,omitempty"`
}

func BandwidthDetailsFromSdk(bandwidthDetails *billingapi.BandwidthDetails) *BandwidthDetails {
	return &BandwidthDetails{
		IngressGb:       bandwidthDetails.IngressGb,
		EgressGb:        bandwidthDetails.EgressGb,
		PackageQuantity: bandwidthDetails.PackageQuantity,
		PackageUnit:     bandwidthDetails.PackageUnit,
	}
}
