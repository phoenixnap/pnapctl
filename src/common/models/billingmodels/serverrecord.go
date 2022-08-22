package billingmodels

import (
	"time"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type ServerRecord struct {
	Id                   string        `json:"id" yaml:"id"`
	ProductCategory      string        `json:"productCategory" yaml:"productCategory"`
	ProductCode          string        `json:"productCode" yaml:"productCode"`
	Location             string        `json:"location" yaml:"location"`
	YearMonth            *string       `json:"yearMonth,omitempty" yaml:"yearMonth,omitempty"`
	StartDateTime        time.Time     `json:"startDateTime" yaml:"startDateTime"`
	EndDateTime          time.Time     `json:"endDateTime" yaml:"endDateTime"`
	Cost                 int64         `json:"cost" yaml:"cost"`
	PriceModel           string        `json:"priceModel" yaml:"priceModel"`
	UnitPrice            float32       `json:"unitPrice" yaml:"unitPrice"`
	UnitPriceDescription string        `json:"unitPriceDescription" yaml:"unitPriceDescription"`
	Quantity             float32       `json:"quantity" yaml:"quantity"`
	Active               bool          `json:"active" yaml:"active"`
	UsageSessionId       string        `json:"usageSessionId" yaml:"usageSessionId"`
	CorrelationId        string        `json:"correlationId" yaml:"correlationId"`
	ReservationId        *string       `json:"reservationId,omitempty" yaml:"reservationId,omitempty"`
	Metadata             ServerDetails `json:"metadata" yaml:"metadata"`
}

func ServerRecordFromSdk(serverRecord *billingapi.ServerRecord) *ServerRecord {
	if serverRecord == nil {
		return nil
	}

	return &ServerRecord{
		Id:                   serverRecord.Id,
		ProductCategory:      serverRecord.ProductCategory,
		ProductCode:          serverRecord.ProductCode,
		Location:             string(serverRecord.Location),
		YearMonth:            serverRecord.YearMonth,
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
		ReservationId:        serverRecord.ReservationId,
		Metadata:             *ServerDetailsFromSdk(&serverRecord.Metadata),
	}
}
