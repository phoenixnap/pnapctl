package tables

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type ProductAvailabilityTable struct {
	ProductCode                 string   `header:"Product Code"`
	ProductCategory             string   `header:"Product Category"`
	LocationAvailabilityDetails []string `header:"Location Availability Details"`
}

func ProductAvailabilityTableFromSdk(sdk billingapi.ProductAvailability) ProductAvailabilityTable {
	return ProductAvailabilityTable{
		ProductCode:                 sdk.ProductCode,
		ProductCategory:             sdk.ProductCategory,
		LocationAvailabilityDetails: iterutils.MapRef(sdk.LocationAvailabilityDetails, models.LocationAvailabilityDetailsToTableString),
	}
}
