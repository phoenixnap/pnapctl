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
	cliOneOf interface{},
	sdkBandwidth billingapi.BandwidthRecord,
) {
	cliBandwidth := cliOneOf.(*ratedusageoneof.BandwidthRecord)

	assert.Equal(test_framework, cliBandwidth.Id, sdkBandwidth.GetId())
	assert.Equal(test_framework, string(cliBandwidth.ProductCategory), sdkBandwidth.GetProductCategory())
	assert.Equal(test_framework, cliBandwidth.ProductCode, sdkBandwidth.GetProductCode())
	assert.Equal(test_framework, billingapi.LocationEnum(cliBandwidth.Location), sdkBandwidth.GetLocation())
	assert.Equal(test_framework, cliBandwidth.YearMonth, sdkBandwidth.GetYearMonth())
	assert.Equal(test_framework, cliBandwidth.StartDateTime, sdkBandwidth.GetStartDateTime())
	assert.Equal(test_framework, cliBandwidth.EndDateTime, sdkBandwidth.GetEndDateTime())
	assert.Equal(test_framework, cliBandwidth.Cost, sdkBandwidth.GetCost())
	assert.Equal(test_framework, cliBandwidth.PriceModel, sdkBandwidth.GetPriceModel())
	assert.Equal(test_framework, cliBandwidth.UnitPrice, sdkBandwidth.GetUnitPrice())
	assert.Equal(test_framework, cliBandwidth.UnitPriceDescription, sdkBandwidth.GetUnitPriceDescription())
	assert.Equal(test_framework, cliBandwidth.Quantity, sdkBandwidth.GetQuantity())
	assert.Equal(test_framework, cliBandwidth.Active, sdkBandwidth.GetActive())
	assert.Equal(test_framework, cliBandwidth.UsageSessionId, sdkBandwidth.GetUsageSessionId())
	assert.Equal(test_framework, cliBandwidth.CorrelationId, sdkBandwidth.GetCorrelationId())
	assert.Equal(test_framework, cliBandwidth.ReservationId, sdkBandwidth.GetReservationId())

	assert.Equal(test_framework, cliBandwidth.Metadata.IngressGb, sdkBandwidth.Metadata.IngressGb)
	assert.Equal(test_framework, cliBandwidth.Metadata.EgressGb, sdkBandwidth.Metadata.EgressGb)
	assert.Equal(test_framework, cliBandwidth.Metadata.PackageQuantity, sdkBandwidth.Metadata.PackageQuantity)
	assert.Equal(test_framework, cliBandwidth.Metadata.PackageUnit, sdkBandwidth.Metadata.PackageUnit)
}

func assertEqualAsOperatingSystemRecord(
	test_framework *testing.T,
	cliOneOf interface{},
	sdkOperatingSystem billingapi.OperatingSystemRecord,
) {
	cliOperatingSystem := cliOneOf.(*ratedusageoneof.OperatingSystemRecord)

	assert.Equal(test_framework, cliOperatingSystem.Id, sdkOperatingSystem.GetId())
	assert.Equal(test_framework, string(cliOperatingSystem.ProductCategory), sdkOperatingSystem.GetProductCategory())
	assert.Equal(test_framework, cliOperatingSystem.ProductCode, sdkOperatingSystem.GetProductCode())
	assert.Equal(test_framework, billingapi.LocationEnum(cliOperatingSystem.Location), sdkOperatingSystem.GetLocation())
	assert.Equal(test_framework, cliOperatingSystem.YearMonth, sdkOperatingSystem.GetYearMonth())
	assert.Equal(test_framework, cliOperatingSystem.StartDateTime, sdkOperatingSystem.GetStartDateTime())
	assert.Equal(test_framework, cliOperatingSystem.EndDateTime, sdkOperatingSystem.GetEndDateTime())
	assert.Equal(test_framework, cliOperatingSystem.Cost, sdkOperatingSystem.GetCost())
	assert.Equal(test_framework, cliOperatingSystem.PriceModel, sdkOperatingSystem.GetPriceModel())
	assert.Equal(test_framework, cliOperatingSystem.UnitPrice, sdkOperatingSystem.GetUnitPrice())
	assert.Equal(test_framework, cliOperatingSystem.UnitPriceDescription, sdkOperatingSystem.GetUnitPriceDescription())
	assert.Equal(test_framework, cliOperatingSystem.Quantity, sdkOperatingSystem.GetQuantity())
	assert.Equal(test_framework, cliOperatingSystem.Active, sdkOperatingSystem.GetActive())
	assert.Equal(test_framework, cliOperatingSystem.UsageSessionId, sdkOperatingSystem.GetUsageSessionId())
	assert.Equal(test_framework, cliOperatingSystem.CorrelationId, sdkOperatingSystem.GetCorrelationId())
	assert.Equal(test_framework, cliOperatingSystem.ReservationId, sdkOperatingSystem.GetReservationId())

	assert.Equal(test_framework, cliOperatingSystem.Metadata.Cores, sdkOperatingSystem.Metadata.Cores)
	assert.Equal(test_framework, cliOperatingSystem.Metadata.CorrelationId, sdkOperatingSystem.Metadata.CorrelationId)
}

func assertEqualAsPublicSubnetRecord(
	test_framework *testing.T,
	cliOneOf interface{},
	sdkPublicSubnet billingapi.PublicSubnetRecord,
) {
	cliPublicSubnet := cliOneOf.(*ratedusageoneof.PublicSubnetRecord)

	assert.Equal(test_framework, cliPublicSubnet.Id, sdkPublicSubnet.GetId())
	assert.Equal(test_framework, string(cliPublicSubnet.ProductCategory), sdkPublicSubnet.GetProductCategory())
	assert.Equal(test_framework, cliPublicSubnet.ProductCode, sdkPublicSubnet.GetProductCode())
	assert.Equal(test_framework, billingapi.LocationEnum(cliPublicSubnet.Location), sdkPublicSubnet.GetLocation())
	assert.Equal(test_framework, cliPublicSubnet.YearMonth, sdkPublicSubnet.GetYearMonth())
	assert.Equal(test_framework, cliPublicSubnet.StartDateTime, sdkPublicSubnet.GetStartDateTime())
	assert.Equal(test_framework, cliPublicSubnet.EndDateTime, sdkPublicSubnet.GetEndDateTime())
	assert.Equal(test_framework, cliPublicSubnet.Cost, sdkPublicSubnet.GetCost())
	assert.Equal(test_framework, cliPublicSubnet.PriceModel, sdkPublicSubnet.GetPriceModel())
	assert.Equal(test_framework, cliPublicSubnet.UnitPrice, sdkPublicSubnet.GetUnitPrice())
	assert.Equal(test_framework, cliPublicSubnet.UnitPriceDescription, sdkPublicSubnet.GetUnitPriceDescription())
	assert.Equal(test_framework, cliPublicSubnet.Quantity, sdkPublicSubnet.GetQuantity())
	assert.Equal(test_framework, cliPublicSubnet.Active, sdkPublicSubnet.GetActive())
	assert.Equal(test_framework, cliPublicSubnet.UsageSessionId, sdkPublicSubnet.GetUsageSessionId())
	assert.Equal(test_framework, cliPublicSubnet.CorrelationId, sdkPublicSubnet.GetCorrelationId())
	assert.Equal(test_framework, cliPublicSubnet.ReservationId, sdkPublicSubnet.GetReservationId())

	assert.Equal(test_framework, cliPublicSubnet.Metadata.Id, sdkPublicSubnet.Metadata.Id)
	assert.Equal(test_framework, cliPublicSubnet.Metadata.Cidr, sdkPublicSubnet.Metadata.Cidr)
	assert.Equal(test_framework, cliPublicSubnet.Metadata.Size, sdkPublicSubnet.Metadata.Size)
}

func assertEqualAsServerRecord(
	test_framework *testing.T,
	cliOneOf interface{},
	sdkServer billingapi.ServerRecord,
) {
	cliServer := cliOneOf.(*ratedusageoneof.ServerRecord)

	assert.Equal(test_framework, cliServer.Id, sdkServer.GetId())
	assert.Equal(test_framework, string(cliServer.ProductCategory), sdkServer.GetProductCategory())
	assert.Equal(test_framework, cliServer.ProductCode, sdkServer.GetProductCode())
	assert.Equal(test_framework, billingapi.LocationEnum(cliServer.Location), sdkServer.GetLocation())
	assert.Equal(test_framework, cliServer.YearMonth, sdkServer.GetYearMonth())
	assert.Equal(test_framework, cliServer.StartDateTime, sdkServer.GetStartDateTime())
	assert.Equal(test_framework, cliServer.EndDateTime, sdkServer.GetEndDateTime())
	assert.Equal(test_framework, cliServer.Cost, sdkServer.GetCost())
	assert.Equal(test_framework, cliServer.PriceModel, sdkServer.GetPriceModel())
	assert.Equal(test_framework, cliServer.UnitPrice, sdkServer.GetUnitPrice())
	assert.Equal(test_framework, cliServer.UnitPriceDescription, sdkServer.GetUnitPriceDescription())
	assert.Equal(test_framework, cliServer.Quantity, sdkServer.GetQuantity())
	assert.Equal(test_framework, cliServer.Active, sdkServer.GetActive())
	assert.Equal(test_framework, cliServer.UsageSessionId, sdkServer.GetUsageSessionId())
	assert.Equal(test_framework, cliServer.CorrelationId, sdkServer.GetCorrelationId())
	assert.Equal(test_framework, cliServer.ReservationId, sdkServer.GetReservationId())

	assert.Equal(test_framework, cliServer.Metadata.Id, sdkServer.Metadata.Id)
	assert.Equal(test_framework, cliServer.Metadata.Hostname, sdkServer.Metadata.Hostname)
}
