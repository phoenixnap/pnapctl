package generators

import (
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func GenerateSshKey() bmcapisdk.SshKey {
	return bmcapisdk.SshKey{
		Id:            RandSeq(10),
		Default:       false,
		Name:          RandSeq(10),
		Key:           RandSeq(10),
		Fingerprint:   RandSeq(10),
		CreatedOn:     time.Now(),
		LastUpdatedOn: time.Now(),
	}
}

func GenerateSshKeyCreate() bmcapisdk.SshKeyCreate {
	return bmcapisdk.SshKeyCreate{
		Default: false,
		Name:    RandSeq(10),
		Key:     RandSeq(10),
	}

}

func GenerateSshKeyUpdate() bmcapisdk.SshKeyUpdate {
	return bmcapisdk.SshKeyUpdate{
		Default: false,
		Name:    RandSeq(10),
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
		Result:   RandSeq(10),
		SshKeyId: RandSeq(10),
	}
}
