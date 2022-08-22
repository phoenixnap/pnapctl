package billingmodels

import (
	"time"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type OperatingSystemRecord struct {
	Id                   string                 `json:"id" yaml:"id"`
	ProductCategory      string                 `json:"productCategory" yaml:"productCategory"`
	ProductCode          string                 `json:"productCode" yaml:"productCode"`
	Location             string                 `json:"location" yaml:"location"`
	YearMonth            *string                `json:"yearMonth,omitempty" yaml:"yearMonth,omitempty"`
	StartDateTime        time.Time              `json:"startDateTime" yaml:"startDateTime"`
	EndDateTime          time.Time              `json:"endDateTime" yaml:"endDateTime"`
	Cost                 int64                  `json:"cost" yaml:"cost"`
	PriceModel           string                 `json:"priceModel" yaml:"priceModel"`
	UnitPrice            float32                `json:"unitPrice" yaml:"unitPrice"`
	UnitPriceDescription string                 `json:"unitPriceDescription" yaml:"unitPriceDescription"`
	Quantity             float32                `json:"quantity" yaml:"quantity"`
	Active               bool                   `json:"active" yaml:"active"`
	UsageSessionId       string                 `json:"usageSessionId" yaml:"usageSessionId"`
	CorrelationId        string                 `json:"correlationId" yaml:"correlationId"`
	ReservationId        *string                `json:"reservationId,omitempty" yaml:"reservationId,omitempty"`
	Metadata             OperatingSystemDetails `json:"metadata" yaml:"metadata"`
}

func OperatingSystemRecordFromSdk(operatingSystemRecord *billingapi.OperatingSystemRecord) *OperatingSystemRecord {
	if operatingSystemRecord == nil {
		return nil
	}

	return &OperatingSystemRecord{
		Id:                   operatingSystemRecord.Id,
		ProductCategory:      operatingSystemRecord.ProductCategory,
		ProductCode:          operatingSystemRecord.ProductCode,
		Location:             string(operatingSystemRecord.Location),
		YearMonth:            operatingSystemRecord.YearMonth,
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
		ReservationId:        operatingSystemRecord.ReservationId,
		Metadata:             *OperatingSystemDetailsFromSdk(&operatingSystemRecord.Metadata),
	}
}
