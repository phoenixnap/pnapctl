package printer

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func PrintTagResponse(tag *tagapisdk.Tag) error {
	tagToPrint := PrepareTagForPrinting(*tag)
	return MainPrinter.PrintOutput(tagToPrint)
}

func PrintTagListResponse(tags []tagapisdk.Tag) error {
	tagListToPrint := iterutils.Map(tags, PrepareTagForPrinting)
	return MainPrinter.PrintOutput(tagListToPrint)
}

func PrepareTagForPrinting(tag tagapisdk.Tag) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.TagFromSdk(tag)
	default:
		return tag
	}
}
