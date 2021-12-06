package sshkeymodels

import (
	"time"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/testsupport/generators"
)

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
