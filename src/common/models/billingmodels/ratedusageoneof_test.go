package billingmodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/billingmodels/ratedusageoneof"
)

// FromSdk tests
func TestRatedUsageActualFromSdk_BandwidthRecord(test_framework *testing.T) {
	bandwidthRecord := GenerateBandwidthRecordSdk()
	ratedUsageResponse := billingapi.RatedUsageGet200ResponseInner{
		BandwidthRecord: bandwidthRecord,
	}

	actual := RatedUsageActualFromSdk(ratedUsageResponse)
	assertEqualAsBandwidthRecord(test_framework, actual, *bandwidthRecord)
}

func TestRatedUsageActualFromSdk_OperatingSystemRecord(test_framework *testing.T) {
	operatingSystemRecord := GenerateOperatingSystemRecordSdk()
	ratedUsageResponse := billingapi.RatedUsageGet200ResponseInner{
		OperatingSystemRecord: operatingSystemRecord,
	}

	actual := RatedUsageActualFromSdk(ratedUsageResponse)
	assertEqualAsOperatingSystemRecord(test_framework, actual, *operatingSystemRecord)
}

func TestRatedUsageActualFromSdk_PublicSubnetRecord(test_framework *testing.T) {
	publicSubnetRecord := GeneratePublicSubnetRecordSdk()
	ratedUsageResponse := billingapi.RatedUsageGet200ResponseInner{
		PublicSubnetRecord: publicSubnetRecord,
	}

	actual := RatedUsageActualFromSdk(ratedUsageResponse)
	assertEqualAsPublicSubnetRecord(test_framework, actual, *publicSubnetRecord)
}

func TestRatedUsageActualFromSdk_ServerRecord(test_framework *testing.T) {
	serverRecord := GenerateServerRecordSdk()
	ratedUsageResponse := billingapi.RatedUsageGet200ResponseInner{
		ServerRecord: serverRecord,
	}

	actual := RatedUsageActualFromSdk(ratedUsageResponse)
	assertEqualAsServerRecord(test_framework, actual, *serverRecord)
}

// Equality asserts

func assertEqualAsBandwidthRecord(
	test_framework *testing.T,
	cliRecord interface{},
	sdkBandwidth billingapi.BandwidthRecord,
) {
	cliBandwidth := cliRecord.(*ratedusageoneof.BandwidthRecord)

	assert.Equal(test_framework, cliBandwidth.Id, sdkBandwidth.Id)
	assert.Equal(test_framework, cliBandwidth.ProductCategory, sdkBandwidth.ProductCategory)
	assert.Equal(test_framework, cliBandwidth.ProductCode, sdkBandwidth.ProductCode)
	assert.Equal(test_framework, string(cliBandwidth.Location), string(sdkBandwidth.Location))
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

func assertEqualAsOperatingSystemRecord(
	test_framework *testing.T,
	cliRecord interface{},
	sdkOperatingSystem billingapi.OperatingSystemRecord,
) {
	cliOperatingSystem := cliRecord.(*ratedusageoneof.OperatingSystemRecord)

	assert.Equal(test_framework, cliOperatingSystem.Id, sdkOperatingSystem.Id)
	assert.Equal(test_framework, cliOperatingSystem.ProductCategory, sdkOperatingSystem.ProductCategory)
	assert.Equal(test_framework, cliOperatingSystem.ProductCode, sdkOperatingSystem.ProductCode)
	assert.Equal(test_framework, string(cliOperatingSystem.Location), string(sdkOperatingSystem.Location))
	assert.Equal(test_framework, cliOperatingSystem.YearMonth, sdkOperatingSystem.YearMonth)
	assert.Equal(test_framework, cliOperatingSystem.StartDateTime, sdkOperatingSystem.StartDateTime)
	assert.Equal(test_framework, cliOperatingSystem.EndDateTime, sdkOperatingSystem.EndDateTime)
	assert.Equal(test_framework, cliOperatingSystem.Cost, sdkOperatingSystem.Cost)
	assert.Equal(test_framework, cliOperatingSystem.PriceModel, sdkOperatingSystem.PriceModel)
	assert.Equal(test_framework, cliOperatingSystem.UnitPrice, sdkOperatingSystem.UnitPrice)
	assert.Equal(test_framework, cliOperatingSystem.UnitPriceDescription, sdkOperatingSystem.UnitPriceDescription)
	assert.Equal(test_framework, cliOperatingSystem.Quantity, sdkOperatingSystem.Quantity)
	assert.Equal(test_framework, cliOperatingSystem.Active, sdkOperatingSystem.Active)
	assert.Equal(test_framework, cliOperatingSystem.UsageSessionId, sdkOperatingSystem.UsageSessionId)
	assert.Equal(test_framework, cliOperatingSystem.CorrelationId, sdkOperatingSystem.CorrelationId)
	assert.Equal(test_framework, cliOperatingSystem.ReservationId, sdkOperatingSystem.ReservationId)

	assert.Equal(test_framework, cliOperatingSystem.Metadata.Cores, sdkOperatingSystem.Metadata.Cores)
	assert.Equal(test_framework, cliOperatingSystem.Metadata.CorrelationId, sdkOperatingSystem.Metadata.CorrelationId)
}

func assertEqualAsPublicSubnetRecord(
	test_framework *testing.T,
	cliRecord interface{},
	sdkPublicSubnet billingapi.PublicSubnetRecord,
) {
	cliPublicSubnet := cliRecord.(*ratedusageoneof.PublicSubnetRecord)

	assert.Equal(test_framework, cliPublicSubnet.Id, sdkPublicSubnet.Id)
	assert.Equal(test_framework, cliPublicSubnet.ProductCategory, sdkPublicSubnet.ProductCategory)
	assert.Equal(test_framework, cliPublicSubnet.ProductCode, sdkPublicSubnet.ProductCode)
	assert.Equal(test_framework, string(cliPublicSubnet.Location), string(sdkPublicSubnet.Location))
	assert.Equal(test_framework, cliPublicSubnet.YearMonth, sdkPublicSubnet.YearMonth)
	assert.Equal(test_framework, cliPublicSubnet.StartDateTime, sdkPublicSubnet.StartDateTime)
	assert.Equal(test_framework, cliPublicSubnet.EndDateTime, sdkPublicSubnet.EndDateTime)
	assert.Equal(test_framework, cliPublicSubnet.Cost, sdkPublicSubnet.Cost)
	assert.Equal(test_framework, cliPublicSubnet.PriceModel, sdkPublicSubnet.PriceModel)
	assert.Equal(test_framework, cliPublicSubnet.UnitPrice, sdkPublicSubnet.UnitPrice)
	assert.Equal(test_framework, cliPublicSubnet.UnitPriceDescription, sdkPublicSubnet.UnitPriceDescription)
	assert.Equal(test_framework, cliPublicSubnet.Quantity, sdkPublicSubnet.Quantity)
	assert.Equal(test_framework, cliPublicSubnet.Active, sdkPublicSubnet.Active)
	assert.Equal(test_framework, cliPublicSubnet.UsageSessionId, sdkPublicSubnet.UsageSessionId)
	assert.Equal(test_framework, cliPublicSubnet.CorrelationId, sdkPublicSubnet.CorrelationId)
	assert.Equal(test_framework, cliPublicSubnet.ReservationId, sdkPublicSubnet.ReservationId)

	assert.Equal(test_framework, cliPublicSubnet.Metadata.Id, sdkPublicSubnet.Metadata.Id)
	assert.Equal(test_framework, cliPublicSubnet.Metadata.Cidr, sdkPublicSubnet.Metadata.Cidr)
	assert.Equal(test_framework, cliPublicSubnet.Metadata.Size, sdkPublicSubnet.Metadata.Size)
}

func assertEqualAsServerRecord(
	test_framework *testing.T,
	cliRecord interface{},
	sdkServer billingapi.ServerRecord,
) {
	cliServer := cliRecord.(*ratedusageoneof.ServerRecord)

	assert.Equal(test_framework, cliServer.Id, sdkServer.Id)
	assert.Equal(test_framework, cliServer.ProductCategory, sdkServer.ProductCategory)
	assert.Equal(test_framework, cliServer.ProductCode, sdkServer.ProductCode)
	assert.Equal(test_framework, string(cliServer.Location), string(sdkServer.Location))
	assert.Equal(test_framework, cliServer.YearMonth, sdkServer.YearMonth)
	assert.Equal(test_framework, cliServer.StartDateTime, sdkServer.StartDateTime)
	assert.Equal(test_framework, cliServer.EndDateTime, sdkServer.EndDateTime)
	assert.Equal(test_framework, cliServer.Cost, sdkServer.Cost)
	assert.Equal(test_framework, cliServer.PriceModel, sdkServer.PriceModel)
	assert.Equal(test_framework, cliServer.UnitPrice, sdkServer.UnitPrice)
	assert.Equal(test_framework, cliServer.UnitPriceDescription, sdkServer.UnitPriceDescription)
	assert.Equal(test_framework, cliServer.Quantity, sdkServer.Quantity)
	assert.Equal(test_framework, cliServer.Active, sdkServer.Active)
	assert.Equal(test_framework, cliServer.UsageSessionId, sdkServer.UsageSessionId)
	assert.Equal(test_framework, cliServer.CorrelationId, sdkServer.CorrelationId)
	assert.Equal(test_framework, cliServer.ReservationId, sdkServer.ReservationId)

	assert.Equal(test_framework, cliServer.Metadata.Id, sdkServer.Metadata.Id)
	assert.Equal(test_framework, cliServer.Metadata.Hostname, sdkServer.Metadata.Hostname)
}
