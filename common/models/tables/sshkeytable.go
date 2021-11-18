package tables

import (
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

type SshKeyTable struct {
	Id            string    `header:"id"`
	Default       bool      `header:"default"`
	Name          string    `header:"name"`
	Key           string    `header:"key"`
	Fingerprint   string    `header:"fingerprint"`
	CreatedOn     time.Time `header:"Created On"`
	LastUpdatedOn time.Time `header:"Last Updated On"`
}

func ToSshKeyTable(sshKey bmcapisdk.SshKey) SshKeyTable {
	return SshKeyTable{
		Id:            sshKey.Id,
		Default:       sshKey.Default,
		Name:          sshKey.Name,
		Key:           sshKey.Key,
		Fingerprint:   sshKey.Fingerprint,
		CreatedOn:     sshKey.CreatedOn,
		LastUpdatedOn: sshKey.LastUpdatedOn,
	}
}
