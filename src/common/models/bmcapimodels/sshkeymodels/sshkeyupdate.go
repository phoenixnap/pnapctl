package sshkeymodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
	"phoenixnap.com/pnapctl/testsupport/generators"
)

type SshKeyUpdate struct {
	Default bool   `json:"default" yaml:"default"`
	Name    string `json:"name" yaml:"name"`
}

func (sshKeyUpdate SshKeyUpdate) toSdk() *bmcapisdk.SshKeyUpdate {
	return &bmcapisdk.SshKeyUpdate{
		Default: sshKeyUpdate.Default,
		Name:    sshKeyUpdate.Name,
	}
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

func GenerateSshKeyUpdateSdk() bmcapisdk.SshKeyUpdate {
	return bmcapisdk.SshKeyUpdate{
		Default: false,
		Name:    generators.RandSeq(10),
	}
}

func GenerateSshKeyUpdateCli() SshKeyUpdate {
	return SshKeyUpdate{
		Default: false,
		Name:    generators.RandSeq(10),
	}
}
