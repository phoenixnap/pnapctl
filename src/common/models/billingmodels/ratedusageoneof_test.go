package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/billingmodels/ratedusageoneof"
)

func TestRatedUsageActualFromSdk_Bandwidth(test_framework *testing.T) {
	bandwidthRecord := GenerateBandwidthRecordSdk()
	ratedUsageResponse := billingapi.RatedUsageGet200ResponseInner{
		BandwidthRecord: bandwidthRecord,
	}

	actual := RatedUsageActualFromSdk(ratedUsageResponse)
	assertEqualAsBandwidth(test_framework, actual, *bandwidthRecord)
}

func assertEqualAsBandwidth(
	test_framework *testing.T,
	cliRecord interface{},
	sdkBandwidth billingapi.BandwidthRecord,
) {
	cliBandwidth := cliRecord.(ratedusageoneof.BandwidthRecord)

	assert.Equal(test_framework, cliBandwidth.Id, sdkBandwidth.Id)
	assert.Equal(test_framework, cliBandwidth.ProductCategory, sdkBandwidth.ProductCategory)
	assert.Equal(test_framework, cliBandwidth.ProductCode, sdkBandwidth.ProductCode)
	assert.Equal(test_framework, cliBandwidth.Location, sdkBandwidth.Location)
	assert.Equal(test_framework, cliBandwidth.YearMonth, sdkBandwidth.YearMonth)
	assert.Equal(test_framework, cliBandwidth.StartDateTime, sdkBandwidth.StartDateTime)
	assert.Equal(test_framework, cliBandwidth.EndDateTime, sdkBandwidth.EndDateTime)
	assert.Equal(test_framework, cliBandwidth.Cost, sdkBandwidth.Cost)
	assert.Equal(test_framework, cliBandwidth.PriceModel, sdkBandwidth.PriceModel)
	assert.Equal(test_framework, cliBandwidth.UnitPrice, sdkBandwidth.UnitPrice)
	assert.Equal(test_framework, cliBandwidth.UnitPriceDescription, sdkBandwidth.UnitPriceDescription)
	assert.Equal(test_framework, cliBandwidth.Quantity, sdkBandwidth.Quantity)
	assert.Equal(test_framework, cliBandwidth.Active, sdkBandwidth.Active)
	assert.Equal(test_framework, cliBandwidth.UsageSessionId, sdkBandwidth.UsageSessionId)
	assert.Equal(test_framework, cliBandwidth.CorrelationId, sdkBandwidth.CorrelationId)
	assert.Equal(test_framework, cliBandwidth.ReservationId, sdkBandwidth.ReservationId)

	assert.Equal(test_framework, cliBandwidth.Metadata.IngressGb, sdkBandwidth.Metadata.IngressGb)
	assert.Equal(test_framework, cliBandwidth.Metadata.EgressGb, sdkBandwidth.Metadata.EgressGb)
	assert.Equal(test_framework, cliBandwidth.Metadata.PackageQuantity, sdkBandwidth.Metadata.PackageQuantity)
	assert.Equal(test_framework, cliBandwidth.Metadata.PackageUnit, sdkBandwidth.Metadata.PackageUnit)
}
