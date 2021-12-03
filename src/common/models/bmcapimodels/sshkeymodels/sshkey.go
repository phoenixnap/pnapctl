package sshkeymodels

import (
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

type SshKey struct {
	Id            string    `json:"id" yaml:"id"`
	Default       bool      `json:"default" yaml:"default"`
	Name          string    `json:"name" yaml:"name"`
	Key           string    `json:"key" yaml:"key"`
	Fingerprint   string    `json:"fingerprint" yaml:"fingerprint"`
	CreatedOn     time.Time `json:"createdOn" yaml:"createdOn"`
	LastUpdatedOn time.Time `json:"lastUpdatedOn" yaml:"lastUpdatedOn"`
}

func SshKeySdkToDto(sshKey bmcapisdk.SshKey) SshKey {
	return SshKey{
		Id:            sshKey.Id,
		Default:       sshKey.Default,
		Name:          sshKey.Name,
		Key:           sshKey.Key,
		Fingerprint:   sshKey.Fingerprint,
		CreatedOn:     sshKey.CreatedOn,
		LastUpdatedOn: sshKey.LastUpdatedOn,
	}
}
