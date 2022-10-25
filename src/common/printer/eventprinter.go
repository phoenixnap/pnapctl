package printer

import (
	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi/v2"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func PrintEventListResponse(events []auditapisdk.Event) error {
	eventListToPrint := iterutils.Map(events, PrepareEventForPrinting)
	return MainPrinter.PrintOutput(eventListToPrint)
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
