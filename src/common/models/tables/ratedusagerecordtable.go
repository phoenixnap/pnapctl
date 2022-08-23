package tables

import (
	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels/ratedusageoneof"
)

var ONE_OF_TYPES = []string{
	"BANDWIDTH_RECORD",
	"OPERATING_SYSTEM_RECORD",
	"PUBLIC_SUBNET_RECORD",
	"SERVER_RECORD",
}

// Metadata keys
const (
	// Bandwidth Record
	EGRESS_GB        string = "Egress (GB)"
	INGRESS_GB       string = "Ingress (GB)"
	PACKAGE_QUANTITY string = "Package Quantity"
	PACKAGE_UNIT     string = "Package Unit"

	// Operating System Record
	CORES          string = "OS Cores"
	CORRELATION_ID string = "Correlation ID"

	// Public Subnet Record
	CIDR      string = "Subnet Cidr"
	SUBNET_ID string = "Subnet ID"
	SIZE      string = "Size"

	// Server Record
	SERVER_ID string = "Server Id"
	HOSTNAME  string = "Hostname"
)

// Full Table

type RatedUsageRecordTable struct {
	Id                   string                        `header:"Id"`
	ProductCategory      ratedusageoneof.Discriminator `header:"Product Category"`
	ProductCode          string                        `header:"Product Code"`
	Location             string                        `header:"Location"`
	YearMonth            string                        `header:"Year Month"`
	StartDateTime        string                        `header:"Start Date Time"`
	EndDateTime          string                        `header:"End Date Time"`
	Cost                 int64                         `header:"Cost"`
	PriceModel           string                        `header:"Price Model"`
	UnitPrice            float32                       `header:"Unit Price"`
	UnitPriceDescription string                        `header:"Unit Price Description"`
	Quantity             float32                       `header:"Quantity"`
	Active               bool                          `header:"Active"`
	UsageSessionId       string                        `header:"Usage Session Id"`
	CorrelationId        string                        `header:"Correlation Id"`
	ReservationId        string                        `header:"Reservation Id"`
	Metadata             map[string]interface{}        `header:"Metadata"`
}

func RatedUsageRecordFromSdk(sdkRecord billingapisdk.RatedUsageGet200ResponseInner) RatedUsageRecordTable {
	ratedUsage := parseCommon(sdkRecord)
	ratedUsage.attachMetadata(sdkRecord)
	return ratedUsage
}

func parseCommon(sdkRecord billingapisdk.RatedUsageGet200ResponseInner) RatedUsageRecordTable {
	ratedUsage := ratedusageoneof.RatedUsageFromSdkOneOf(&sdkRecord)

	return RatedUsageRecordTable{
		Id:                   ratedUsage.Id,
		ProductCategory:      ratedUsage.ProductCategory,
		ProductCode:          ratedUsage.ProductCode,
		Location:             ratedUsage.Location,
		YearMonth:            ratedUsage.YearMonth,
		StartDateTime:        ratedUsage.StartDateTime.String(),
		EndDateTime:          ratedUsage.EndDateTime.String(),
		Cost:                 ratedUsage.Cost,
		PriceModel:           ratedUsage.PriceModel,
		UnitPrice:            ratedUsage.UnitPrice,
		UnitPriceDescription: ratedUsage.UnitPriceDescription,
		Quantity:             ratedUsage.Quantity,
		Active:               ratedUsage.Active,
		UsageSessionId:       ratedUsage.UsageSessionId,
		CorrelationId:        ratedUsage.CorrelationId,
		ReservationId:        ratedUsage.ReservationId,
	}
}

func (table *RatedUsageRecordTable) attachMetadata(sdkRecord billingapisdk.RatedUsageGet200ResponseInner) {
	switch table.ProductCategory {

	case ratedusageoneof.BANDWIDTH:
		table.Metadata = map[string]interface{}{
			EGRESS_GB:        sdkRecord.BandwidthRecord.Metadata.EgressGb,
			INGRESS_GB:       sdkRecord.BandwidthRecord.Metadata.IngressGb,
			PACKAGE_QUANTITY: sdkRecord.BandwidthRecord.Metadata.PackageQuantity,
			PACKAGE_UNIT:     sdkRecord.BandwidthRecord.Metadata.PackageUnit,
		}

	case ratedusageoneof.OPERATING_SYSTEM:
		table.Metadata = map[string]interface{}{
			CORES:          sdkRecord.OperatingSystemRecord.Metadata.Cores,
			CORRELATION_ID: sdkRecord.OperatingSystemRecord.Metadata.CorrelationId,
		}

	case ratedusageoneof.PUBLIC_SUBNET:
		table.Metadata = map[string]interface{}{
			CIDR:      sdkRecord.PublicSubnetRecord.Metadata.Cidr,
			SUBNET_ID: sdkRecord.PublicSubnetRecord.Metadata.Id,
			SIZE:      sdkRecord.PublicSubnetRecord.Metadata.Size,
		}

	case ratedusageoneof.SERVER:
		table.Metadata = map[string]interface{}{
			SERVER_ID: sdkRecord.ServerRecord.Metadata.Id,
			HOSTNAME:  sdkRecord.ServerRecord.Metadata.Hostname,
		}
	}
}

// Short Version

type ShortRatedUsageRecordTable struct {
	Id              string                        `header:"Id"`
	ProductCategory ratedusageoneof.Discriminator `header:"Product Category"`
	ProductCode     string                        `header:"Product Code"`
	YearMonth       string                        `header:"Year Month"`
	Cost            int64                         `header:"Cost"`
}

// Extracts a ShortRatedUsageRecordTable using the full table.
func ShortRatedUsageRecordFromSdk(sdkRecord billingapisdk.RatedUsageGet200ResponseInner) *ShortRatedUsageRecordTable {
	fullTable := parseCommon(sdkRecord)

	return &ShortRatedUsageRecordTable{
		Id:              fullTable.Id,
		ProductCategory: fullTable.ProductCategory,
		ProductCode:     fullTable.ProductCode,
		YearMonth:       fullTable.YearMonth,
		Cost:            fullTable.Cost,
	}
}
