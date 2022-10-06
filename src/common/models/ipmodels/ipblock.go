package ipmodels

import (
	"time"

	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
)

type IpBlock struct {
	Id                   string                   `yaml:"id" json:"id"`
	Location             string                   `yaml:"location" json:"location"`
	CidrBlockSize        string                   `yaml:"cidrBlockSize" json:"cidrBlockSize"`
	Cidr                 string                   `yaml:"cidr" json:"cidr"`
	Status               string                   `yaml:"status" json:"status"`
	AssignedResourceId   *string                  `yaml:"assignedResourceId,omitempty" json:"assignedResourceId,omitempty"`
	AssignedResourceType *string                  `yaml:"assignedResourceType,omitempty" json:"assignedResourceType,omitempty"`
	Description          *string                  `yaml:"description,omitempty" json:"description,omitempty"`
	Tags                 []ipapisdk.TagAssignment `json:"tags,omitempty"`
	IsBringYourOwn       bool                     `json:"isBringYourOwn"`
	CreatedOn            time.Time                `json:"createdOn"`
}

func IpBlockFromSdk(ipBlock ipapisdk.IpBlock) IpBlock {
	return IpBlock{
		Id:                   ipBlock.Id,
		Location:             ipBlock.Location,
		CidrBlockSize:        ipBlock.CidrBlockSize,
		Cidr:                 ipBlock.Cidr,
		Status:               ipBlock.Status,
		AssignedResourceId:   ipBlock.AssignedResourceId,
		AssignedResourceType: ipBlock.AssignedResourceType,
		Description:          ipBlock.Description,
		Tags:                 ipBlock.Tags,
	}
}
