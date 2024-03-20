package tables

import "github.com/phoenixnap/go-sdk-bmc/billingapi/v3"

type ReservationTable struct {
	Id                  string  `header:"ID"`
	ProductCode         string  `header:"Product Code"`
	ProductCategory     string  `header:"Product Category"`
	Location            string  `header:"Location"`
	ReservationModel    string  `header:"Reservation"`
	InitialInvoiceModel string  `header:"Initial Invoice"`
	StartDateTime       string  `header:"Start Date"`
	EndDateTime         string  `header:"End Date"`
	LastRenewalDateTime string  `header:"Last Renewal"`
	NextRenewalDateTime string  `header:"Next Renewal"`
	AutoRenew           bool    `header:"Auto Renew"`
	Sku                 string  `header:"Sku"`
	Price               float32 `header:"Price"`
	PriceUnit           string  `header:"Unit"`
	AssignedResourceId  string  `header:"Assigned Resource ID"`
	NextBillingDate     string  `header:"Next billing date for Reservation"`
}

type ShortReservationTable struct {
	Id              string  `header:"ID"`
	ProductCode     string  `header:"Product Code"`
	ProductCategory string  `header:"Product Category"`
	Location        string  `header:"Location"`
	Price           float32 `header:"Price"`
	PriceUnit       string  `header:"Unit"`
}

func ReservationTableFromSdk(sdk billingapi.Reservation) ReservationTable {
	return ReservationTable{
		Id:                  sdk.Id,
		ProductCode:         sdk.ProductCode,
		ProductCategory:     string(sdk.ProductCategory),
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
		NextBillingDate:     DerefString(sdk.NextBillingDate),
	}
}

func ShortReservationTableFromSdk(sdk billingapi.Reservation) ShortReservationTable {
	return ShortReservationTable{
		Id:              sdk.Id,
		ProductCode:     sdk.ProductCode,
		ProductCategory: string(sdk.ProductCategory),
		Location:        string(sdk.Location),
		Price:           sdk.Price,
		PriceUnit:       string(sdk.PriceUnit),
	}
}
