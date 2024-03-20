package tables

import (
	"github.com/phoenixnap/go-sdk-bmc/locationapi/v3"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type Location struct {
	Location            string   `header:"Location"`
	LocationDescription string   `header:"Description"`
	ProductCategories   []string `header:"Product Categories`
}

func ToLocationTable(location locationapi.Location) Location {
	return Location{
		Location:            string(location.Location),
		LocationDescription: *location.LocationDescription,
		ProductCategories:   iterutils.MapRef(location.ProductCategories, models.ProductCategoryToTableString),
	}
}
