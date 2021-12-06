package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type TagAssignmentRequest struct {
	Name  string  `yaml:"name" json:"name"`
	Value *string `yaml:"value,omitempty" json:"value,omitempty"`
}

func (tagAssignmentRequest TagAssignmentRequest) toSdk() bmcapisdk.TagAssignmentRequest {
	var tagAssignmentRequestSdk = bmcapisdk.TagAssignmentRequest{
		Name:  tagAssignmentRequest.Name,
		Value: tagAssignmentRequest.Value,
	}

	return tagAssignmentRequestSdk
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
