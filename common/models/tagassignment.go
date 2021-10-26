package models

import (
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
)

type TagAssignmentRequest struct {
	Name  string  `yaml:"name" json:"name"`
	Value *string `yaml:"value,omitempty" json:"value,omitempty"`
}

type TagAssignment struct {
	Id           string  `yaml:"id" json:"id"`
	Name         string  `yaml:"name" json:"name"`
	Value        *string `yaml:"value,omitempty" json:"value,omitempty"`
	IsBillingTag bool    `yaml:"isBillingTag" json:"isBillingTag"`
}

func mapTagAssignmentRequestToSdk(tagAssignmentRequest *[]TagAssignmentRequest) *[]bmcapi.TagAssignmentRequest {
	if tagAssignmentRequest == nil {
		return nil
	}

	var tagAssignmentRequests []bmcapi.TagAssignmentRequest

	for _, tagAssignmentRequest := range *tagAssignmentRequest {
		tagAssignmentRequests = append(tagAssignmentRequests, tagAssignmentRequest.toSdk())
	}

	return &tagAssignmentRequests
}

func (tagAssignmentRequest TagAssignmentRequest) toSdk() bmcapi.TagAssignmentRequest {
	var tagAssignmentRequestSdk = bmcapi.TagAssignmentRequest{
		Name:  tagAssignmentRequest.Name,
		Value: tagAssignmentRequest.Value,
	}

	return tagAssignmentRequestSdk
}

func tagAssignmentSdkToDto(tagAssignment *[]bmcapi.TagAssignment) *[]TagAssignment {
	if tagAssignment == nil {
		return nil
	}

	var tagAssignments []TagAssignment

	for _, bmcTagAssignment := range *tagAssignment {
		mappedTagAssignment := &TagAssignment{
			Id:           bmcTagAssignment.Id,
			Name:         bmcTagAssignment.Name,
			Value:        bmcTagAssignment.Value,
			IsBillingTag: bmcTagAssignment.IsBillingTag,
		}

		tagAssignments = append(tagAssignments, *mappedTagAssignment)
	}

	return &tagAssignments
}
