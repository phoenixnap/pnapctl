package tables

import (
	"encoding/json"
	"time"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

// Full Table

type RatedUsageRecordTable struct {
	Id                   string                 `header:"Id"`
	ProductCategory      string                 `header:"Product Category"`
	ProductCode          string                 `header:"Product Code"`
	Location             string                 `header:"Location"`
	YearMonth            string                 `header:"Year Month"`
	StartDateTime        time.Time              `header:"Start Date Time"`
	EndDateTime          time.Time              `header:"End Date Time"`
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

// Extracts a RatedUsageRecordTable from the OneOf response.
// Done by marshalling into JSON and unmarshalling into a struct that holds a common representation.
func RatedUsageRecordFromSdk(sdkRecord billingapisdk.RatedUsageGet200ResponseInner) RatedUsageRecordTable {
	// Converts the oneOf record into the common representation
	jsonRecord, _ := sdkRecord.MarshalJSON()
	var common RatedUsageRecordCommon
	json.Unmarshal(jsonRecord, &common)

	return RatedUsageRecordTable{
		Id:                   common.Id,
		ProductCategory:      common.ProductCategory,
		ProductCode:          common.ProductCode,
		Location:             common.Location,
		YearMonth:            DerefString(common.YearMonth),
		StartDateTime:        common.StartDateTime,
		EndDateTime:          common.EndDateTime,
		Cost:                 common.Cost,
		PriceModel:           common.PriceModel,
		UnitPrice:            common.UnitPrice,
		UnitPriceDescription: common.UnitPriceDescription,
		Quantity:             common.Quantity,
		Active:               common.Active,
		UsageSessionId:       common.UsageSessionId,
		CorrelationId:        common.CorrelationId,
		ReservationId:        DerefString(common.ReservationId),
		Metadata:             common.Metadata,
	}
}

type RatedUsageRecordCommon struct {
	Id                   string                 `json:"id"`
	ProductCategory      string                 `json:"productCategory"`
	ProductCode          string                 `json:"productCode"`
	Location             string                 `json:"location"`
	YearMonth            *string                `json:"yearMonth"`
	StartDateTime        time.Time              `json:"startDateTime"`
	EndDateTime          time.Time              `json:"endDateTime"`
	Cost                 int64                  `json:"cost"`
	PriceModel           string                 `json:"priceModel"`
	UnitPrice            float32                `json:"unitPrice"`
	UnitPriceDescription string                 `json:"unitPriceDescription"`
	Quantity             float32                `json:"quantity"`
	Active               bool                   `json:"active"`
	UsageSessionId       string                 `json:"usageSessionId"`
	CorrelationId        string                 `json:"correlationId"`
	ReservationId        *string                `json:"reservationId"`
	Metadata             map[string]interface{} `json:"Metadata"`
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
