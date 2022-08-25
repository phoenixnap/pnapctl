package ratedusageoneof

import (
	"time"

	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

// Uses the discriminator - productCategory
type Discriminator string

const (
	// Bandwidth Record
	BANDWIDTH Discriminator = "bandwidth"

	// Operating System Record
	OPERATING_SYSTEM Discriminator = "operating-system"

	// Public Subnet Record
	PUBLIC_SUBNET Discriminator = "public-ip"

	// Server Record
	SERVER Discriminator = "bmc-server"
)

// Matches all elements in the OneOf due to common fields.
type RatedUsageSdk interface {
	GetId() string
	GetProductCategory() string
	GetProductCode() string
	GetLocation() billingapisdk.LocationEnum
	GetYearMonth() string
	GetStartDateTime() time.Time
	GetEndDateTime() time.Time
	GetCost() int64
	GetPriceModel() string
	GetUnitPrice() float32
	GetUnitPriceDescription() string
	GetQuantity() float32
	GetActive() bool
	GetUsageSessionId() string
	GetCorrelationId() string
	GetReservationId() string
}

// All common fields that the interface maps to.
type RatedUsageCommon struct {
	Id                   string        `json:"id" yaml:"id"`
	ProductCategory      Discriminator `json:"productCategory" yaml:"productCategory"`
	ProductCode          string        `json:"productCode" yaml:"productCode"`
	Location             string        `json:"location" yaml:"location"`
	YearMonth            string        `json:"yearMonth,omitempty" yaml:"yearMonth,omitempty"`
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
	ReservationId        string        `json:"reservationId,omitempty" yaml:"reservationId,omitempty"`
}

// Uses type assertions to convert the OneOf into a RatedUsage.
func RatedUsageFromSdkOneOf(sdk *billingapisdk.RatedUsageGet200ResponseInner) *RatedUsageCommon {
	if sdk == nil {
		return nil
	}

	actualInstance := sdk.GetActualInstance().(RatedUsageSdk)

	return &RatedUsageCommon{
		Id:                   actualInstance.GetId(),
		ProductCategory:      Discriminator(actualInstance.GetProductCategory()),
		ProductCode:          actualInstance.GetProductCode(),
		Location:             string(actualInstance.GetLocation()),
		YearMonth:            actualInstance.GetYearMonth(),
		StartDateTime:        actualInstance.GetStartDateTime(),
		EndDateTime:          actualInstance.GetEndDateTime(),
		Cost:                 actualInstance.GetCost(),
		PriceModel:           actualInstance.GetPriceModel(),
		UnitPrice:            actualInstance.GetUnitPrice(),
		UnitPriceDescription: actualInstance.GetUnitPriceDescription(),
		Quantity:             actualInstance.GetQuantity(),
		Active:               actualInstance.GetActive(),
		UsageSessionId:       actualInstance.GetUsageSessionId(),
		CorrelationId:        actualInstance.GetCorrelationId(),
		ReservationId:        actualInstance.GetReservationId(),
	}
}

// Checks whether the struct type is actually any of the ones passed.
func (rec *RatedUsageCommon) IsActually(ds ...Discriminator) bool {
	return iterutils.Contains(ds, rec.ProductCategory)
}
