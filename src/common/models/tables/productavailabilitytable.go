package tables

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

type ProductAvailabilityTable struct {
	ProductCode                 string
	ProductCategory             string
	LocationAvailabilityDetails []string
}

func ProductAvailabilityTableFromSdk(sdk *billingapi.ProductAvailability) *ProductAvailabilityTable {
	if sdk == nil {
		return nil
	}

	return &ProductAvailabilityTable{}
}
