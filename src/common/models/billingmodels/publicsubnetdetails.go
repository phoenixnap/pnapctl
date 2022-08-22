package billingmodels

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

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
