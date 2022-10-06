package ipmodels

import (
	"fmt"

	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
)

type TagAssignment struct {
	// The unique id of the tag.
	Id string `yaml:"id" json:"id"`
	// The name of the tag.
	Name string `yaml:"name" json:"name"`
	// The value of the tag assigned to the resource.
	Value *string `yaml:"value,omitempty" json:"value,omitempty"`
	// Whether or not to show the tag as part of billing and invoices
	IsBillingTag bool `yaml:"isBillingTag" json:"isBillingTag"`
	// Who the tag was created by.
	CreatedBy *string `yaml:"createdBy,omitempty" json:"createdBy,omitempty"`
}

func TagAssignmentFromSdk(tagAssignment ipapisdk.TagAssignment) TagAssignment {
	return TagAssignment{
		Id:           tagAssignment.Id,
		Name:         tagAssignment.Name,
		Value:        tagAssignment.Value,
		IsBillingTag: tagAssignment.IsBillingTag,
		CreatedBy:    tagAssignment.CreatedBy,
	}
}

func TagAssignmentToTableString(tagAssignment *ipapisdk.TagAssignment) string {

	if tagAssignment == nil {
		return ""
	}

	Value, CreatedBy := "", ""

	if tagAssignment.Value != nil {
		Value = *tagAssignment.Value
	}

	if tagAssignment.CreatedBy != nil {
		CreatedBy = *tagAssignment.CreatedBy
	}

	return fmt.Sprintf("ID: %s\nName: %s\nValue: %s\nIsBillingTag: %t\nCreated By: %s", tagAssignment.Id, tagAssignment.Name, Value, tagAssignment.IsBillingTag, CreatedBy)
}
