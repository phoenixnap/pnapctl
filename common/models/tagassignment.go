package models

import (
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
)

type TagAssignmentRequest struct {
	Name  string  `yaml:"name" json:"name"`
	Value *string `yaml:"value,omitempty" json:"value,omitempty"`
}

func tagAssignmentRequestDtoToSdk(tagAssignmentRequest *[]TagAssignmentRequest) *[]bmcapi.TagAssignmentRequest {
	if tagAssignmentRequest == nil {
		return nil
	}

	var list []bmcapi.TagAssignmentRequest

	for _, x := range *tagAssignmentRequest {
		converted := &bmcapi.TagAssignmentRequest{
			Name:  x.Name,
			Value: x.Value,
		}

		list = append(list, *converted)
	}

	return &list
}
