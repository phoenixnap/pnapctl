package tables

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"phoenixnap.com/pnapctl/common/models"
)

type TagTable struct {
	Id                  string   `header:"Id"`
	Name                string   `header:"Name"`
	Values              []string `header:"Values"`
	Description         string   `header:"Description"`
	IsBillingTag        bool     `header:"Is Billing Tag"`
	ResourceAssignments []string `header:"Resource Assignments"`
	CreatedBy           string   `header:"Created By"`
}

func TagFromSdk(tag tagapisdk.Tag) TagTable {
	var resourceAssignments []string

	if tag.ResourceAssignments != nil {
		for _, x := range tag.ResourceAssignments {
			resourceAssignments = append(resourceAssignments, models.ResourceAssignmentToTableString(x))
		}
	}

	return TagTable{
		Id:                  tag.Id,
		Name:                tag.Name,
		Values:              tag.Values,
		Description:         DerefString(tag.Description),
		IsBillingTag:        tag.IsBillingTag,
		ResourceAssignments: resourceAssignments,
		CreatedBy:           *tag.CreatedBy,
	}
}
