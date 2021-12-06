package sshkeymodels

import (
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/testsupport/generators"
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

func GenerateSshKeySdk() bmcapisdk.SshKey {
	return bmcapisdk.SshKey{
		Id:            generators.RandSeq(10),
		Default:       false,
		Name:          generators.RandSeq(10),
		Key:           generators.RandSeq(10),
		Fingerprint:   generators.RandSeq(10),
		CreatedOn:     time.Now(),
		LastUpdatedOn: time.Now(),
	}
}

func GenerateSshKeyListSdk(n int) []bmcapisdk.SshKey {
	var sshKeyList []bmcapisdk.SshKey
	for i := 0; i < n; i++ {
		sshKeyList = append(sshKeyList, GenerateSshKeySdk())
	}
	return sshKeyList
}

func GenerateSshKeyDeleteResultSdk() bmcapisdk.DeleteSshKeyResult {
	return bmcapisdk.DeleteSshKeyResult{
		Result:   generators.RandSeq(10),
		SshKeyId: generators.RandSeq(10),
	}
}
