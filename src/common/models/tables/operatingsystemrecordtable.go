package tables

import (
	"time"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
)

type OperatingSystemRecordTable struct {
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

type ShortOperatingSystemRecordTable struct {
}

func OperatingSystemRecordTableFromSdk(operatingSystemRecord billingapisdk.OperatingSystemRecord) OperatingSystemRecordTable {
	return OperatingSystemRecordTable{
		Id:                   operatingSystemRecord.Id,
		ProductCategory:      operatingSystemRecord.ProductCategory,
		ProductCode:          operatingSystemRecord.ProductCode,
		Location:             string(operatingSystemRecord.Location),
		YearMonth:            DerefString(operatingSystemRecord.YearMonth),
		StartDateTime:        operatingSystemRecord.StartDateTime,
		EndDateTime:          operatingSystemRecord.EndDateTime,
		Cost:                 operatingSystemRecord.Cost,
		PriceModel:           operatingSystemRecord.PriceModel,
		UnitPrice:            operatingSystemRecord.UnitPrice,
		UnitPriceDescription: operatingSystemRecord.UnitPriceDescription,
		Quantity:             operatingSystemRecord.Quantity,
		Active:               operatingSystemRecord.Active,
		UsageSessionId:       operatingSystemRecord.UsageSessionId,
		CorrelationId:        operatingSystemRecord.CorrelationId,
		ReservationId:        DerefString(operatingSystemRecord.ReservationId),
		Metadata:             billingmodels.OperatingSystemDetailsToTableString(&operatingSystemRecord.Metadata),
	}
}

func ShortOperatingSystemRecordTableFromSdk(operatingSystemRecord billingapisdk.OperatingSystemRecord) ShortOperatingSystemRecordTable {
	return ShortOperatingSystemRecordTable{}
}
