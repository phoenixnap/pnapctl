package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type RelinquishIpBlock struct {
	DeleteIpBlocks *bool `json:"deleteIpBlocks,omitempty" yaml:"deleteIpBlocks,omitempty"`
}

func (relinquishIpBlock RelinquishIpBlock) ToSdk() *bmcapisdk.RelinquishIpBlock {
	return &bmcapisdk.RelinquishIpBlock{
		DeleteIpBlocks: relinquishIpBlock.DeleteIpBlocks,
	}
}

func CreateRelinquishIpBlockRequestFromFile(filename string, commandname string) (*bmcapisdk.RelinquishIpBlock, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var relinquishIpBlock RelinquishIpBlock

	err = files.Unmarshal(data, &relinquishIpBlock, commandname)

	if err != nil {
		return nil, err
	}

	return relinquishIpBlock.ToSdk(), nil
}
