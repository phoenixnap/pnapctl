package tables

import (
	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"phoenixnap.com/pnapctl/common/models"
)

type IpBlock struct {
	Id                   string   `header:"ID"`
	Location             string   `header:"Location"`
	CidrBlockSize        string   `header:"Cidr Block Size"`
	Cidr                 string   `header:"Cidr"`
	Status               string   `header:"Status"`
	AssignedResourceId   string   `header:"Assigned Resource ID"`
	AssignedResourceType string   `header:"Assigned Resource Type"`
	Description          string   `header:"Description"`
	Tags                 []string `header:"Tags"`
	IsBringYourOwn       bool     `header:"Bring Your Own"`
	CreatedOn            string   `header:"Created On"`
}

type IpBlockShort struct {
	Id            string   `header:"ID"`
	Location      string   `header:"Location"`
	CidrBlockSize string   `header:"Cidr Block Size"`
	Cidr          string   `header:"Cidr"`
	Status        string   `header:"Status"`
	Description   string   `header:"Description"`
	Tags          []string `header:"Tags"`
	CreatedOn     string   `header:"Created On"`
}

func ToIpBlockTable(ipBlock ipapisdk.IpBlock) IpBlock {
	var tags []string

	for _, tag := range ipBlock.Tags {
		tags = append(tags, models.TagAssignmentToTableString(&tag))
	}

	return IpBlock{
		Id:                   ipBlock.Id,
		Location:             ipBlock.Location,
		CidrBlockSize:        ipBlock.CidrBlockSize,
		Cidr:                 ipBlock.Cidr,
		Status:               ipBlock.Status,
		AssignedResourceId:   DerefString(ipBlock.AssignedResourceId),
		AssignedResourceType: DerefString(ipBlock.AssignedResourceType),
		Description:          DerefString(ipBlock.Description),
		Tags:                 tags,
		IsBringYourOwn:       ipBlock.IsBringYourOwn,
		CreatedOn:            ipBlock.CreatedOn.String(),
	}
}

func ToShortIpBlockTable(ipBlock ipapisdk.IpBlock) IpBlockShort {
	var tags []string

	for _, tag := range ipBlock.Tags {
		tags = append(tags, models.TagAssignmentToTableString(&tag))
	}

	return IpBlockShort{
		Id:            ipBlock.Id,
		Location:      ipBlock.Location,
		CidrBlockSize: ipBlock.CidrBlockSize,
		Cidr:          ipBlock.Cidr,
		Status:        ipBlock.Status,
		Description:   DerefString(ipBlock.Description),
		Tags:          tags,
		CreatedOn:     ipBlock.CreatedOn.String(),
	}
}
