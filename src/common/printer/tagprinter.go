package printer

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
	"phoenixnap.com/pnapctl/common/models/tables"
)

func PrintTagResponse(tag *tagapisdk.Tag) error {
	tagToPrint := PrepareTagForPrinting(*tag)
	return MainPrinter.PrintOutput(tagToPrint)
}

func PrintTagListResponse(tags []tagapisdk.Tag) error {
	tagListToPrint := PrepareTagListForPrinting(tags)
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

func PrepareTagListForPrinting(tags []tagapisdk.Tag) []interface{} {
	var tagList []interface{}

	for _, tag := range tags {
		tagList = append(tagList, PrepareTagForPrinting(tag))
	}

	return tagList
}
