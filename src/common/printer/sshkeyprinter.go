package printer

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func PrintSshKeyResponse(sshKey *bmcapisdk.SshKey, full bool) error {
	sshKeyToPrint := PrepareSshKeyForPrinting(*sshKey, full)
	return MainPrinter.PrintOutput(sshKeyToPrint)
}

func PrintSshKeyListResponse(sshKeys []bmcapisdk.SshKey, full bool) error {
	sshKeyListToPrint := iterutils.Map(sshKeys, withFull(full, PrepareSshKeyForPrinting))
	return MainPrinter.PrintOutput(sshKeyListToPrint)
}

func PrepareSshKeyForPrinting(sshKey bmcapisdk.SshKey, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case full && table:
		return tables.ToSshKeyTableFull(sshKey)
	case !full && table:
		return tables.ToSshKeyTable(sshKey)
	default:
		return sshKey
	}
}
