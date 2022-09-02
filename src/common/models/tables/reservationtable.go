package tables

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

type ReservationTable struct {
	Id                  string
	ProductCode         string
	ProductCategory     string
	Location            string
	ReservationModel    string
	InitialInvoiceModel string
	StartDateTime       string
	EndDateTime         string
	LastRenewalDateTime string
	NextRenewalDateTime string
	AutoRenew           bool
	Sku                 string
	Price               float32
	PriceUnit           string
	AssignedResourceId  string
}

func ReservationTableFromSdk(sdk *billingapi.Reservation) *ReservationTable {
	if sdk == nil {
		return nil
	}

	return &ReservationTable{
		Id:                  sdk.Id,
		ProductCode:         sdk.ProductCode,
		ProductCategory:     sdk.ProductCategory,
		Location:            string(sdk.Location),
		ReservationModel:    string(sdk.ReservationModel),
		InitialInvoiceModel: DerefString(sdk.InitialInvoiceModel),
		StartDateTime:       sdk.StartDateTime.String(),
		EndDateTime:         DerefStringable(sdk.EndDateTime),
		LastRenewalDateTime: DerefStringable(sdk.LastRenewalDateTime),
		NextRenewalDateTime: DerefStringable(sdk.NextRenewalDateTime),
		AutoRenew:           sdk.AutoRenew,
		Sku:                 sdk.Sku,
		Price:               sdk.Price,
		PriceUnit:           string(sdk.PriceUnit),
		AssignedResourceId:  DerefString(sdk.AssignedResourceId),
	}
}
