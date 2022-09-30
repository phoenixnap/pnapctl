package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type ServerPatch struct {
	Hostname    *string `yaml:"hostname,omitempty" json:"hostname,omitempty"`
	Description *string `yaml:"description,omitempty" json:"description,omitempty"`
}

func PatchServerRequestFromFile(filename string, commandname string) (*bmcapisdk.ServerPatch, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file using the struct
	var serverPatch ServerPatch

	err = files.Unmarshal(data, &serverPatch, commandname)

	if err != nil {
		return nil, err
	}

	return serverPatch.ToSdk(), nil
}

func (serverPatch ServerPatch) ToSdk() *bmcapisdk.ServerPatch {
	return &bmcapisdk.ServerPatch{
		Hostname:    serverPatch.Hostname,
		Description: serverPatch.Description,
	}
}
