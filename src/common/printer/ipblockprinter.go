package printer

import (
	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi"
	"phoenixnap.com/pnapctl/common/models/ipmodels"
	"phoenixnap.com/pnapctl/common/models/tables"
)

func PrintIpBlockResponse(ipBlock ipapisdk.IpBlock, commandName string) error {
	ipBlockToPrint := PrepareIpBlockForPrinting(ipBlock)
	return MainPrinter.PrintOutput(ipBlockToPrint, commandName)
}

func PrintIpBlockListResponse(ipBlocks []ipapisdk.IpBlock, commandName string) error {
	ipBlockListToPrint := PrepareIpBlockListForPrinting(ipBlocks)
	return MainPrinter.PrintOutput(ipBlockListToPrint, commandName)
}

func PrepareIpBlockListForPrinting(ipBlocks []ipapisdk.IpBlock) []interface{} {
	var ipBlockList []interface{}

	for _, ipBlock := range ipBlocks {
		ipBlockList = append(ipBlockList, PrepareIpBlockForPrinting(ipBlock))
	}

	return ipBlockList
}

func PrepareIpBlockForPrinting(ipBlock ipapisdk.IpBlock) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ToIpBlockTable(ipBlock)
	default:
		return ipmodels.IpBlockFromSdk(ipBlock)
	}
}
