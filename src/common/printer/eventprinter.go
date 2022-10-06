package printer

import (
	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	"phoenixnap.com/pnapctl/common/models/tables"
)

func PrintEventListResponse(events []auditapisdk.Event, commandName string) error {
	eventListToPrint := PrepareEventListForPrinting(events)
	return MainPrinter.PrintOutput(eventListToPrint, commandName)
}

func PrepareEventForPrinting(event auditapisdk.Event) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ToEventTable(event)
	default:
		return event
	}
}

func PrepareEventListForPrinting(events []auditapisdk.Event) []interface{} {
	var eventList []interface{}

	for _, event := range events {
		eventList = append(eventList, PrepareEventForPrinting(event))
	}

	return eventList
}
