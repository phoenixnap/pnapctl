package billingmodels

import (
	"time"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type Reservation struct {
	Id                  string                                    `json:"id" yaml:"id"`
	ProductCode         string                                    `json:"productCode" yaml:"productCode"`
	ProductCategory     string                                    `json:"productCategory" yaml:"productCategory"`
	Location            billingapi.LocationEnum                   `json:"location" yaml:"location"`
	ReservationModel    billingapi.ReservationModelEnum           `json:"reservationModel" yaml:"reservationModel"`
	InitialInvoiceModel *billingapi.ReservationInvoicingModelEnum `json:"initialInvoiceModel,omitempty" yaml:"initialInvoiceModel,omitempty"`
	StartDateTime       time.Time                                 `json:"startDateTime" yaml:"startDateTime"`
	EndDateTime         *time.Time                                `json:"endDateTime,omitempty" yaml:"endDateTime,omitempty"`
	LastRenewalDateTime *time.Time                                `json:"lastRenewalDateTime,omitempty" yaml:"lastRenewalDateTime,omitempty"`
	NextRenewalDateTime *time.Time                                `json:"nextRenewalDateTime,omitempty" yaml:"nextRenewalDateTime,omitempty"`
	AutoRenew           bool                                      `json:"autoRenew" yaml:"autoRenew"`
	Sku                 string                                    `json:"sku" yaml:"sku"`
	Price               float32                                   `json:"price" yaml:"price"`
	PriceUnit           billingapi.PriceUnitEnum                  `json:"priceUnit" yaml:"priceUnit"`
	AssignedResourceId  *string                                   `json:"assignedResourceId,omitempty" yaml:"assignedResourceId,omitempty"`
}

func ReservationFromSdk(sdk *billingapi.Reservation) *Reservation {
	if sdk == nil {
		return nil
	}

	return &Reservation{
		Id:                  sdk.Id,
		ProductCode:         sdk.ProductCode,
		ProductCategory:     sdk.ProductCategory,
		Location:            sdk.Location,
		ReservationModel:    sdk.ReservationModel,
		InitialInvoiceModel: sdk.InitialInvoiceModel,
		StartDateTime:       sdk.StartDateTime,
		EndDateTime:         sdk.EndDateTime,
		LastRenewalDateTime: sdk.LastRenewalDateTime,
		NextRenewalDateTime: sdk.NextRenewalDateTime,
		AutoRenew:           sdk.AutoRenew,
		Sku:                 sdk.Sku,
		Price:               sdk.Price,
		PriceUnit:           sdk.PriceUnit,
		AssignedResourceId:  sdk.AssignedResourceId,
	}
}
