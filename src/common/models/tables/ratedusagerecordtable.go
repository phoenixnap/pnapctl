package tables

import (
	"encoding/json"
	"fmt"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/ctlerrors"
)

var ONE_OF_TYPES = []string{
	"BANDWIDTH_RECORD",
	"OPERATING_SYSTEM_RECORD",
	"PUBLIC_SUBNET_RECORD",
	"SERVER_RECORD",
}

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
	Id                   string                 `header:"Id"`
	ProductCategory      string                 `header:"Product Category"`
	ProductCode          string                 `header:"Product Code"`
	Location             string                 `header:"Location"`
	YearMonth            string                 `header:"Year Month"`
	StartDateTime        string                 `header:"Start Date Time"`
	EndDateTime          string                 `header:"End Date Time"`
	Cost                 int64                  `header:"Cost"`
	PriceModel           string                 `header:"Price Model"`
	UnitPrice            float32                `header:"Unit Price"`
	UnitPriceDescription string                 `header:"Unit Price Description"`
	Quantity             float32                `header:"Quantity"`
	Active               bool                   `header:"Active"`
	UsageSessionId       string                 `header:"Usage Session Id"`
	CorrelationId        string                 `header:"Correlation Id"`
	ReservationId        string                 `header:"Reservation Id"`
	Metadata             map[string]interface{} `header:"Metadata"`
}

func RatedUsageRecordFromSdk(sdkRecord billingapisdk.RatedUsageGet200ResponseInner, commandName string) (RatedUsageRecordTable, error) {
	if sdkRecord.BandwidthRecord != nil {
		return fromBandwidthRecord(*sdkRecord.BandwidthRecord), nil
	} else if sdkRecord.OperatingSystemRecord != nil {
		return fromOperatingSystemRecord(*sdkRecord.OperatingSystemRecord), nil
	} else if sdkRecord.PublicSubnetRecord != nil {
		return fromPublicSubnetRecord(*sdkRecord.PublicSubnetRecord), nil
	} else if sdkRecord.ServerRecord != nil {
		return fromServerRecord(*sdkRecord.ServerRecord), nil
	} else {
		return RatedUsageRecordTable{}, ctlerrors.CreateCLIError(ctlerrors.OneOfNoFieldsPopulated, commandName,
			fmt.Errorf("RatedUsage was none of the following: %v. Couldn't turn into table", ONE_OF_TYPES))
	}
}

func fromBandwidthRecord(bandwidthRecord billingapisdk.BandwidthRecord) RatedUsageRecordTable {
	metadata := make(map[string]interface{})

	metadata[EGRESS_GB] = bandwidthRecord.Metadata.EgressGb
	metadata[INGRESS_GB] = bandwidthRecord.Metadata.IngressGb
	metadata[PACKAGE_QUANTITY] = bandwidthRecord.Metadata.PackageQuantity
	metadata[PACKAGE_UNIT] = bandwidthRecord.Metadata.PackageUnit

	return RatedUsageRecordTable{
		Id:                   bandwidthRecord.Id,
		ProductCategory:      bandwidthRecord.ProductCategory,
		ProductCode:          bandwidthRecord.ProductCode,
		Location:             string(bandwidthRecord.Location),
		YearMonth:            DerefString(bandwidthRecord.YearMonth),
		StartDateTime:        bandwidthRecord.StartDateTime.String(),
		EndDateTime:          bandwidthRecord.EndDateTime.String(),
		Cost:                 bandwidthRecord.Cost,
		PriceModel:           bandwidthRecord.PriceModel,
		UnitPrice:            bandwidthRecord.UnitPrice,
		UnitPriceDescription: bandwidthRecord.UnitPriceDescription,
		Quantity:             bandwidthRecord.Quantity,
		Active:               bandwidthRecord.Active,
		UsageSessionId:       bandwidthRecord.UsageSessionId,
		CorrelationId:        bandwidthRecord.CorrelationId,
		ReservationId:        DerefString(bandwidthRecord.ReservationId),
		Metadata:             metadata,
	}
}

func fromOperatingSystemRecord(operatingSystemRecord billingapisdk.OperatingSystemRecord) RatedUsageRecordTable {
	metadata := make(map[string]interface{})

	metadata[CORES] = operatingSystemRecord.Metadata.Cores
	metadata[CORRELATION_ID] = operatingSystemRecord.Metadata.CorrelationId

	return RatedUsageRecordTable{
		Id:                   operatingSystemRecord.Id,
		ProductCategory:      operatingSystemRecord.ProductCategory,
		ProductCode:          operatingSystemRecord.ProductCode,
		Location:             string(operatingSystemRecord.Location),
		YearMonth:            DerefString(operatingSystemRecord.YearMonth),
		StartDateTime:        operatingSystemRecord.StartDateTime.String(),
		EndDateTime:          operatingSystemRecord.EndDateTime.String(),
		Cost:                 operatingSystemRecord.Cost,
		PriceModel:           operatingSystemRecord.PriceModel,
		UnitPrice:            operatingSystemRecord.UnitPrice,
		UnitPriceDescription: operatingSystemRecord.UnitPriceDescription,
		Quantity:             operatingSystemRecord.Quantity,
		Active:               operatingSystemRecord.Active,
		UsageSessionId:       operatingSystemRecord.UsageSessionId,
		CorrelationId:        operatingSystemRecord.CorrelationId,
		ReservationId:        DerefString(operatingSystemRecord.ReservationId),
		Metadata:             metadata,
	}
}

func fromPublicSubnetRecord(publicSubnetRecord billingapisdk.PublicSubnetRecord) RatedUsageRecordTable {
	metadata := make(map[string]interface{})

	metadata[CIDR] = publicSubnetRecord.Metadata.Cidr
	metadata[SUBNET_ID] = publicSubnetRecord.Metadata.Id
	metadata[SIZE] = publicSubnetRecord.Metadata.Size

	return RatedUsageRecordTable{
		Id:                   publicSubnetRecord.Id,
		ProductCategory:      publicSubnetRecord.ProductCategory,
		ProductCode:          publicSubnetRecord.ProductCode,
		Location:             string(publicSubnetRecord.Location),
		YearMonth:            DerefString(publicSubnetRecord.YearMonth),
		StartDateTime:        publicSubnetRecord.StartDateTime.String(),
		EndDateTime:          publicSubnetRecord.EndDateTime.String(),
		Cost:                 publicSubnetRecord.Cost,
		PriceModel:           publicSubnetRecord.PriceModel,
		UnitPrice:            publicSubnetRecord.UnitPrice,
		UnitPriceDescription: publicSubnetRecord.UnitPriceDescription,
		Quantity:             publicSubnetRecord.Quantity,
		Active:               publicSubnetRecord.Active,
		UsageSessionId:       publicSubnetRecord.UsageSessionId,
		CorrelationId:        publicSubnetRecord.CorrelationId,
		ReservationId:        DerefString(publicSubnetRecord.ReservationId),
		Metadata:             metadata,
	}
}

func fromServerRecord(serverRecord billingapisdk.ServerRecord) RatedUsageRecordTable {
	metadata := make(map[string]interface{})

	metadata[SERVER_ID] = serverRecord.Metadata.Id
	metadata[HOSTNAME] = serverRecord.Metadata.Hostname

	return RatedUsageRecordTable{
		Id:                   serverRecord.Id,
		ProductCategory:      serverRecord.ProductCategory,
		ProductCode:          serverRecord.ProductCode,
		Location:             string(serverRecord.Location),
		YearMonth:            DerefString(serverRecord.YearMonth),
		StartDateTime:        serverRecord.StartDateTime.String(),
		EndDateTime:          serverRecord.EndDateTime.String(),
		Cost:                 serverRecord.Cost,
		PriceModel:           serverRecord.PriceModel,
		UnitPrice:            serverRecord.UnitPrice,
		UnitPriceDescription: serverRecord.UnitPriceDescription,
		Quantity:             serverRecord.Quantity,
		Active:               serverRecord.Active,
		UsageSessionId:       serverRecord.UsageSessionId,
		CorrelationId:        serverRecord.CorrelationId,
		ReservationId:        DerefString(serverRecord.ReservationId),
		Metadata:             metadata,
	}
}

// TODO: Finish Short version

type ShortRatedUsageRecordTable struct {
	// TODO: Fill...
}

// Extracts a ShortRatedUsageRecordTable from the OneOf response.
// Done by marshalling into JSON and unmarshalling into a struct that holds a common representation.
func ShortRatedUsageRecordFromSdk(sdkRecord billingapisdk.RatedUsageGet200ResponseInner) ShortRatedUsageRecordTable {
	// Converts the oneOf record into the common representation
	jsonRecord, _ := sdkRecord.MarshalJSON()
	var common ShortRatedUsageRecordCommon
	json.Unmarshal(jsonRecord, &common)

	return ShortRatedUsageRecordTable{
		// TODO: Fill...
	}
}

type ShortRatedUsageRecordCommon struct {
	// TODO: Fill...
}
