package bmcapimodels

import (
	"fmt"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	files "phoenixnap.com/pnap-cli/common/fileprocessor"
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

func mapTagAssignmentRequestToSdk(tagAssignmentRequest *[]TagAssignmentRequest) *[]bmcapisdk.TagAssignmentRequest {
	if tagAssignmentRequest == nil {
		return nil
	}

	var tagAssignmentRequests []bmcapisdk.TagAssignmentRequest

	for _, tagAssignmentRequest := range *tagAssignmentRequest {
		tagAssignmentRequests = append(tagAssignmentRequests, tagAssignmentRequest.toSdk())
	}

	return &tagAssignmentRequests
}

func (tagAssignmentRequest TagAssignmentRequest) toSdk() bmcapisdk.TagAssignmentRequest {
	var tagAssignmentRequestSdk = bmcapisdk.TagAssignmentRequest{
		Name:  tagAssignmentRequest.Name,
		Value: tagAssignmentRequest.Value,
	}

	return tagAssignmentRequestSdk
}

func TagAssignmentSdkToDto(tagAssignment *[]bmcapisdk.TagAssignment) *[]TagAssignment {
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

func (t TagAssignment) ToTableString() string {
	var tagValue string

	if t.Value == nil {
		tagValue = ""
	} else {
		tagValue = ": " + *t.Value
	}
	return fmt.Sprintf("(%s) %s%s", t.Id, t.Name, tagValue)
}

func TagsToTableStrings(tags *[]bmcapisdk.TagAssignment) []string {
	var tagStrings []string
	if tags == nil {
		tagStrings = []string{}
	} else {
		dtoTags := TagAssignmentSdkToDto(tags)
		for _, tag := range *dtoTags {
			tagStrings = append(tagStrings, tag.ToTableString())
		}
	}

	return tagStrings
}

func TagServerRequestFromFile(filename string, commandname string) (*[]bmcapisdk.TagAssignmentRequest, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file using the struct
	var tagAssignmentRequests []TagAssignmentRequest

	err = files.Unmarshal(data, &tagAssignmentRequests, commandname)

	if err != nil {
		return nil, err
	}

	return mapTagAssignmentRequestToSdk(&tagAssignmentRequests), nil
}
