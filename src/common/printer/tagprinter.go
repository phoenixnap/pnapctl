package printer

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/models/tagmodels"
)

func PrintTagResponse(tag tagapisdk.Tag, commandName string) error {
	tagToPrint := PrepareTagForPrinting(tag)
	return MainPrinter.PrintOutput(tagToPrint, commandName)
}

func PrintTagListResponse(tags []tagapisdk.Tag, commandName string) error {
	tagListToPrint := PrepareTagListForPrinting(tags)
	return MainPrinter.PrintOutput(tagListToPrint, commandName)
}

func PrepareTagForPrinting(tag tagapisdk.Tag) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.TagFromSdk(tag)
	default:
		return *tagmodels.TagFromSdk(&tag)
	}
}

func PrepareTagListForPrinting(tags []tagapisdk.Tag) []interface{} {
	var tagList []interface{}

	for _, tag := range tags {
		tagList = append(tagList, PrepareTagForPrinting(tag))
	}

	return tagList
}
