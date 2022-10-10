package servermodels

import (
	"fmt"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

type TagAssignment struct {
	Id           string  `yaml:"id" json:"id"`
	Name         string  `yaml:"name" json:"name"`
	Value        *string `yaml:"value,omitempty" json:"value,omitempty"`
	IsBillingTag bool    `yaml:"isBillingTag" json:"isBillingTag"`
	CreatedBy    *string `yaml:"createdBy,omitempty" json:"createdBy,omitempty"`
}

func TagAssignmentListFromSdk(tagAssignment []bmcapisdk.TagAssignment) []TagAssignment {
	if tagAssignment == nil {
		return nil
	}

	var tagAssignments []TagAssignment

	for _, bmcTagAssignment := range tagAssignment {
		mappedTagAssignment := &TagAssignment{
			Id:           bmcTagAssignment.Id,
			Name:         bmcTagAssignment.Name,
			Value:        bmcTagAssignment.Value,
			IsBillingTag: bmcTagAssignment.IsBillingTag,
			CreatedBy:    bmcTagAssignment.CreatedBy,
		}

		tagAssignments = append(tagAssignments, *mappedTagAssignment)
	}

	return tagAssignments
}

func (t TagAssignment) toTableString() string {
	var tagValue string

	if t.Value == nil {
		tagValue = ""
	} else {
		tagValue = ": " + *t.Value
	}
	return fmt.Sprintf("(%s) %s%s", t.Id, t.Name, tagValue)
}

func TagsToTableStrings(tags []bmcapisdk.TagAssignment) []string {
	var tagStrings []string
	if tags == nil {
		tagStrings = []string{}
	} else {
		tagDetails := TagAssignmentListFromSdk(tags)
		for _, tag := range tagDetails {
			tagStrings = append(tagStrings, tag.toTableString())
		}
	}

	return tagStrings
}
