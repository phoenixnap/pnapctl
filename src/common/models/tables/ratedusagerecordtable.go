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
	EGRESS_GB        = "Egress (GB)"
	INGRESS_GB       = "Ingress (GB)"
	PACKAGE_QUANTITY = "Package Quantity"
	PACKAGE_UNIT     = "Package Unit"

	// Operating System Record
	CORES          = "OS Cores"
	CORRELATION_ID = "Correlation ID"

	// Public Subnet Record
	CIDR      = "Subnet Cidr"
	SUBNET_ID = "Subnet ID"
	SIZE      = "Size"

	// Server Record
	SERVER_ID = "Server Id"
	HOSTNAME  = "Hostname"
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

func RatedUsageRecordTableFromSdk(sdk billingapisdk.RatedUsageGet200ResponseInner) RatedUsageRecordTable {
	ratedUsage := RatedUsageRecordTable{}
	ratedUsage.parseCommon(sdk)
	ratedUsage.attachMetadata(sdk)
	return ratedUsage
}

func (r *RatedUsageRecordTable) parseCommon(sdk billingapisdk.RatedUsageGet200ResponseInner) {
	ratedUsage := ratedusageoneof.RatedUsageFromSdkOneOf(&sdk)

	r.Id = ratedUsage.Id
	r.ProductCategory = ratedUsage.ProductCategory
	r.ProductCode = ratedUsage.ProductCode
	r.Location = ratedUsage.Location
	r.YearMonth = ratedUsage.YearMonth
	r.StartDateTime = ratedUsage.StartDateTime.String()
	r.EndDateTime = ratedUsage.EndDateTime.String()
	r.Cost = ratedUsage.Cost
	r.PriceModel = ratedUsage.PriceModel
	r.UnitPrice = ratedUsage.UnitPrice
	r.UnitPriceDescription = ratedUsage.UnitPriceDescription
	r.Quantity = ratedUsage.Quantity
	r.Active = ratedUsage.Active
	r.UsageSessionId = ratedUsage.UsageSessionId
	r.CorrelationId = ratedUsage.CorrelationId
	r.ReservationId = ratedUsage.ReservationId
}

func (table *RatedUsageRecordTable) attachMetadata(sdk billingapisdk.RatedUsageGet200ResponseInner) {
	switch table.ProductCategory {

	case ratedusageoneof.BANDWIDTH:
		table.Metadata = map[string]interface{}{
			EGRESS_GB:        sdk.BandwidthRecord.Metadata.EgressGb,
			INGRESS_GB:       sdk.BandwidthRecord.Metadata.IngressGb,
			PACKAGE_QUANTITY: sdk.BandwidthRecord.Metadata.PackageQuantity,
			PACKAGE_UNIT:     sdk.BandwidthRecord.Metadata.PackageUnit,
		}

	case ratedusageoneof.OPERATING_SYSTEM:
		table.Metadata = map[string]interface{}{
			CORES:          sdk.OperatingSystemRecord.Metadata.Cores,
			CORRELATION_ID: sdk.OperatingSystemRecord.Metadata.CorrelationId,
		}

	case ratedusageoneof.PUBLIC_SUBNET:
		table.Metadata = map[string]interface{}{
			CIDR:      sdk.PublicSubnetRecord.Metadata.Cidr,
			SUBNET_ID: sdk.PublicSubnetRecord.Metadata.Id,
			SIZE:      sdk.PublicSubnetRecord.Metadata.Size,
		}

	case ratedusageoneof.SERVER:
		table.Metadata = map[string]interface{}{
			SERVER_ID: sdk.ServerRecord.Metadata.Id,
			HOSTNAME:  sdk.ServerRecord.Metadata.Hostname,
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
func ShortRatedUsageRecordFromSdk(sdk billingapisdk.RatedUsageGet200ResponseInner) *ShortRatedUsageRecordTable {
	fullTable := RatedUsageRecordTable{}
	fullTable.parseCommon(sdk)

	return &ShortRatedUsageRecordTable{
		Id:              fullTable.Id,
		ProductCategory: fullTable.ProductCategory,
		ProductCode:     fullTable.ProductCode,
		YearMonth:       fullTable.YearMonth,
		Cost:            fullTable.Cost,
	}
}
