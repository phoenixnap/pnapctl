package tables

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

type SshKeyTableFull struct {
	Id            string  `header:"id"`
	Default       bool    `header:"default"`
	Name          string  `header:"name"`
	Fingerprint   string  `header:"fingerprint"`
	CreatedOn     string  `header:"Created On"`
	LastUpdatedOn string  `header:"Last Updated On"`
	Key           *string `header:"key"`
}

type SshKeyTable struct {
	Id            string `header:"id"`
	Default       bool   `header:"default"`
	Name          string `header:"name"`
	Fingerprint   string `header:"fingerprint"`
	CreatedOn     string `header:"Created On"`
	LastUpdatedOn string `header:"Last Updated On"`
}

func ToSshKeyTableFull(sshKey bmcapisdk.SshKey) SshKeyTableFull {
	return SshKeyTableFull{
		Id:            sshKey.Id,
		Default:       sshKey.Default,
		Name:          sshKey.Name,
		Key:           &sshKey.Key,
		Fingerprint:   sshKey.Fingerprint,
		CreatedOn:     sshKey.CreatedOn.String(),
		LastUpdatedOn: sshKey.LastUpdatedOn.String(),
	}
}

func ToSshKeyTable(sshKey bmcapisdk.SshKey) SshKeyTable {
	return SshKeyTable{
		Id:            sshKey.Id,
		Default:       sshKey.Default,
		Name:          sshKey.Name,
		Fingerprint:   sshKey.Fingerprint,
		CreatedOn:     sshKey.CreatedOn.String(),
		LastUpdatedOn: sshKey.LastUpdatedOn.String(),
	}
}
