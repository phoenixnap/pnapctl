package printer

import (
	locationapisdk "github.com/phoenixnap/go-sdk-bmc/locationapi/v2"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func PrintLocationListResponse(locations []locationapisdk.Location) error {
	locationListToPrint := iterutils.Map(locations, PrepareLocationForPrinting)
	return MainPrinter.PrintOutput(locationListToPrint)
}

func PrepareLocationForPrinting(location locationapisdk.Location) interface{} {
	table := OutputIsTable()

	switch {
	case table:
		return tables.ToLocationTable(location)
	default:
		return location
	}
}
