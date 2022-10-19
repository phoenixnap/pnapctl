package printer

import (
	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"phoenixnap.com/pnapctl/common/models/tables"
)

func PrintIpBlockResponse(ipBlock *ipapisdk.IpBlock, full bool) error {
	ipBlockToPrint := PrepareIpBlockForPrinting(*ipBlock, full)
	return MainPrinter.PrintOutput(ipBlockToPrint)
}

func PrintIpBlockListResponse(ipBlocks []ipapisdk.IpBlock, full bool) error {
	ipBlockListToPrint := PrepareIpBlockListForPrinting(ipBlocks, full)
	return MainPrinter.PrintOutput(ipBlockListToPrint)
}

func PrepareIpBlockListForPrinting(ipBlocks []ipapisdk.IpBlock, full bool) []interface{} {
	var ipBlockList []interface{}

	for _, ipBlock := range ipBlocks {
		ipBlockList = append(ipBlockList, PrepareIpBlockForPrinting(ipBlock, full))
	}

	return ipBlockList
}

func PrepareIpBlockForPrinting(ipBlock ipapisdk.IpBlock, full bool) interface{} {
	table := OutputIsTable()

	switch {
	case table && full:
		return tables.ToIpBlockTable(ipBlock)
	case table:
		return tables.ToShortIpBlockTable(ipBlock)
	default:
		return ipBlock
	}
}
