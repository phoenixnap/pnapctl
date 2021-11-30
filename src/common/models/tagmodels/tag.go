package tagmodels

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
)

type Tag struct {
	Id                  string                `json:"id" yaml:"id"`
	Name                string                `json:"name" yaml:"name"`
	Values              *[]string             `json:"values" yaml:"values"`
	Description         *string               `json:"description" yaml:"description"`
	IsBillingTag        bool                  `json:"isBillingTag" yaml:"isBillingTag"`
	ResourceAssignments *[]ResourceAssignment `json:"resourceAssignments" yaml:"resourceAssignments"`
}

func (tag *Tag) ToSdk() *tagapisdk.Tag {
	var assignments *[]tagapisdk.ResourceAssignment

	if tag.ResourceAssignments != nil {
		assignments = &[]tagapisdk.ResourceAssignment{}
		for _, assignment := range *tag.ResourceAssignments {
			*assignments = append(*assignments, *assignment.ToSdk())
		}
	}

	return &tagapisdk.Tag{
		Id:                  tag.Id,
		Name:                tag.Name,
		Values:              tag.Values,
		Description:         tag.Description,
		IsBillingTag:        tag.IsBillingTag,
		ResourceAssignments: assignments,
	}
}

func TagFromSdk(tag *tagapisdk.Tag) *Tag {
	var assignments *[]ResourceAssignment

	if tag == nil {
		return nil
	}

	if tag.ResourceAssignments != nil {
		assignments = &[]ResourceAssignment{}
		for _, assignment := range *tag.ResourceAssignments {
			*assignments = append(*assignments, *ResourceAssignmentFromSdk(&assignment))
		}
	}

	return &Tag{
		Id:                  tag.Id,
		Name:                tag.Name,
		Values:              tag.Values,
		Description:         tag.Description,
		IsBillingTag:        tag.IsBillingTag,
		ResourceAssignments: assignments,
	}
}
