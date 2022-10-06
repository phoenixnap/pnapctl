package printer

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/sshkeymodels"
	"phoenixnap.com/pnapctl/common/models/tables"
)

func PrintSshKeyResponse(sshKey *bmcapisdk.SshKey, full bool, commandName string) error {
	sshKeyToPrint := PrepareSshKeyForPrinting(*sshKey, full)
	return MainPrinter.PrintOutput(sshKeyToPrint, commandName)
}

func PrintSshKeyListResponse(sshKeys []bmcapisdk.SshKey, full bool, commandName string) error {
	sshKeyListToPrint := PrepareSshKeyListForPrinting(sshKeys, full)
	return MainPrinter.PrintOutput(sshKeyListToPrint, commandName)
}

func PrepareSshKeyListForPrinting(quotas []bmcapisdk.SshKey, full bool) []interface{} {
	var sshKeyList []interface{}

	for _, sdkSshKey := range quotas {
		sshKeyList = append(sshKeyList, PrepareSshKeyForPrinting(sdkSshKey, full))
	}

	return sshKeyList
}

func PrepareSshKeyForPrinting(sshKey bmcapisdk.SshKey, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case full && table:
		return tables.ToSshKeyTableFull(sshKey)
	case !full && table:
		return tables.ToSshKeyTable(sshKey)
	default:
		return sshkeymodels.SshKeyFromSdk(sshKey)
	}
}
