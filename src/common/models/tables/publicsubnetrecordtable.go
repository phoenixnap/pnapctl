package tables

import (
	"time"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels"
)

type PublicSubnetRecordTable struct {
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

type ShortPublicSubnetRecordTable struct {
}

func PublicSubnetRecordTableFromSdk(publicSubnetRecord billingapisdk.PublicSubnetRecord) PublicSubnetRecordTable {
	return PublicSubnetRecordTable{
		Id:                   publicSubnetRecord.Id,
		ProductCategory:      publicSubnetRecord.ProductCategory,
		ProductCode:          publicSubnetRecord.ProductCode,
		Location:             string(publicSubnetRecord.Location),
		YearMonth:            DerefString(publicSubnetRecord.YearMonth),
		StartDateTime:        publicSubnetRecord.StartDateTime,
		EndDateTime:          publicSubnetRecord.EndDateTime,
		Cost:                 publicSubnetRecord.Cost,
		PriceModel:           publicSubnetRecord.PriceModel,
		UnitPrice:            publicSubnetRecord.UnitPrice,
		UnitPriceDescription: publicSubnetRecord.UnitPriceDescription,
		Quantity:             publicSubnetRecord.Quantity,
		Active:               publicSubnetRecord.Active,
		UsageSessionId:       publicSubnetRecord.UsageSessionId,
		CorrelationId:        publicSubnetRecord.CorrelationId,
		ReservationId:        DerefString(publicSubnetRecord.ReservationId),
		Metadata:             billingmodels.PublicSubnetDetailsToTableString(&publicSubnetRecord.Metadata),
	}
}

func ShortPublicSubnetRecordTableFromSdk(publicSubnetRecord billingapisdk.PublicSubnetRecord) ShortPublicSubnetRecordTable {
	return ShortPublicSubnetRecordTable{}
}
