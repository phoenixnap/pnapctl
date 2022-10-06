package tables

import (
	"time"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
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
	Id                   string                     `header:"Id"`
	ProductCategory      string                     `header:"Product Category"`
	ProductCode          string                     `header:"Product Code"`
	Location             billingapisdk.LocationEnum `header:"Location"`
	YearMonth            string                     `header:"Year Month"`
	StartDateTime        string                     `header:"Start Date Time"`
	EndDateTime          string                     `header:"End Date Time"`
	Cost                 int64                      `header:"Cost"`
	PriceModel           string                     `header:"Price Model"`
	UnitPrice            float32                    `header:"Unit Price"`
	UnitPriceDescription string                     `header:"Unit Price Description"`
	Quantity             float32                    `header:"Quantity"`
	Active               bool                       `header:"Active"`
	UsageSessionId       string                     `header:"Usage Session Id"`
	CorrelationId        string                     `header:"Correlation Id"`
	ReservationId        string                     `header:"Reservation Id"`
	Metadata             map[string]interface{}     `header:"Metadata"`
}

func RatedUsageRecordTableFromSdk(sdk billingapisdk.RatedUsageGet200ResponseInner) *RatedUsageRecordTable {
	ratedUsage := parseCommonRatedUsage(sdk)
	if ratedUsage == nil {
		return nil
	}

	ratedUsage.attachUnique(sdk)
	return ratedUsage
}

func parseCommonRatedUsage(sdk billingapisdk.RatedUsageGet200ResponseInner) *RatedUsageRecordTable {
	ratedUsage := models.GetFromAllOf[billingapisdk.RatedUsageRecord](sdk)
	if ratedUsage == nil {
		return nil
	}

	return &RatedUsageRecordTable{
		Id:                   ratedUsage.Id,
		ProductCategory:      ratedUsage.ProductCategory,
		ProductCode:          ratedUsage.ProductCode,
		Location:             ratedUsage.Location,
		YearMonth:            DerefString(ratedUsage.YearMonth),
		StartDateTime:        ratedUsage.StartDateTime.Format(time.RFC1123),
		EndDateTime:          ratedUsage.EndDateTime.Format(time.RFC1123),
		Cost:                 ratedUsage.Cost,
		PriceModel:           ratedUsage.PriceModel,
		UnitPrice:            ratedUsage.UnitPrice,
		UnitPriceDescription: ratedUsage.UnitPriceDescription,
		Quantity:             ratedUsage.Quantity,
		Active:               ratedUsage.Active,
		UsageSessionId:       ratedUsage.UsageSessionId,
		CorrelationId:        ratedUsage.CorrelationId,
		ReservationId:        DerefString(ratedUsage.ReservationId),
	}
}

func (table *RatedUsageRecordTable) attachUnique(sdk billingapisdk.RatedUsageGet200ResponseInner) {
	switch table.ProductCategory {

	case generators.RatedUsageBandwidth:
		table.Metadata = map[string]interface{}{
			EGRESS_GB:        sdk.BandwidthRecord.Metadata.EgressGb,
			INGRESS_GB:       sdk.BandwidthRecord.Metadata.IngressGb,
			PACKAGE_QUANTITY: sdk.BandwidthRecord.Metadata.PackageQuantity,
			PACKAGE_UNIT:     sdk.BandwidthRecord.Metadata.PackageUnit,
		}

	case generators.RatedUsageOperatingSystem:
		table.Metadata = map[string]interface{}{
			CORES:          sdk.OperatingSystemRecord.Metadata.Cores,
			CORRELATION_ID: sdk.OperatingSystemRecord.Metadata.CorrelationId,
		}

	case generators.RatedUsagePublicSubnet:
		table.Metadata = map[string]interface{}{
			CIDR:      sdk.PublicSubnetRecord.Metadata.Cidr,
			SUBNET_ID: sdk.PublicSubnetRecord.Metadata.Id,
			SIZE:      sdk.PublicSubnetRecord.Metadata.Size,
		}

	case generators.RatedUsageServer:
		table.Metadata = map[string]interface{}{
			SERVER_ID: sdk.ServerRecord.Metadata.Id,
			HOSTNAME:  sdk.ServerRecord.Metadata.Hostname,
		}
	}
}

// Short Version

type ShortRatedUsageRecordTable struct {
	Id              string `header:"Id"`
	ProductCategory string `header:"Product Category"`
	ProductCode     string `header:"Product Code"`
	YearMonth       string `header:"Year Month"`
	Cost            int64  `header:"Cost"`
}

// Extracts a ShortRatedUsageRecordTable using the full table.
func ShortRatedUsageRecordFromSdk(sdk billingapisdk.RatedUsageGet200ResponseInner) *ShortRatedUsageRecordTable {
	fullTable := parseCommonRatedUsage(sdk)
	if fullTable == nil {
		return nil
	}

	return &ShortRatedUsageRecordTable{
		Id:              fullTable.Id,
		ProductCategory: fullTable.ProductCategory,
		ProductCode:     fullTable.ProductCode,
		YearMonth:       fullTable.YearMonth,
		Cost:            fullTable.Cost,
	}
}
