package bmcapimodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type SshKeyCreate struct {
	Default bool   `json:"default" yaml:"default"`
	Name    string `json:"name" yaml:"name"`
	Key     string `json:"key" yaml:"key"`
}

type SshKeyUpdate struct {
	Default bool   `json:"default" yaml:"default"`
	Name    string `json:"name" yaml:"name"`
}

func (sshKeyCreate SshKeyCreate) toSdk() *bmcapisdk.SshKeyCreate {
	return &bmcapisdk.SshKeyCreate{
		Default: sshKeyCreate.Default,
		Name:    sshKeyCreate.Name,
		Key:     sshKeyCreate.Key,
	}
}

func (sshKeyUpdate SshKeyUpdate) toSdk() *bmcapisdk.SshKeyUpdate {
	return &bmcapisdk.SshKeyUpdate{
		Default: sshKeyUpdate.Default,
		Name:    sshKeyUpdate.Name,
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

func CreateSshKeyUpdateRequestFromFile(filename string, commandname string) (*bmcapisdk.SshKeyUpdate, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var sshKeyUpdate SshKeyUpdate

	err = files.Unmarshal(data, &sshKeyUpdate, commandname)

	if err != nil {
		return nil, err
	}

	return sshKeyUpdate.toSdk(), nil
}
