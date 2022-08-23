package tables

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
)

// Full version
func TestRatedUsageRecordFromBandwidthSdk(test_framework *testing.T) {
	record := billingapi.RatedUsageGet200ResponseInner{
		BandwidthRecord: billingmodels.GenerateBandwidthRecordSdk(),
	}
	table, err := RatedUsageRecordFromSdk(record, "get rated-usage")

	assert.Nil(test_framework, err)
	assertFullBandwidthRecordsEqual(test_framework, *record.BandwidthRecord, table)
}

func TestRatedUsageRecordFromOperatingSystemSdk(test_framework *testing.T) {
	record := billingapi.RatedUsageGet200ResponseInner{
		OperatingSystemRecord: billingmodels.GenerateOperatingSystemRecordSdk(),
	}
	table, err := RatedUsageRecordFromSdk(record, "get rated-usage")

	assert.Nil(test_framework, err)
	assertFullOperatingSystemRecordsEqual(test_framework, *record.OperatingSystemRecord, table)
}

func TestRatedUsageRecordFromPublicSubnetSdk(test_framework *testing.T) {
	record := billingapi.RatedUsageGet200ResponseInner{
		PublicSubnetRecord: billingmodels.GeneratePublicSubnetRecordSdk(),
	}
	table, err := RatedUsageRecordFromSdk(record, "get rated-usage")

	assert.Nil(test_framework, err)
	assertFullPublicSubnetRecordsEqual(test_framework, *record.PublicSubnetRecord, table)
}

func TestRatedUsageRecordFromServerSdk(test_framework *testing.T) {
	record := billingapi.RatedUsageGet200ResponseInner{
		ServerRecord: billingmodels.GenerateServerRecordSdk(),
	}
	table, err := RatedUsageRecordFromSdk(record, "get rated-usage")

	assert.Nil(test_framework, err)
	assertFullServerRecordsEqual(test_framework, *record.ServerRecord, table)
}

func TestRatedUsageRecordFromSdk_NoFieldsError(test_framework *testing.T) {
	record := billingapi.RatedUsageGet200ResponseInner{}
	_, err := RatedUsageRecordFromSdk(record, "get rated-usage")

	assert.NotNil(test_framework, err)
	expectedMessage := "Command 'get rated-usage' has been performed, but something went wrong. Error code: " + ctlerrors.OneOfNoFieldsPopulated
	assert.Equal(test_framework, err.Error(), expectedMessage)
}

// Full assertions
func assertFullBandwidthRecordsEqual(test_framework *testing.T, bandwidthRecord billingapi.BandwidthRecord, cliOneOf RatedUsageRecordTable) {
	assert.Equal(test_framework, bandwidthRecord.Id, cliOneOf.Id)

	assert.Equal(test_framework, bandwidthRecord.Id, cliOneOf.Id)
	assert.Equal(test_framework, bandwidthRecord.ProductCategory, cliOneOf.ProductCategory)
	assert.Equal(test_framework, bandwidthRecord.ProductCode, cliOneOf.ProductCode)
	assert.Equal(test_framework, string(bandwidthRecord.Location), cliOneOf.Location)
	assert.Equal(test_framework, DerefString(bandwidthRecord.YearMonth), cliOneOf.YearMonth)
	assert.Equal(test_framework, bandwidthRecord.StartDateTime.String(), cliOneOf.StartDateTime)
	assert.Equal(test_framework, bandwidthRecord.EndDateTime.String(), cliOneOf.EndDateTime)
	assert.Equal(test_framework, bandwidthRecord.Cost, cliOneOf.Cost)
	assert.Equal(test_framework, bandwidthRecord.PriceModel, cliOneOf.PriceModel)
	assert.Equal(test_framework, bandwidthRecord.UnitPrice, cliOneOf.UnitPrice)
	assert.Equal(test_framework, bandwidthRecord.UnitPriceDescription, cliOneOf.UnitPriceDescription)
	assert.Equal(test_framework, bandwidthRecord.Quantity, cliOneOf.Quantity)
	assert.Equal(test_framework, bandwidthRecord.Active, cliOneOf.Active)
	assert.Equal(test_framework, bandwidthRecord.UsageSessionId, cliOneOf.UsageSessionId)
	assert.Equal(test_framework, bandwidthRecord.CorrelationId, cliOneOf.CorrelationId)
	assert.Equal(test_framework, DerefString(bandwidthRecord.ReservationId), cliOneOf.ReservationId)

	assert.Equal(test_framework, bandwidthRecord.Metadata.IngressGb, cliOneOf.Metadata[INGRESS_GB])
	assert.Equal(test_framework, bandwidthRecord.Metadata.EgressGb, cliOneOf.Metadata[EGRESS_GB])
	assert.Equal(test_framework, bandwidthRecord.Metadata.PackageQuantity, cliOneOf.Metadata[PACKAGE_QUANTITY])
	assert.Equal(test_framework, bandwidthRecord.Metadata.PackageUnit, cliOneOf.Metadata[PACKAGE_UNIT])
}

func assertFullOperatingSystemRecordsEqual(test_framework *testing.T, operatingSystemRecord billingapi.OperatingSystemRecord, cliOneOf RatedUsageRecordTable) {
	assert.Equal(test_framework, operatingSystemRecord.Id, cliOneOf.Id)

	assert.Equal(test_framework, operatingSystemRecord.Id, cliOneOf.Id)
	assert.Equal(test_framework, operatingSystemRecord.ProductCategory, cliOneOf.ProductCategory)
	assert.Equal(test_framework, operatingSystemRecord.ProductCode, cliOneOf.ProductCode)
	assert.Equal(test_framework, string(operatingSystemRecord.Location), cliOneOf.Location)
	assert.Equal(test_framework, DerefString(operatingSystemRecord.YearMonth), cliOneOf.YearMonth)
	assert.Equal(test_framework, operatingSystemRecord.StartDateTime.String(), cliOneOf.StartDateTime)
	assert.Equal(test_framework, operatingSystemRecord.EndDateTime.String(), cliOneOf.EndDateTime)
	assert.Equal(test_framework, operatingSystemRecord.Cost, cliOneOf.Cost)
	assert.Equal(test_framework, operatingSystemRecord.PriceModel, cliOneOf.PriceModel)
	assert.Equal(test_framework, operatingSystemRecord.UnitPrice, cliOneOf.UnitPrice)
	assert.Equal(test_framework, operatingSystemRecord.UnitPriceDescription, cliOneOf.UnitPriceDescription)
	assert.Equal(test_framework, operatingSystemRecord.Quantity, cliOneOf.Quantity)
	assert.Equal(test_framework, operatingSystemRecord.Active, cliOneOf.Active)
	assert.Equal(test_framework, operatingSystemRecord.UsageSessionId, cliOneOf.UsageSessionId)
	assert.Equal(test_framework, operatingSystemRecord.CorrelationId, cliOneOf.CorrelationId)
	assert.Equal(test_framework, DerefString(operatingSystemRecord.ReservationId), cliOneOf.ReservationId)

	assert.Equal(test_framework, operatingSystemRecord.Metadata.Cores, cliOneOf.Metadata[CORES])
	assert.Equal(test_framework, operatingSystemRecord.Metadata.CorrelationId, cliOneOf.Metadata[CORRELATION_ID])
}

func assertFullPublicSubnetRecordsEqual(test_framework *testing.T, publicSubnetRecord billingapi.PublicSubnetRecord, cliOneOf RatedUsageRecordTable) {
	assert.Equal(test_framework, publicSubnetRecord.Id, cliOneOf.Id)

	assert.Equal(test_framework, publicSubnetRecord.Id, cliOneOf.Id)
	assert.Equal(test_framework, publicSubnetRecord.ProductCategory, cliOneOf.ProductCategory)
	assert.Equal(test_framework, publicSubnetRecord.ProductCode, cliOneOf.ProductCode)
	assert.Equal(test_framework, string(publicSubnetRecord.Location), cliOneOf.Location)
	assert.Equal(test_framework, DerefString(publicSubnetRecord.YearMonth), cliOneOf.YearMonth)
	assert.Equal(test_framework, publicSubnetRecord.StartDateTime.String(), cliOneOf.StartDateTime)
	assert.Equal(test_framework, publicSubnetRecord.EndDateTime.String(), cliOneOf.EndDateTime)
	assert.Equal(test_framework, publicSubnetRecord.Cost, cliOneOf.Cost)
	assert.Equal(test_framework, publicSubnetRecord.PriceModel, cliOneOf.PriceModel)
	assert.Equal(test_framework, publicSubnetRecord.UnitPrice, cliOneOf.UnitPrice)
	assert.Equal(test_framework, publicSubnetRecord.UnitPriceDescription, cliOneOf.UnitPriceDescription)
	assert.Equal(test_framework, publicSubnetRecord.Quantity, cliOneOf.Quantity)
	assert.Equal(test_framework, publicSubnetRecord.Active, cliOneOf.Active)
	assert.Equal(test_framework, publicSubnetRecord.UsageSessionId, cliOneOf.UsageSessionId)
	assert.Equal(test_framework, publicSubnetRecord.CorrelationId, cliOneOf.CorrelationId)
	assert.Equal(test_framework, DerefString(publicSubnetRecord.ReservationId), cliOneOf.ReservationId)

	assert.Equal(test_framework, publicSubnetRecord.Metadata.Cidr, cliOneOf.Metadata[CIDR])
	assert.Equal(test_framework, publicSubnetRecord.Metadata.Id, cliOneOf.Metadata[SUBNET_ID])
	assert.Equal(test_framework, publicSubnetRecord.Metadata.Size, cliOneOf.Metadata[SIZE])
}

func assertFullServerRecordsEqual(test_framework *testing.T, serverRecord billingapi.ServerRecord, cliOneOf RatedUsageRecordTable) {
	assert.Equal(test_framework, serverRecord.Id, cliOneOf.Id)

	assert.Equal(test_framework, serverRecord.Id, cliOneOf.Id)
	assert.Equal(test_framework, serverRecord.ProductCategory, cliOneOf.ProductCategory)
	assert.Equal(test_framework, serverRecord.ProductCode, cliOneOf.ProductCode)
	assert.Equal(test_framework, string(serverRecord.Location), cliOneOf.Location)
	assert.Equal(test_framework, DerefString(serverRecord.YearMonth), cliOneOf.YearMonth)
	assert.Equal(test_framework, serverRecord.StartDateTime.String(), cliOneOf.StartDateTime)
	assert.Equal(test_framework, serverRecord.EndDateTime.String(), cliOneOf.EndDateTime)
	assert.Equal(test_framework, serverRecord.Cost, cliOneOf.Cost)
	assert.Equal(test_framework, serverRecord.PriceModel, cliOneOf.PriceModel)
	assert.Equal(test_framework, serverRecord.UnitPrice, cliOneOf.UnitPrice)
	assert.Equal(test_framework, serverRecord.UnitPriceDescription, cliOneOf.UnitPriceDescription)
	assert.Equal(test_framework, serverRecord.Quantity, cliOneOf.Quantity)
	assert.Equal(test_framework, serverRecord.Active, cliOneOf.Active)
	assert.Equal(test_framework, serverRecord.UsageSessionId, cliOneOf.UsageSessionId)
	assert.Equal(test_framework, serverRecord.CorrelationId, cliOneOf.CorrelationId)
	assert.Equal(test_framework, DerefString(serverRecord.ReservationId), cliOneOf.ReservationId)

	assert.Equal(test_framework, serverRecord.Metadata.Hostname, cliOneOf.Metadata[HOSTNAME])
	assert.Equal(test_framework, serverRecord.Metadata.Id, cliOneOf.Metadata[SERVER_ID])
}
