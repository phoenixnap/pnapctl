package tables

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestProductAvailabilityTableFromSdk_NotNil(t *testing.T) {
	sdk := generators.Generate[billingapi.ProductAvailability]()
	tbl := ProductAvailabilityTableFromSdk(sdk)

	assertProductAvailabilityTablesEqual(t, sdk, tbl)
}

func assertProductAvailabilityTablesEqual(t *testing.T, sdk billingapi.ProductAvailability, tbl ProductAvailabilityTable) {
	assert.Equal(t, sdk.ProductCode, tbl.ProductCode)
	assert.Equal(t, sdk.ProductCategory, tbl.ProductCategory)

	sdkAsTableString := iterutils.MapRef(sdk.LocationAvailabilityDetails, models.LocationAvailabilityDetailsToTableString)
	assert.Equal(t, sdkAsTableString, tbl.LocationAvailabilityDetails)
}
