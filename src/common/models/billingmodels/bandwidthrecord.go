package billingmodels

import (
	"time"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type BandwidthRecord struct {
	Id                   string           `json:"id" yaml:"id"`
	ProductCategory      string           `json:"productCategory" yaml:"productCategory"`
	ProductCode          string           `json:"productCode" yaml:"productCode"`
	Location             string           `json:"location" yaml:"location"`
	YearMonth            *string          `json:"yearMonth,omitempty" yaml:"yearMonth,omitempty"`
	StartDateTime        time.Time        `json:"startDateTime" yaml:"startDateTime"`
	EndDateTime          time.Time        `json:"endDateTime" yaml:"endDateTime"`
	Cost                 int64            `json:"cost" yaml:"cost"`
	PriceModel           string           `json:"priceModel" yaml:"priceModel"`
	UnitPrice            float32          `json:"unitPrice" yaml:"unitPrice"`
	UnitPriceDescription string           `json:"unitPriceDescription" yaml:"unitPriceDescription"`
	Quantity             float32          `json:"quantity" yaml:"quantity"`
	Active               bool             `json:"active" yaml:"active"`
	UsageSessionId       string           `json:"usageSessionId" yaml:"usageSessionId"`
	CorrelationId        string           `json:"correlationId" yaml:"correlationId"`
	ReservationId        *string          `json:"reservationId,omitempty" yaml:"reservationId,omitempty"`
	Metadata             BandwidthDetails `json:"metadata" yaml:"metadata"`
}

func BandwidthRecordFromSdk(bandwidthRecord *billingapi.BandwidthRecord) *BandwidthRecord {
	if bandwidthRecord == nil {
		return nil
	}

	return &BandwidthRecord{
		Id:                   bandwidthRecord.Id,
		ProductCategory:      bandwidthRecord.ProductCategory,
		ProductCode:          bandwidthRecord.ProductCode,
		Location:             string(bandwidthRecord.Location),
		YearMonth:            bandwidthRecord.YearMonth,
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
		ReservationId:        bandwidthRecord.ReservationId,
		Metadata:             *BandwidthDetailsFromSdk(&bandwidthRecord.Metadata),
	}
}
