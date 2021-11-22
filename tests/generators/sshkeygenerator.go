package generators

import (
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func GenerateSshKey() bmcapisdk.SshKey {
	return bmcapisdk.SshKey{
		Id:            randSeq(10),
		Default:       false,
		Name:          randSeq(10),
		Key:           randSeq(10),
		Fingerprint:   randSeq(10),
		CreatedOn:     time.Now(),
		LastUpdatedOn: time.Now(),
	}
}

func GenerateSshKeyCreate() bmcapisdk.SshKeyCreate {
	return bmcapisdk.SshKeyCreate{
		Default: false,
		Name:    randSeq(10),
		Key:     randSeq(10),
	}

}

func GenerateSshKeys(n int) []bmcapisdk.SshKey {
	var sshKeyList []bmcapisdk.SshKey
	for i := 0; i < n; i++ {
		sshKeyList = append(sshKeyList, GenerateSshKey())
	}
	return sshKeyList
}

func GenerateSshKeyDeleteResult() bmcapisdk.DeleteSshKeyResult {
	return bmcapisdk.DeleteSshKeyResult{
		Result:   randSeq(10),
		SshKeyId: randSeq(10),
	}
}
