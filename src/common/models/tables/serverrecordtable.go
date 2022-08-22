package tables

import (
	"time"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
)

type ServerRecordTable struct {
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

type ShortServerRecordTable struct {
}

func ServerRecordTableFromSdk(serverRecord billingapisdk.ServerRecord) ServerRecordTable {
	return ServerRecordTable{
		Id:                   serverRecord.Id,
		ProductCategory:      serverRecord.ProductCategory,
		ProductCode:          serverRecord.ProductCode,
		Location:             string(serverRecord.Location),
		YearMonth:            DerefString(serverRecord.YearMonth),
		StartDateTime:        serverRecord.StartDateTime,
		EndDateTime:          serverRecord.EndDateTime,
		Cost:                 serverRecord.Cost,
		PriceModel:           serverRecord.PriceModel,
		UnitPrice:            serverRecord.UnitPrice,
		UnitPriceDescription: serverRecord.UnitPriceDescription,
		Quantity:             serverRecord.Quantity,
		Active:               serverRecord.Active,
		UsageSessionId:       serverRecord.UsageSessionId,
		CorrelationId:        serverRecord.CorrelationId,
		ReservationId:        DerefString(serverRecord.ReservationId),
		Metadata:             billingmodels.ServerDetailsToTableString(&serverRecord.Metadata),
	}
}

func ShortServerRecordTableFromSdk(serverRecord billingapisdk.ServerRecord) ShortServerRecordTable {
	return ShortServerRecordTable{}
}
