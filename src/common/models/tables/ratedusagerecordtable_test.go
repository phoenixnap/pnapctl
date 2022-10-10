package tables

import (
	"testing"
	"time"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
)

// Full version
func TestRatedUsageRecordFromBandwidthSdk(test_framework *testing.T) {
	record := billingapi.RatedUsageGet200ResponseInner{
		BandwidthRecord: generators.GenerateBandwidthRecordSdk(),
	}
	table := *RatedUsageRecordTableFromSdk(record)

	assertFullBandwidthRecordsEqual(test_framework, *record.BandwidthRecord, table)
}

func TestRatedUsageRecordFromOperatingSystemSdk(test_framework *testing.T) {
	record := billingapi.RatedUsageGet200ResponseInner{
		OperatingSystemRecord: generators.GenerateOperatingSystemRecordSdk(),
	}
	table := *RatedUsageRecordTableFromSdk(record)

	assertFullOperatingSystemRecordsEqual(test_framework, *record.OperatingSystemRecord, table)
}

func TestRatedUsageRecordFromPublicSubnetSdk(test_framework *testing.T) {
	record := billingapi.RatedUsageGet200ResponseInner{
		PublicSubnetRecord: generators.GeneratePublicSubnetRecordSdk(),
	}
	table := *RatedUsageRecordTableFromSdk(record)

	assertFullPublicSubnetRecordsEqual(test_framework, *record.PublicSubnetRecord, table)
}

func TestRatedUsageRecordFromServerSdk(test_framework *testing.T) {
	record := billingapi.RatedUsageGet200ResponseInner{
		ServerRecord: generators.GenerateServerRecordSdk(),
	}
	table := *RatedUsageRecordTableFromSdk(record)

	assertFullServerRecordsEqual(test_framework, *record.ServerRecord, table)
}

func TestRatedUsageRecordFromStorageSdk(test_framework *testing.T) {
	record := billingapi.RatedUsageGet200ResponseInner{
		StorageRecord: generators.GenerateStorageRecordSdk(),
	}
	table := *RatedUsageRecordTableFromSdk(record)

	assertFullStorageRecordsEqual(test_framework, *record.StorageRecord, table)
}

// Full assertions
func assertFullBandwidthRecordsEqual(test_framework *testing.T, bandwidthRecord billingapi.BandwidthRecord, cliOneOf RatedUsageRecordTable) {
	assert.Equal(test_framework, bandwidthRecord.Id, cliOneOf.Id)
	assert.Equal(test_framework, bandwidthRecord.ProductCategory, string(cliOneOf.ProductCategory))
	assert.Equal(test_framework, bandwidthRecord.ProductCode, cliOneOf.ProductCode)
	assert.Equal(test_framework, bandwidthRecord.Location, cliOneOf.Location)
	assert.Equal(test_framework, DerefString(bandwidthRecord.YearMonth), cliOneOf.YearMonth)
	assert.Equal(test_framework, bandwidthRecord.StartDateTime.Format(time.RFC1123), cliOneOf.StartDateTime)
	assert.Equal(test_framework, bandwidthRecord.EndDateTime.Format(time.RFC1123), cliOneOf.EndDateTime)
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
	assert.Equal(test_framework, operatingSystemRecord.ProductCategory, string(cliOneOf.ProductCategory))
	assert.Equal(test_framework, operatingSystemRecord.ProductCode, cliOneOf.ProductCode)
	assert.Equal(test_framework, operatingSystemRecord.Location, cliOneOf.Location)
	assert.Equal(test_framework, DerefString(operatingSystemRecord.YearMonth), cliOneOf.YearMonth)
	assert.Equal(test_framework, operatingSystemRecord.StartDateTime.Format(time.RFC1123), cliOneOf.StartDateTime)
	assert.Equal(test_framework, operatingSystemRecord.EndDateTime.Format(time.RFC1123), cliOneOf.EndDateTime)
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
	assert.Equal(test_framework, publicSubnetRecord.ProductCategory, string(cliOneOf.ProductCategory))
	assert.Equal(test_framework, publicSubnetRecord.ProductCode, cliOneOf.ProductCode)
	assert.Equal(test_framework, publicSubnetRecord.Location, cliOneOf.Location)
	assert.Equal(test_framework, DerefString(publicSubnetRecord.YearMonth), cliOneOf.YearMonth)
	assert.Equal(test_framework, publicSubnetRecord.StartDateTime.Format(time.RFC1123), cliOneOf.StartDateTime)
	assert.Equal(test_framework, publicSubnetRecord.EndDateTime.Format(time.RFC1123), cliOneOf.EndDateTime)
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
	assert.Equal(test_framework, serverRecord.ProductCategory, string(cliOneOf.ProductCategory))
	assert.Equal(test_framework, serverRecord.ProductCode, cliOneOf.ProductCode)
	assert.Equal(test_framework, serverRecord.Location, cliOneOf.Location)
	assert.Equal(test_framework, DerefString(serverRecord.YearMonth), cliOneOf.YearMonth)
	assert.Equal(test_framework, serverRecord.StartDateTime.Format(time.RFC1123), cliOneOf.StartDateTime)
	assert.Equal(test_framework, serverRecord.EndDateTime.Format(time.RFC1123), cliOneOf.EndDateTime)
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

func assertFullStorageRecordsEqual(test_framework *testing.T, storageRecord billingapi.StorageRecord, cliOneOf RatedUsageRecordTable) {
	assert.Equal(test_framework, storageRecord.Id, cliOneOf.Id)
	assert.Equal(test_framework, storageRecord.ProductCategory, string(cliOneOf.ProductCategory))
	assert.Equal(test_framework, storageRecord.ProductCode, cliOneOf.ProductCode)
	assert.Equal(test_framework, storageRecord.Location, cliOneOf.Location)
	assert.Equal(test_framework, DerefString(storageRecord.YearMonth), cliOneOf.YearMonth)
	assert.Equal(test_framework, storageRecord.StartDateTime.Format(time.RFC1123), cliOneOf.StartDateTime)
	assert.Equal(test_framework, storageRecord.EndDateTime.Format(time.RFC1123), cliOneOf.EndDateTime)
	assert.Equal(test_framework, storageRecord.Cost, cliOneOf.Cost)
	assert.Equal(test_framework, storageRecord.PriceModel, cliOneOf.PriceModel)
	assert.Equal(test_framework, storageRecord.UnitPrice, cliOneOf.UnitPrice)
	assert.Equal(test_framework, storageRecord.UnitPriceDescription, cliOneOf.UnitPriceDescription)
	assert.Equal(test_framework, storageRecord.Quantity, cliOneOf.Quantity)
	assert.Equal(test_framework, storageRecord.Active, cliOneOf.Active)
	assert.Equal(test_framework, storageRecord.UsageSessionId, cliOneOf.UsageSessionId)
	assert.Equal(test_framework, storageRecord.CorrelationId, cliOneOf.CorrelationId)
	assert.Equal(test_framework, DerefString(storageRecord.ReservationId), cliOneOf.ReservationId)

	assert.Equal(test_framework, storageRecord.Metadata.NetworkStorageId, cliOneOf.Metadata[NETWORK_STORAGE_ID])
	assert.Equal(test_framework, storageRecord.Metadata.NetworkStorageName, cliOneOf.Metadata[NETWORK_STORAGE_NAME])
	assert.Equal(test_framework, storageRecord.Metadata.VolumeId, cliOneOf.Metadata[VOLUME_ID])
	assert.Equal(test_framework, storageRecord.Metadata.VolumeName, cliOneOf.Metadata[VOLUME_NAME])
	assert.Equal(test_framework, storageRecord.Metadata.CapacityInGb, cliOneOf.Metadata[CAPACITY_IN_GB])
	assert.Equal(test_framework, storageRecord.Metadata.CreatedOn, cliOneOf.Metadata[CREATED_ON])
}
