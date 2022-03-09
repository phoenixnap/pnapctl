package tables

import ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi"

type IpBlock struct {
	Id                   string  `header:"ID"`
	Location             string  `header:"Location"`
	CidrBlockSize        string  `header:"Cidr Block Size"`
	Cidr                 string  `header:"Cidr"`
	Status               string  `header:"Status"`
	AssignedResourceId   *string `header:"Assigned Resource ID"`
	AssignedResourceType *string `header:"Assigned Resource Type"`
}

func ToIpBlockTable(ipBlock ipapisdk.IpBlock) IpBlock {
	return IpBlock{
		Id:                   ipBlock.Id,
		Location:             ipBlock.Location,
		CidrBlockSize:        ipBlock.CidrBlockSize,
		Cidr:                 ipBlock.Cidr,
		Status:               ipBlock.Status,
		AssignedResourceId:   ipBlock.AssignedResourceId,
		AssignedResourceType: ipBlock.AssignedResourceType,
	}
}
