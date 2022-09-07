package ipmodels

import (
	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type TagAssignmentRequest struct {
	Name  string  `yaml:"name" json:"name"`
	Value *string `yaml:"value,omitempty" json:"value,omitempty"`
}

func PutIpBlockTagRequestFromFile(filename string, commandname string) ([]ipapisdk.TagAssignmentRequest, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var TagAssignmentRequests []TagAssignmentRequest
	err = files.Unmarshal(data, &TagAssignmentRequests, commandname)

	if err != nil {
		return nil, err
	}

	return mapTagAssignmentRequestToSdk(&TagAssignmentRequests), nil
}

func (tagAssignmentRequest TagAssignmentRequest) ToSdk() ipapisdk.TagAssignmentRequest {
	var tagAssignmentRequestSdk = ipapisdk.TagAssignmentRequest{
		Name:  tagAssignmentRequest.Name,
		Value: tagAssignmentRequest.Value,
	}

	return tagAssignmentRequestSdk
}

func mapTagAssignmentRequestToSdk(TagAssignmentRequest *[]TagAssignmentRequest) []ipapisdk.TagAssignmentRequest {

	if TagAssignmentRequest == nil {
		return nil
	}

	var tagAssignmentRequests []ipapisdk.TagAssignmentRequest

	for _, TagAssignmentRequest := range *TagAssignmentRequest {
		tagAssignmentRequests = append(tagAssignmentRequests, TagAssignmentRequest.ToSdk())
	}

	return tagAssignmentRequests
}
