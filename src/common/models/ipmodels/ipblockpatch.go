package ipmodels

import (
	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type IpBlockPatch struct {
	Description *string `yaml:"description,omitempty" json:"description,omitempty"`
}

func PatchIpBlockRequestFromFile(filename string, commandname string) (*ipapisdk.IpBlockPatch, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var ipBlockPatch IpBlockPatch

	err = files.Unmarshal(data, &ipBlockPatch, commandname)

	if err != nil {
		return nil, err
	}

	return ipBlockPatch.ToSdk(), nil
}

func (ipBlockPatch IpBlockPatch) ToSdk() *ipapisdk.IpBlockPatch {
	return &ipapisdk.IpBlockPatch{
		Description: ipBlockPatch.Description,
	}
}
