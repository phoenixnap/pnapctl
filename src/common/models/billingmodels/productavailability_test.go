package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestProductAvailabilityFromSdk_NotNil(t *testing.T) {
	sdk := *GenerateProductAvailability()
	cli := ProductAvailabilityFromSdk(sdk)

	assertEqualProductAvailability(t, sdk, cli)
}

func assertEqualProductAvailability(t *testing.T, sdk billingapi.ProductAvailability, cli ProductAvailability) {
	assert.Equal(t, sdk.ProductCategory, cli.ProductCategory)
	assert.Equal(t, sdk.ProductCode, cli.ProductCode)

	testutil.ForEachPair(sdk.LocationAvailabilityDetails, cli.LocationAvailabilityDetails).
		Do(t, assertLocationAvailabilityDetail)
}

func assertLocationAvailabilityDetail(t *testing.T, sdk billingapi.LocationAvailabilityDetail, cli LocationAvailabilityDetail) {
	assert.Equal(t, sdk.Location, cli.Location)
	assert.Equal(t, sdk.MinQuantityRequested, cli.MinQuantityRequested)
	assert.Equal(t, sdk.MinQuantityAvailable, cli.MinQuantityAvailable)
	assert.Equal(t, sdk.AvailableQuantity, cli.AvailableQuantity)

	testutil.ForEachPair(sdk.Solutions, cli.Solutions).
		Do(t, testutil.AssertEqual[string])
}
