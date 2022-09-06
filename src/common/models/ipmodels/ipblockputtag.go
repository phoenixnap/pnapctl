package ipmodels

import (
	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type IpBlockPutTag struct {
	Name string `yaml:"name" json:"name"`
}

func PutIpBlockTagRequestFromFile(filename string, commandname string) (*[]ipapisdk.TagAssignmentRequest, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var IpBlockPutTag IpBlockPutTag

	err = files.Unmarshal(data, &IpBlockPutTag, commandname)

	if err != nil {
		return nil, err
	}

	return IpBlockPutTag.ToSdk(), nil
}

func (ipBlockPutTag IpBlockPutTag) ToSdk() *[]ipapisdk.TagAssignmentRequest {
	return &[]ipapisdk.TagAssignmentRequest{
		*ipapisdk.NewTagAssignmentRequest(ipBlockPutTag.Name),
	}
}
