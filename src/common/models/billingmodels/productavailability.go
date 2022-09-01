package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type ProductAvailability struct {
	ProductCode                 string                       `json:"productCode" yaml:"productCode"`
	ProductCategory             string                       `json:"productCategory" yaml:"productCategory"`
	LocationAvailabilityDetails []LocationAvailabilityDetail `json:"locationAvailabilityDetails" yaml:"locationAvailabilityDetails"`
}

type LocationAvailabilityDetail struct {
	Location             billingapi.LocationEnum `json:"location" yaml:"location"`
	MinQuantityRequested float32                 `json:"minQuantityRequested" yaml:"minQuantityRequested"`
	MinQuantityAvailable bool                    `json:"minQuantityAvailable" yaml:"minQuantityAvailable"`
	AvailableQuantity    float32                 `json:"availableQuantity" yaml:"availableQuantity"`
	Solutions            []string                `json:"solutions" yaml:"solutions"`
}

func ProductAvailabilityFromSdk(sdk *billingapi.ProductAvailability) *ProductAvailability {
	if sdk == nil {
		return nil
	}

	return &ProductAvailability{
		ProductCode:     sdk.ProductCode,
		ProductCategory: sdk.ProductCategory,
		LocationAvailabilityDetails: iterutils.Map(
			sdk.LocationAvailabilityDetails,
			locationAvailabilityDetailsFromSdk,
		),
	}
}

func locationAvailabilityDetailsFromSdk(sdk billingapi.LocationAvailabilityDetail) LocationAvailabilityDetail {
	return LocationAvailabilityDetail{
		Location:             sdk.Location,
		MinQuantityRequested: sdk.MinQuantityRequested,
		MinQuantityAvailable: sdk.MinQuantityAvailable,
		AvailableQuantity:    sdk.AvailableQuantity,
		Solutions:            sdk.Solutions,
	}
}
