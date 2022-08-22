package billingmodels

import (
	"time"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type PublicSubnetRecord struct {
	Id                   string              `json:"id" yaml:"id"`
	ProductCategory      string              `json:"productCategory" yaml:"productCategory"`
	ProductCode          string              `json:"productCode" yaml:"productCode"`
	Location             string              `json:"location" yaml:"location"`
	YearMonth            *string             `json:"yearMonth,omitempty" yaml:"yearMonth,omitempty"`
	StartDateTime        time.Time           `json:"startDateTime" yaml:"startDateTime"`
	EndDateTime          time.Time           `json:"endDateTime" yaml:"endDateTime"`
	Cost                 int64               `json:"cost" yaml:"cost"`
	PriceModel           string              `json:"priceModel" yaml:"priceModel"`
	UnitPrice            float32             `json:"unitPrice" yaml:"unitPrice"`
	UnitPriceDescription string              `json:"unitPriceDescription" yaml:"unitPriceDescription"`
	Quantity             float32             `json:"quantity" yaml:"quantity"`
	Active               bool                `json:"active" yaml:"active"`
	UsageSessionId       string              `json:"usageSessionId" yaml:"usageSessionId"`
	CorrelationId        string              `json:"correlationId" yaml:"correlationId"`
	ReservationId        *string             `json:"reservationId,omitempty" yaml:"reservationId,omitempty"`
	Metadata             PublicSubnetDetails `json:"metadata" yaml:"metadata"`
}

func PublicSubnetRecordFromSdk(publicSubnetRecord *billingapi.PublicSubnetRecord) *PublicSubnetRecord {
	if publicSubnetRecord == nil {
		return nil
	}

	return &PublicSubnetRecord{
		Id:                   publicSubnetRecord.Id,
		ProductCategory:      publicSubnetRecord.ProductCategory,
		ProductCode:          publicSubnetRecord.ProductCode,
		Location:             string(publicSubnetRecord.Location),
		YearMonth:            publicSubnetRecord.YearMonth,
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
		ReservationId:        publicSubnetRecord.ReservationId,
		Metadata:             *PublicSubnetDetailsFromSdk(&publicSubnetRecord.Metadata),
	}
}
