package ipmodels

import (
	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type IpBlockCreate struct {
	Location      string `yaml:"location" json:"location"`
	CidrBlockSize string `yaml:"cidrBlockSize" json:"cidrBlockSize"`
}

func CreateIpBlockRequestFromFile(filename string, commandname string) (*ipapisdk.IpBlockCreate, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var ipBlockCreate IpBlockCreate

	err = files.Unmarshal(data, &ipBlockCreate, commandname)

	if err != nil {
		return nil, err
	}

	return ipBlockCreate.ToSdk(), nil
}

func (ipBlockCreate IpBlockCreate) ToSdk() *ipapisdk.IpBlockCreate {
	return &ipapisdk.IpBlockCreate{
		Location:      ipBlockCreate.Location,
		CidrBlockSize: ipBlockCreate.CidrBlockSize,
	}
}
