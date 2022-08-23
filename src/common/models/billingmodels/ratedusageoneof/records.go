package ratedusageoneof

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

// Record types
type BandwidthRecord struct {
	RatedUsage
	Metadata BandwidthDetails `json:"metadata" yaml:"metadata"`
}

type OperatingSystemRecord struct {
	RatedUsage
	Metadata OperatingSystemDetails `json:"metadata" yaml:"metadata"`
}

type PublicSubnetRecord struct {
	RatedUsage
	Metadata PublicSubnetDetails `json:"metadata" yaml:"metadata"`
}

type ServerRecord struct {
	RatedUsage
	Metadata ServerDetails `json:"metadata" yaml:"metadata"`
}

// Metadata
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

type PublicSubnetDetails struct {
	Id   *string `json:"id,omitempty" yaml:"id,omitempty"`
	Cidr string  `json:"cidr" yaml:"cidr"`
	Size string  `json:"size" yaml:"size"`
}

func PublicSubnetDetailsFromSdk(publicSubnetDetails *billingapi.PublicSubnetDetails) *PublicSubnetDetails {
	return &PublicSubnetDetails{
		Id:   publicSubnetDetails.Id,
		Cidr: publicSubnetDetails.Cidr,
		Size: publicSubnetDetails.Size,
	}
}

type ServerDetails struct {
	Id       string `json:"id" yaml:"id"`
	Hostname string `json:"hostname" yaml:"hostname"`
}

func ServerDetailsFromSdk(serverDetails *billingapi.ServerDetails) *ServerDetails {
	return &ServerDetails{
		Id:       serverDetails.Id,
		Hostname: serverDetails.Hostname,
	}
}
