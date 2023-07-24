package tables

import (
	"strings"

	"github.com/phoenixnap/go-sdk-bmc/locationapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type Location struct {
	Location            string `header:"Location"`
	LocationDescription string `header:"Description"`
	ProductCategories   string `header:"Product Categories`
}

func ToLocationTable(location locationapi.Location) Location {
	return Location{
		Location:            string(location.Location),
		LocationDescription: *location.LocationDescription,
		ProductCategories:   strings.Join(iterutils.MapRef(location.ProductCategories, models.ProductCategoryToTableString), "\n\n"),
	}
}
