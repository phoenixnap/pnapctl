package bmcapimodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	files "phoenixnap.com/pnap-cli/common/fileprocessor"
)

type SshKeyCreate struct {
	Default bool   `json:"default" yaml:"default"`
	Name    string `json:"name" yaml:"name"`
	Key     string `json:"key" yaml:"key"`
}

func (sshKeyCreate SshKeyCreate) toSdk() *bmcapisdk.SshKeyCreate {
	return &bmcapisdk.SshKeyCreate{
		Default: sshKeyCreate.Default,
		Name:    sshKeyCreate.Name,
		Key:     sshKeyCreate.Key,
	}
}

func CreateSshKeyCreateRequestFromFile(filename string, commandname string) (*bmcapisdk.SshKeyCreate, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var sshKeyCreate SshKeyCreate

	err = files.Unmarshal(data, &sshKeyCreate, commandname)

	if err != nil {
		return nil, err
	}

	return sshKeyCreate.toSdk(), nil
}
