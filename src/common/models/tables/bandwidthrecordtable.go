package tables

import (
	"time"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
)

type BandwidthRecordTable struct {
	Id                   string    `header:"Id"`
	ProductCategory      string    `header:"Product Category"`
	ProductCode          string    `header:"Product Code"`
	Location             string    `header:"Location"`
	YearMonth            string    `header:"Year Month"`
	StartDateTime        time.Time `header:"Start Date Time"`
	EndDateTime          time.Time `header:"End Date Time"`
	Cost                 int64     `header:"Cost"`
	PriceModel           string    `header:"Price Model"`
	UnitPrice            float32   `header:"Unit Price"`
	UnitPriceDescription string    `header:"Unit Price Description"`
	Quantity             float32   `header:"Quantity"`
	Active               bool      `header:"Active"`
	UsageSessionId       string    `header:"Usage Session Id"`
	CorrelationId        string    `header:"Correlation Id"`
	ReservationId        string    `header:"Reservation Id"`
	Metadata             string    `header:"Metadata"`
}

type ShortBandwidthRecordTable struct {
}

func BandwidthRecordTableFromSdk(bandwidthRecord billingapisdk.BandwidthRecord) BandwidthRecordTable {
	return BandwidthRecordTable{
		Id:                   bandwidthRecord.Id,
		ProductCategory:      bandwidthRecord.ProductCategory,
		ProductCode:          bandwidthRecord.ProductCode,
		Location:             string(bandwidthRecord.Location),
		YearMonth:            DerefString(bandwidthRecord.YearMonth),
		StartDateTime:        bandwidthRecord.StartDateTime,
		EndDateTime:          bandwidthRecord.EndDateTime,
		Cost:                 bandwidthRecord.Cost,
		PriceModel:           bandwidthRecord.PriceModel,
		UnitPrice:            bandwidthRecord.UnitPrice,
		UnitPriceDescription: bandwidthRecord.UnitPriceDescription,
		Quantity:             bandwidthRecord.Quantity,
		Active:               bandwidthRecord.Active,
		UsageSessionId:       bandwidthRecord.UsageSessionId,
		CorrelationId:        bandwidthRecord.CorrelationId,
		ReservationId:        DerefString(bandwidthRecord.ReservationId),
		Metadata:             billingmodels.BandwidthDetailsToTableString(&bandwidthRecord.Metadata),
	}
}

func ShortBandwidthRecordTableFromSdk(bandwidthRecord billingapisdk.BandwidthRecord) ShortBandwidthRecordTable {
	return ShortBandwidthRecordTable{}
}
