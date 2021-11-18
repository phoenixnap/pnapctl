package printer

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	sshKeyModel "phoenixnap.com/pnap-cli/common/models/bmcapimodels"
	"phoenixnap.com/pnap-cli/common/models/tables"
)

func PrintSshKeyResponse(sshKey bmcapisdk.SshKey, commandName string) error {
	sshKeyToPrint := PrepareSshKeyForPrinting(sshKey)
	return MainPrinter.PrintOutput(sshKeyToPrint, commandName)
}

func PrintSshKeyListResponse(sshKeys []bmcapisdk.SshKey, commandName string) error {
	sshKeyListToPrint := PrepareSshKeyListForPrinting(sshKeys)
	return MainPrinter.PrintOutput(sshKeyListToPrint, commandName)
}

func PrepareSshKeyListForPrinting(quotas []bmcapisdk.SshKey) []interface{} {
	var sshKeyList []interface{}

	for _, sdkSshKey := range quotas {
		sshKeyList = append(sshKeyList, PrepareSshKeyForPrinting(sdkSshKey))
	}

	return sshKeyList
}

func PrepareSshKeyForPrinting(sshKey bmcapisdk.SshKey) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ToSshKeyTable(sshKey)
	default:
		return sshKeyModel.SshKeySdkToDto(sshKey)
	}
}
