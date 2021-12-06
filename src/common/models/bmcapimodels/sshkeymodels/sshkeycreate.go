package sshkeymodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
	"phoenixnap.com/pnapctl/testsupport/generators"
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

func GenerateSshKeyCreateSdk() bmcapisdk.SshKeyCreate {
	return bmcapisdk.SshKeyCreate{
		Default: false,
		Name:    generators.RandSeq(10),
		Key:     generators.RandSeq(10),
	}

}

func GenerateSshKeyCreateCli() SshKeyCreate {
	return SshKeyCreate{
		Default: false,
		Name:    generators.RandSeq(10),
		Key:     generators.RandSeq(10),
	}

}
